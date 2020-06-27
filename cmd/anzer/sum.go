package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli"

	"github.com/tariel-x/anzer/pkg/cache"
	"github.com/tariel-x/anzer/pkg/git"
)

type SumCmd struct {
	cache *cache.Manager
}

func (c *SumCmd) findLatestCommitID(f string) (string, error) {
	repo, err := git.NewRepository(f)
	if err != nil {
		return "", err
	}
	return repo.LatestCommit()
}

func newSumCmd(c *cli.Context) (*SumCmd, error) {
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

	cmd := &SumCmd{cache: cm}

	return cmd, nil
}

func (s *SumCmd) update(f string) error {
	commitID, err := s.findLatestCommitID(f)
	if err != nil {
		return err
	}
	if _, err := s.cache.SetFunction(f, commitID, nil); err != nil {
		return err
	}
	sumFile, err := os.Create(defaultSumFileName)
	if err != nil {
		return fmt.Errorf("can not write to %s: %w", defaultSumFileName, err)
	}
	defer sumFile.Close()
	return s.cache.Flush(sumFile)
}

func SumUpdate(c *cli.Context) error {
	s, err := newSumCmd(c)
	if err != nil {
		return err
	}
	return s.update(c.String("function"))
}
