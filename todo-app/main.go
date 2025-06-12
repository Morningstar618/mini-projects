package main

import (
	"flag"
	"fmt"
	"os"

	"todo/lib/printer"
	"todo/lib/todo"
)

const filename = "todo.csv"

func main() {
	list := flag.Bool("l", false, "List todo")
	delete := flag.String("d", "", "delete todo")
	update := flag.String("u", "", "update todo")
	add := flag.String("a", "", "add todo")

	flag.Parse()

	if len(os.Args) < 2 {
		green := printer.PrintGreen()
		fmt.Printf("[*] Usage: %s\n", green("todo -h"))
	}

	file, err := os.Open(filename)
	if err != nil {
		os.WriteFile(filename, []byte(""), 0644)
	}

	if len(*delete) > 0 {
		todo.DeleteTodo(delete, file)
	}

	if len(*update) > 0 {
		todo.UpdateTodo(update, file)
	}

	if len(*add) > 0 {
		todo.AddTodo(add, file)
	}

	if *list {
		todo.ListTodo(file)
	}

	file.Close()
}
