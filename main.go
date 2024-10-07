package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"colorscheme/render"
)

//go:embed themes.json
var data []byte
var themes map[string][]string

func init() {
	if err := json.Unmarshal(data, &themes); err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		for name := range themes {
			fmt.Printf("%s\n", name)
		}

		return
	}

	colors := themes[args[0]]
	if len(colors) == 0 {
		os.Exit(1)
	}

	if len(args) == 1 {
		os.Exit(1)
	}

	render.Init(args[1])

	if len(args) == 2 {
		render.Render(args[0], colors, -1)
	} else {
		args = args[2:]
		for _, arg := range args {
			index := 0
			fmt.Sscanf(arg, "%d", &index)
			render.Render(args[0], colors, index)
		}
	}
}
