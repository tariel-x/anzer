package main

import "github.com/tariel-x/anzer/types"

type Service struct {
	InType  types.JsonSchema
	OutType types.JsonSchema
	Config  struct {
		EnvIn  string
		EnvOut string
		Envs   map[string]string
	}
}

type Link struct {
}
