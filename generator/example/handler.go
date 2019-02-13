package main

import (
	"log"
	"math/rand"
)

type anzerIn struct {
	Language string
}

type anzerOut struct {
	Assegnee string
}

var (
	humans = map[string][]string{
		"ivan":     {"go", "php"},
		"nikita":   {"go", "php"},
		"alexandr": {"go", "js"},
		"valentin": {"go", "js", "php"},
		"maxim":    {"js"},
		"alice":    {"java"},
	}
)

func handle(input anzerIn) anzerOut {
	candidates := make([]string, 0, len(humans))
	if input.Language == "" {
		for human := range humans {
			candidates = append(candidates, human)
		}
	} else {

		for human, langs := range humans {
			if in(input.Language, langs) {
				candidates = append(candidates, human)
			}
		}
	}
	log.Printf("Candidates num %d", len(candidates))

	num := rand.Intn(len(candidates))

	log.Printf("New assegnee %s for language %s", candidates[num], input.Language)

	out := anzerOut{
		Assegnee: candidates[num],
	}
	return out
}

func in(search string, scope []string) bool {
	for _, occurrence := range scope {
		if search == occurrence {
			return true
		}
	}
	return false
}
