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
		s.CreateItem(fmt.Sprintf("item-%d", i), fmt.Sprintf("description-%d", i), "")
	}
}

func (s *store) WriteStore() {
	WriteToFile(s.convertItemsToRecords())
}

func (s *store) ListItems(showArchived bool) {
	fmt.Println("Id                                   | Created At              | Last Updated            | Name | Status | Description")
	for _, item := range s.items {
		if item.status != Archived || showArchived {
			fmt.Println(item)
		}
	}
}

func (s *store) GetItemByName(name string) (item, error) {
	for _, item := range s.items {
		if item.name == name {
			return item, nil
		}
	}
	return item{}, fmt.Errorf("Item not found")
}

func (s *store) GetItemById(id string) (item, error) {
	for _, item := range s.items {
		if item.id == id {
			return item, nil
		}
	}
	return item{}, fmt.Errorf("Item not found")
}

func (s *store) CreateItem(name string, description string, status string) error {
	// Currently not forcing uniqueness on names... This is going to be a problem
	// but I'll punt it for now
	now := GetCurrentTimeString()
	if status == "" {
		status = "backlog"
	}
	typedStatus, err := StringToStatus(status)
	if err != nil {
		return err
	}
	item := &item{
		id:          Guid(),
		createdAt:   now,
		lastUpdated: now,
		name:        name,
		status:      typedStatus,
		description: description,
	}
	s.items = append(s.items, *item)
	s.WriteStore()
	return nil
}

func (s *store) ProgressItem(id string, name string) error {
	for i, item := range s.items {
		if item.name == name && name != "" || item.id == id && id != "" {
			s.items[i].ProgressStatus()
			s.items[i].lastUpdated = GetCurrentTimeString()
			s.WriteStore()
			return nil
		}
	}
	return fmt.Errorf("Item not found")
}

func (s *store) convertItemsToRecords() [][]string {
	records := [][]string{}

	for _, item := range s.items {
		record := []string{item.id, item.createdAt, item.lastUpdated, item.name, string(item.status), item.description}
		records = append(records, record)
	}
	return records
}

func (s *store) DeleteItem(name string, id string) error {
	for i, item := range s.items {
		if item.name == name && name != "" || item.id == id && id != "" {
			s.items = append(s.items[:i], s.items[i+1:]...)
			s.WriteStore()
			return nil
		}
	}
	return fmt.Errorf("Item not found")
}

func (s *store) UpdateItem(id string, name string, description string, status string, allowEmpty bool) error {
	var itemToUpdate item
	if id == "" {
		itemToUpdate, _ = s.GetItemByName(name)
	} else {
		itemToUpdate, _ = s.GetItemById(id)
	}
	for i, item := range s.items {
		if item.id == itemToUpdate.id {
			if name != "" || allowEmpty {
				s.items[i].name = name
			}
			if description != "" || allowEmpty {
				s.items[i].description = description
			}
			if status != "" {
				status, err := StringToStatus(status)
				if err != nil {
					return err
				}
				s.items[i].status = status
			}
			s.items[i].lastUpdated = GetCurrentTimeString()
			s.WriteStore()
			return nil
		}
	}
	return fmt.Errorf("Item not found")
}
