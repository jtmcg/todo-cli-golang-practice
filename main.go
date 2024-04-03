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
	// - [x] Add create command
	// - [x] Add update command
	// - [x] Add delete command
	// - [x] Add some sort of data store. Probably just going to use a static CSV file for now.
	// - [] Add todo list-type functionality

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
		if len(args) == 1 || args[1] == "all" {
			store.ListItems()
		} else if len(args) == 2 {
			store.GetItemByName(args[1])
		} else {
			fmt.Println("Too many arguments for list command")
		}
	case "create-test-store":
		store.CreateTestStore()
	case "create-item":
		if len(args) < 2 {
			fmt.Println("Please provide a name for the item")
			return
		}
		store.CreateItem(args[1], args[2])
	case "delete-item":
		if len(args) < 2 {
			fmt.Println("Please provide a name for the item to delete")
			return
		}
		store.DeleteItem(args[1])
	case "update-item":
		if len(args) < 4 {
			fmt.Println("Please provide a name for the item to update, the field to update, and the new value")
			return
		}
		store.UpdateItem(args[1], args[2], args[3])
	case "progress:":
		if len(args) < 2 {
			fmt.Println("Please provide a name for the item to progress")
			return
		}
		store.ProgressItem(args[1])
	}
}
