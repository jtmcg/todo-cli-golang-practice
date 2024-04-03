package main

import "errors"

type Status string

const (
	Backlog    Status = "backlog"
	Planned    Status = "planned"
	InProgress Status = "in-progress"
	InReview   Status = "in-review"
	Done       Status = "done"
	Archived   Status = "archived"
)

func StringToStatus(status string) (Status, error) {
	switch status {
	case "backlog":
		return Backlog, nil
	case "planned":
		return Planned, nil
	case "in-progress":
		return InProgress, nil
	case "in-review":
		return InReview, nil
	case "done":
		return Done, nil
	case "archived":
		return Archived, nil
	default:
		err := errors.New("Invalid status: " + status)
		return "", err
	}
}
