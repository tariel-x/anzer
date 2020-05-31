package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"

	"github.com/tariel-x/anzer/pkg/cache"
	"github.com/tariel-x/anzer/pkg/git"
	l "github.com/tariel-x/anzer/pkg/lang"
	"github.com/tariel-x/anzer/pkg/platform"
	"github.com/tariel-x/anzer/pkg/platform/models"
)

const (
	defaultCacheLocation = "/.anzer_cache"
	defaultSumFileName   = "anzer.sum"
)

type BuildCmd struct {
	input    string
	platform platform.Platform
	cache    *cache.Manager
}

func Build(c *cli.Context) error {
	cmd, err := newBuildCmd(c)
	if err != nil {
		return err
	}
	return cmd.build()
}

func newBuildCmd(c *cli.Context) (*BuildCmd, error) {
	plat, err := getPlatform(c)
	if err != nil {
		return nil, err
	}
	if err := plat.Connect(); err != nil {
		return nil, err
	}

	input := c.String("input")
	if input == "" {
		return nil, errNoInput
	}

	cacheLocation := c.String("cacheLocation")
	if cacheLocation == "" {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}
		cacheLocation = filepath.Join(usr.HomeDir, defaultCacheLocation)
	}

	f, err := os.OpenFile(defaultSumFileName, os.O_RDWR, 0666)
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(defaultSumFileName)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	defer f.Close()
	cm, err := cache.NewManager(f, cacheLocation)
	if err != nil {
		return nil, err
	}

	cmd := &BuildCmd{
		input:    input,
		platform: plat,
		cache:    cm,
	}
	return cmd, nil
}

func (b *BuildCmd) build() error {
	f, err := os.Open(b.input)
	if err != nil {
		return err
	}
	defer f.Close()

	composes, err := platform.ParseLazy(f)
	if err != nil {
		return err
	}

	for _, compose := range composes {
		if err := b.buildCompose(compose); err != nil {
			return err
		}
	}
	sumFile, err := os.Create(defaultSumFileName)
	if err != nil {
		return fmt.Errorf("can not write to %s: %w", defaultSumFileName, err)
	}
	defer sumFile.Close()
	return b.cache.Flush(sumFile)
}

func (b *BuildCmd) buildCompose(compose l.Composable) error {
	log.Printf("build composition %s = %s", compose.GetName(), compose.Definition())

	if err := compose.Invalid(); err != nil {
		return err
	}

	chain, err := b.toChain(compose)
	if err != nil {
		return err
	}

	components := make([]models.PublishedFunction, 0, len(chain))
	for _, f := range chain {
		action, err := b.resolveFunc(f)
		if err != nil {
			return err
		}
		component, err := b.publishFunc(f, action)
		if err != nil {
			return fmt.Errorf("publish function %s: %w", f.Definition(), err)
		}
		components = append(components, component)
	}

	log.Printf("make link for %v", components)
	lnk, err := b.platform.Link(compose.GetName(), components)
	if err != nil {
		return err
	}
	fmt.Println("Your function name is:", aurora.Cyan(lnk.Name))
	return nil
}

func (b *BuildCmd) toChain(f l.Composable) ([]l.Runnable, error) {
	chain := []l.Runnable{}
	switch ft := f.(type) {
	case l.Alias:
		for _, sf := range ft.Compose {
			subchain, err := b.toChain(sf)
			if err != nil {
				return nil, err
			}
			chain = append(chain, subchain...)
		}
	case l.F:
		chain = append(chain, ft)
	case l.EitherBind:
		chain = append(chain, ft)
	}
	return chain, nil
}

func (b *BuildCmd) resolveFunc(f l.Runnable) (io.Reader, error) {
	action, commitID, err := b.loadCached(f)
	if err == nil {
		log.Printf("loaded cached function %s %s", f.GetName(), commitID)
		return action, nil
	}

	log.Printf("error loading cached function %s", err)

	if commitID == "" {
		commitID, err = b.findLatestCommitID(f)
		if err != nil {
			return nil, fmt.Errorf("can not find latest commit id for %s: %w", f.Definition(), err)
		}
	}

	log.Printf("build function %s@%s", f.Definition(), commitID)
	action, err = b.buildFunc(f, commitID)
	if err != nil {
		return nil, fmt.Errorf("build function %s: %w", f.Definition(), err)
	}
	location, err := b.cache.SetFunction(f.GetLink().String(), commitID, nil)
	if err != nil {
		return nil, err
	}

	zipFile, err := ioutil.ReadAll(action)
	if err != nil {
		return nil, err
	}
	if err := ioutil.WriteFile(location, zipFile, 0666); err != nil {
		return nil, err
	}

	return action, nil
}

func (b *BuildCmd) loadCached(f l.Runnable) (io.Reader, string, error) {
	location, commitID, err := b.cache.FunctionWithCommit(f.GetLink().String(), nil)
	if err != nil {
		return nil, commitID, err
	}
	file, err := os.Open(location)
	return file, commitID, err
}

func (b *BuildCmd) buildFunc(f l.Runnable, commitID string) (io.Reader, error) {
	dockerGenerator, err := platform.GetDockerGenerator(f.GetRuntime())
	if err != nil {
		return nil, err
	}
	opts, err := dockerGenerator.GetBuildOptions(&models.BuildOpts{
		Debug:    false,
		F:        f,
		CommitID: commitID,
	})
	if err != nil {
		return nil, err
	}
	env, err := getEnv()
	if err != nil {
		return nil, err
	}
	opts.Env = env

	builder, err := platform.NewBuilder()
	if err != nil {
		return nil, err
	}

	return builder.BuildWithImage(opts, f.GetLink())
}

func (b *BuildCmd) publishFunc(f l.Runnable, action io.Reader) (models.PublishedFunction, error) {
	name := strings.Replace(string(f.GetLink()), "/", "_", -1)
	function, err := b.platform.Create(action, name, f.GetRuntime())
	if err != nil {
		return models.PublishedFunction{}, err
	}
	return function, nil
}

func (b *BuildCmd) findLatestCommitID(f l.Runnable) (string, error) {
	repo, err := git.NewRepository(f.GetLink().String())
	if err != nil {
		return "", err
	}
	return repo.LatestCommit()
}
