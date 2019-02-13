package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type whiskInput struct {
	Value json.RawMessage `json:"value"`
}

type whiskOutput struct {
	Value json.RawMessage `json:"value"`
}

func main() {
	// debugging
	var debug = os.Getenv("OW_DEBUG") != ""

	if debug {
		filename := os.Getenv("OW_DEBUG")
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(f)
			defer f.Close()
		}
		log.Printf("ACTION ENV: %v", os.Environ())
	}

	// input
	out := os.NewFile(3, "pipe")
	defer out.Close()
	reader := bufio.NewReader(os.Stdin)

	// read-eval-print loop
	if debug {
		log.Println("started")
	}
	for {
		// read one line
		inbuf, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		if debug {
			log.Printf(">>>'%s'>>>", inbuf)
		}
		// parse one line
		var rawInput map[string]interface{}
		input := whiskInput{}
		err = json.Unmarshal(inbuf, &rawInput)
		if err != nil {
			log.Println(err.Error())
			fmt.Fprintf(out, "{ error: %q}\n", err.Error())
			continue
		}
		err = json.Unmarshal(inbuf, &input)
		if err != nil {
			log.Println(err.Error())
			fmt.Fprintf(out, "{ error: %q}\n", err.Error())
			continue
		}
		if debug {
			log.Printf("%v\n", rawInput)
		}
		// set environment variables
		err = json.Unmarshal(inbuf, &rawInput)
		for k, v := range rawInput {
			if k == "value" {
				continue
			}
			if s, ok := v.(string); ok {
				os.Setenv("__OW_"+strings.ToUpper(k), s)
			}
		}

		// process the request
		result, _ := callHandler(input)
		// encode the answer
		output, err := json.Marshal(&result)
		if err != nil {
			log.Println(err.Error())
			fmt.Fprintf(out, "{ error: %q}\n", err.Error())
			continue
		}
		output = bytes.Replace(output, []byte("\n"), []byte(""), -1)
		if debug {
			log.Printf("'<<<%s'<<<", output)
		}
		fmt.Fprintf(out, "%s\n", output)
	}
}

type anzerIn struct {
	Language string
}

type anzerOut struct {
	Assegnee string
}

// Main selects assignee for MR
func callHandler(input whiskInput) (whiskOutput, error) {
	anzHandlerInput := anzerIn{}
	if err := json.Unmarshal(input.Value, &anzHandlerInput); err != nil {
		return whiskOutput{}, err
	}
	anzHandlerOutput := handle(typeIn(anzHandlerInput))
	handlerOutput := anzerOut(anzHandlerOutput)
	output, err := json.Marshal(handlerOutput)
	return whiskOutput{
		Value: output,
	}, err
}
