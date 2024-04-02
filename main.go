package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a command. Use 'help' to see available commands.")
	}

	// TODO::
	// - [] Extract flags from args
	// - [] Add help command
	// - [x] Add list command
	// - [] Add create command
	// - [] Add update command
	// - [] Add delete command
	// - [] Add some sort of data store. Probably just going to use a static CSV file for now.

	// Add CRUD endpoints for todo list items
	store, err := GetStore()
	if err != nil {
		if err.Error() == "open store.csv: no such file or directory" {
			CreateStore()
			store, err = GetStore()
		} else {
			fmt.Println(err)
			return
		}
	}

	switch args[0] {
	case "list":
		store.ListItems()
	case "create-test-store":
		store.CreateTestStore()
	case "create-item":
		if len(args) < 2 {
			fmt.Println("Please provide a name for the item")
			return
		}
		store.CreateItem(args[1])
	}
}
