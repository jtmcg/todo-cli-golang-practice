package main

import (
	"fmt"
)

type item struct {
	id   string
	name string
}

type store struct {
	items []item
}

func (s *store) WriteStore() {
	WriteToFile(s.convertItemsToRecords())
}

func (s *store) ListItems() {
	for _, item := range s.items {
		fmt.Println(item)
	}
}

func (s *store) CreateItem(name string) {
	item := &item{id: Guid(), name: name}
	s.items = append(s.items, *item)
	s.WriteStore()
}

func (s *store) convertItemsToRecords() [][]string {
	records := [][]string{}

	for _, item := range s.items {
		record := []string{item.id, item.name}
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
