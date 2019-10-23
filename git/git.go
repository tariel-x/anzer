package git

type Repository struct {
}

func NewRepository(repo string) (*Repository, error) {
	//r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
	//	URL: "https://github.com/src-d/go-siva",
	//})
	return nil, nil
}

func (r *Repository) LatestCommit() (string, error) {
	return "", nil
}
