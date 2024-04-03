package main

import (
	"fmt"
)

type item struct {
	id          string
	createdAt   string
	lastUpdated string
	name        string
	status      Status
	description string
}

func (i *item) ProgressStatus() {
	switch i.status {
	case Backlog:
		i.status = Planned
	case Planned:
		i.status = InProgress
	case InProgress:
		i.status = InReview
	case InReview:
		i.status = Done
	case Done:
		i.status = Archived
	case Archived:
		i.status = Archived
	}
}

type store struct {
	items []item
}

func (s *store) CreateTestStore() {
	DeleteStore()
	for i := 0; i < 10; i++ {
		s.CreateItem(fmt.Sprintf("item-%d", i), fmt.Sprintf("description-%d", i))
	}
}

func (s *store) WriteStore() {
	WriteToFile(s.convertItemsToRecords())
}

func (s *store) ListItems() {
	fmt.Println("Id                                   | Created At              | Last Updated            | Name | Status | Description")
	for _, item := range s.items {
		fmt.Println(item)
	}
}

func (s *store) GetItemByName(name string) {
	for _, item := range s.items {
		if item.name == name {
			fmt.Println(item)
			return
		}
	}
	fmt.Println("Item not found")
}

func (s *store) CreateItem(name string, description string) {
	now := GetCurrentTimeString()
	fmt.Println(now)
	item := &item{
		id:          Guid(),
		createdAt:   now,
		lastUpdated: now,
		name:        name,
		status:      Backlog,
		description: description,
	}
	s.items = append(s.items, *item)
	fmt.Println(s.items)
	s.WriteStore()
}

func (s *store) ProgressItem(name string) {
	// This isn't working yet
	for i, item := range s.items {
		if item.name == name {
			s.items[i].ProgressStatus()
			s.items[i].lastUpdated = GetCurrentTimeString()
			s.WriteStore()
			return
		}
	}
	fmt.Println("Item not found")
}

func (s *store) convertItemsToRecords() [][]string {
	records := [][]string{}

	for _, item := range s.items {
		record := []string{item.id, item.createdAt, item.lastUpdated, item.name, string(item.status), item.description}
		records = append(records, record)
	}
	return records
}

func (s *store) DeleteItem(name string) {
	for i, item := range s.items {
		if item.name == name {
			s.items = append(s.items[:i], s.items[i+1:]...)
			s.WriteStore()
			return
		}
	}
	fmt.Println("Item not found")
}

func (s *store) UpdateItem(name string, field string, value string) {
	for i, item := range s.items {
		if item.name == name {
			switch field {
			case "name":
				s.items[i].name = value
			case "status":
				s.items[i].status = ToStatus(value)
			case "description":
				s.items[i].description = value
			default:
				fmt.Println("Field not found or not updateable")
				return
			}
			s.items[i].lastUpdated = GetCurrentTimeString()
			s.WriteStore()
			return
		}
	}
	fmt.Println("Item not found")
}
