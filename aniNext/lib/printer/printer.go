package printer

import "github.com/fatih/color"

func PrintRed() func(a ...interface{}) string {
	col := color.New(color.FgRed).SprintFunc()
	return col
}

func PrintGreen() func(a ...interface{}) string {
	col := color.New(color.FgGreen).SprintFunc()
	return col
}

func PrintCyan() func(a ...interface{}) string {
	col := color.New(color.FgCyan).SprintFunc()
	return col
}
