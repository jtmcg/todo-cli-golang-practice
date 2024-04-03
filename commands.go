package main

import (
	"errors"
	"fmt"
)

type Command string

const (
	Help            Command = "help"
	List            Command = "list"
	CreateTestStore Command = "create-test-store"
	Create          Command = "create"
	Delete          Command = "delete"
	Update          Command = "update"
	Progress        Command = "progress"
	ChangeStatus    Command = "change-status"
	Archive         Command = "archive"
)

func StringToCommand(command string) (Command, error) {
	switch command {
	case "help", "h":
		return Help, nil
	case "list", "ls":
		return List, nil
	case "create-test-store":
		return CreateTestStore, nil
	case "create", "create-item":
		return Create, nil
	case "delete", "delete-item", "del":
		return Delete, nil
	case "update", "update-item":
		return Update, nil
	case "progress":
		return Progress, nil
	case "change-status":
		return ChangeStatus, nil
	case "archive":
		return Archive, nil
	default:
		err := errors.New("Invalid command: " + command + ". Use 'help' to see available commands")
		return "", err
	}
}

type Usage struct {
	aliases        string
	description    string
	flags          string
	requiredFlags  string
	optionalFlags  string
	additionalInfo string
}

func GetUsage(command Command) string {
	var usage Usage
	switch command {
	case Help:
		usage = Usage{
			aliases:        "help, h",
			description:    "list available commands and their usage",
			requiredFlags:  "",
			optionalFlags:  "",
			additionalInfo: "you may pass in any commands as arguments to see their usage. No commands shows all commands' usage.",
		}
	case List:
		usage = Usage{
			aliases:        "list, ls",
			description:    "list item(s) in the todo list",
			requiredFlags:  "",
			optionalFlags:  "--all, --name, --id, --include-archived",
			additionalInfo: "Unless accessed by name or id, archived items will not be shown without the --include-archived flag",
		}
	case CreateTestStore:
		usage = Usage{
			aliases:        "create-test-store",
			description:    "create a test store with 10 items",
			requiredFlags:  "",
			optionalFlags:  "",
			additionalInfo: "",
		}
	case Create:
		usage = Usage{
			aliases:        "create, create-item",
			description:    "create an item",
			requiredFlags:  "--name",
			optionalFlags:  "--description, --status",
			additionalInfo: "The description may be empty. If no status is provided, the default status is 'backlog'",
		}
	case Delete:
		usage = Usage{
			aliases:        "delete, delete-item, del",
			description:    "delete an item",
			requiredFlags:  "--name or --id",
			optionalFlags:  "",
			additionalInfo: "",
		}
	case Update:
		usage = Usage{
			aliases:        "update, update-item",
			description:    "update an item",
			requiredFlags:  "--name or --id",
			optionalFlags:  "--description, --status, --allow-empty",
			additionalInfo: "If you'd like to update the name, you must provide the id. If no status is provided, the status will not be updated. If --allow-empty is provided, the name and description may be empty",
		}
	case Progress:
		usage = Usage{
			aliases:        "progress",
			description:    "progress an item",
			requiredFlags:  "--name or --id",
			optionalFlags:  "",
			additionalInfo: "Status progression: backlog -> planned -> in-progress -> in-review -> done -> archived",
		}
	case ChangeStatus:
		usage = Usage{
			aliases:        "change-status",
			description:    "change the status of an item",
			requiredFlags:  "--name or --id, --status",
			optionalFlags:  "",
			additionalInfo: "Status options: backlog, planned, in-progress, in-review, done, archived",
		}
	case Archive:
		usage = Usage{
			aliases:        "archive",
			description:    "archive an item",
			requiredFlags:  "--name or --id",
			optionalFlags:  "",
			additionalInfo: "Archived items will not show up in the list command without the --include-archived flag",
		}
	}
	return fmt.Sprintf("Command: %s\nDescription: %s\n\tRequired Flags: %s\n\tOptional Flags: %s\n\tAdditional Info: %s\n",
		usage.aliases, usage.description, usage.requiredFlags, usage.optionalFlags, usage.additionalInfo)
}
