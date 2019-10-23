package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type Repository struct {
	repo *git.Repository
}

func NewRepository(repo string) (*Repository, error) {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/src-d/go-siva",
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
