package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todo/lib/printer"

	"github.com/fatih/color"
)

const filename = "todo.csv"

func ListTodo(file *os.File) {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("[!] Unable to read csv file!")
	}

	// <------- Colors ---------->
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Println("No. | Todo")
	fmt.Println("----|-------------------")

	// <------ Dynamic Padding -------->
	columnWidth := 4

	for _, record := range records {
		id := strings.TrimSpace(record[0])
		task := strings.TrimSpace(record[1])

		coloredID := yellow(id)
		coloredTask := cyan(task)

		paddingNeeded := max(columnWidth-len(id), 0)
		padding := strings.Repeat(" ", paddingNeeded)

		fmt.Printf("%s%s| %s\n", coloredID, padding, coloredTask)
	}
}

func DeleteTodo(delete *string, file *os.File) {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		red := printer.PrintRed()
		fmt.Printf("%s\n", red("[!] Failed to read csv file while deleting todo"))
	}

	var newRecord [][]string
	for _, record := range records {
		id := strings.TrimSpace(record[0])

		if strings.Compare(id, *delete) != 0 {
			newRecord = append(newRecord, record)
		}
	}

	writerFile, err := os.Create(filename)
	if err != nil {
		red := printer.PrintRed()
		fmt.Printf("%s\n", red("[!] Failed to truncate file while deleting todo"))
	}

	writer := csv.NewWriter(writerFile)
	writer.WriteAll(newRecord)
	writer.Flush()

	if err := writer.Error(); err != nil {
		red := printer.PrintRed()
		fmt.Printf("%s\n", red("[!] Failed to write to file"))
	}
}

func UpdateTodo(update *string, file *os.File) {
	splitString := strings.Split(*update, ",")
	update_id := splitString[0]
	newTask := splitString[1]

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		red := printer.PrintRed()
		fmt.Printf("%s\n", red("[!] Failed to read csv file while updating todo"))
	}

	var newRecord [][]string

	for _, record := range records {
		id := strings.TrimSpace(record[0])

		if strings.Compare(id, update_id) == 0 {
			record[1] = newTask
		}

		newRecord = append(newRecord, record)
	}

	writerFile, err := os.Create(filename)
	if err != nil {
		red := printer.PrintRed()
		fmt.Printf("%s\n", red("[!] Failed to truncate csv file while updating todo"))
	}

	writer := csv.NewWriter(writerFile)
	writer.WriteAll(newRecord)
	writer.Flush()

	if writer.Error() != nil {
		red := printer.PrintRed()
		fmt.Printf("%s\n", red("[!] Failed to write to the csv file while updating todo"))
	}
}

func AddTodo(add *string, file *os.File) {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		green := printer.PrintGreen()
		fmt.Printf("%s\n", green("[*] Creating todo.csv"))

		_, err = os.Create(filename)
		if err != nil {
			red := printer.PrintRed()
			fmt.Printf("%s\n", red("[!] Failed to create new file while adding todo"))
		}
	}

	newTask := *add
	var newID int

	if len(records) == 0 {
		newID = 1
	} else {

		tempID, _ := strconv.Atoi(records[len(records)-1][0])
		if err != nil {
			red := printer.PrintRed()
			fmt.Printf("%s\n", red("[!] Failed to convert string to int while adding todo"))
		}

		newID = tempID + 1
	}

	newRecord := []string{strconv.Itoa(newID), newTask}
	records = append(records, newRecord)

	writerFile, err := os.Create(filename)
	if err != nil {
		red := printer.PrintRed()
		fmt.Printf("%s\n", red("[!] Failed to truncate file while adding todo"))
	}

	writer := csv.NewWriter(writerFile)
	writer.WriteAll(records)
	writer.Flush()

	if err := writer.Error(); err != nil {
		red := printer.PrintRed()
		fmt.Printf("%s\n", red("[!] Failed to write to file"))
	}
}
