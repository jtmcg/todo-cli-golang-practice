package main

type Status string

const (
	Backlog    Status = "backlog"
	Planned    Status = "planned"
	InProgress Status = "in-progress"
	InReview   Status = "in-review"
	Done       Status = "done"
	Archived   Status = "archived"
)

func ToStatus(status string) Status {
	switch status {
	case "backlog":
		return Backlog
	case "planned":
		return Planned
	case "in-progress":
		return InProgress
	case "in-review":
		return InReview
	case "done":
		return Done
	case "archived":
		return Archived
	default:
		return Backlog
	}
}
