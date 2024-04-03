package main

import (
	"flag"
	"fmt"
)

func main() {

	// Flags must go before any commands

	// Flag for "all items"
	var all bool
	flag.BoolVar(&all, "a", false, "all items (shortcut)")
	flag.BoolVar(&all, "all", false, "all items")

	// Flag for "item name"
	var name string
	flag.StringVar(&name, "n", "", "item name (shortcut)")
	flag.StringVar(&name, "name", "", "item name")

	// Flag for "item id"
	var id string
	flag.StringVar(&id, "id", "", "item id")

	// Flag for "item description"
	var description string
	flag.StringVar(&description, "d", "", "item description (shortcut)")
	flag.StringVar(&description, "description", "", "item description")

	// Flag for "item status"
	var status string
	flag.StringVar(&status, "s", "", "item status (shortcut)")
	flag.StringVar(&status, "status", "", "item status")

	// Flag for "allow empty"
	var allowEmpty bool
	flag.BoolVar(&allowEmpty, "e", false, "allow empty (shortcut)")
	flag.BoolVar(&allowEmpty, "allow-empty", false, "allow empty values to be passed in for the item name or description when updating an item")

	// Flag for "showing archived items"
	var showArchived bool
	flag.BoolVar(&showArchived, "include-archived", false, "show archived items")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = append(args, "help")
	}

	// TODO::
	// - [x] Extract flags from args
	// - [x] Add help command
	// - [x] Add list command
	// - [x] Add create command
	// - [x] Add update command
	// - [x] Add delete command
	// - [x] Add some sort of data store. Probably just going to use a static CSV file for now.
	// - [x] Add todo list-type functionality
	// - [] Bulk actions

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

	cmd, err := StringToCommand(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	switch cmd {
	case Help:
		var cmds []Command
		if len(args) == 1 {
			cmds = []Command{Help, List, CreateTestStore, Create, Delete, Update, Progress, Archive, ChangeStatus}
		} else {
			for _, arg := range args[1:] {
				cmd, err := StringToCommand(arg)
				if err != nil {
					fmt.Println(err)
					return
				}
				cmds = append(cmds, cmd)
			}
		}

		for _, cmd := range cmds {
			fmt.Println(GetUsage(cmd))
		}
	case List:
		if name != "" {
			item, err := store.GetItemByName(name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(item)
			}
		} else if id != "" {
			item, err := store.GetItemById(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(item)
			}
		} else if all || len(args) == 1 {
			store.ListItems(showArchived)
		} else {
			fmt.Println("Too many arguments for list command")
		}
	case CreateTestStore:
		store.CreateTestStore()
	case Create:
		if name == "" {
			fmt.Println("Please provide a name for the item to create")
			return
		}
		store.CreateItem(name, description, status)
	case Delete:
		if name == "" && id == "" {
			fmt.Println("Please provide a name or id for the item to delete")
			return
		}
		store.DeleteItem(name, id)
	case Update:
		if name == "" && id == "" {
			fmt.Println("Please provide a name or id for the item to update. If you'd like to update the name, you must provide the id.")
			return
		}
		store.UpdateItem(id, name, description, status, allowEmpty)
	case Progress:
		if name == "" && id == "" {
			fmt.Println("Please provide a name or id for the item to progress")
			return
		}
		store.ProgressItem(id, name)
	case Archive:
		store.UpdateItem(id, name, "", "archived", false)
	case ChangeStatus:
		store.UpdateItem(id, name, "", status, false)
	default:
		fmt.Println("Invalid command: " + args[0] + ". Use 'help' to see available commands.")
	}
}
