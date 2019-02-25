package generator

func (cg CodeGenerator) GenerateDocker() string {
	return dockerfile
}

var dockerfile = `
FROM go:1.11
`
