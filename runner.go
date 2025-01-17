package aoc

import (
	"fmt"

	internal "github.com/igorwulff/aoc/internal"
)

func Run() {
	args := internal.ProcessArgs()

	plugin := internal.PluginProcessor{Args: args}
	if err := plugin.Build(); err != nil {
		fmt.Println(err)
		return
	}

	input, err := plugin.GetInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	plugin.RunTests()

	output, err := plugin.CallFunc(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nSolution:\n", output)
}
