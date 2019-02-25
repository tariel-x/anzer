package golang

func (cg Generator) GenerateDocker() string {
	return dockerfile
}

var dockerfile = `
FROM go:1.11
`
