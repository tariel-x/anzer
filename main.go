package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
