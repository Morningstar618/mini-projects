package printer

import "github.com/fatih/color"

func PrintRed() func(a ...interface{}) string {
	red := color.New(color.FgRed).SprintFunc()
	return red
}

func PrintGreen() func(a ...interface{}) string {
	red := color.New(color.FgGreen).SprintFunc()
	return red
}
