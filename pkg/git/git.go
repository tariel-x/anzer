package git

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type Repository struct {
	repo *git.Repository
}

func NewRepository(repo string) (*Repository, error) {
	//TODO: not only https
	rawUrl := fmt.Sprintf("https://%s", repo)
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	p := strings.Split(u.Path, "/")
	if len(p) > 3 {
		u.Path = path.Join(p[:3]...)
	}
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: u.String(),
	})
	if err != nil {
		return nil, err
	}
	return &Repository{repo: r}, nil
}

func (r *Repository) LatestCommit() (string, error) {
	log, err := r.repo.Log(&git.LogOptions{})
	if err != nil {
		return "", err
	}
	defer log.Close()
	commit, err := log.Next()
	if err != nil {
		return "", err
	}
	return commit.Hash.String(), nil
}
