package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type item struct {
	id   string
	name string
}

type store struct {
	items []item
}

func CreateStore() {
	file, err := os.Create("store.csv")
	if err != nil {
		fmt.Println("Error creating store file")
	}
	file.Close()
	fmt.Println("Store Created")
}

func (s *store) CreateTestStore() {
	items := []item{
		{id: Guid(), name: "item1"},
		{id: Guid(), name: "item2"},
	}
	s.items = items
	s.WriteStore()
}

func (s *store) WriteStore() {
	file, err := os.Open("store.csv")
	if err != nil {
		fmt.Println("Error opening store file")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range s.items {
		record := []string{item.id, item.name}
		fmt.Println(record)
		err := writer.Write(record)
		// this seems like a dangerous way to do this because if it errors then I'll
		// lose the data that was already written to the file. It's probably fine for
		// this use case but worth calling out. This issue would be easy to fix if
		// I was using a SQL database instead of a CSV file, so I'm going to ignore it
		if err != nil {
			fmt.Println("Error writing to store file")
		}
	}
	fmt.Println("Store Updated")
}

func GetStore() (*store, error) {
	// TODO:: Add error handling for a store that doesn't exist yet
	file, err := os.Open("store.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	items := []item{}
	for _, record := range records {
		items = append(items, item{id: record[0], name: record[1]})
	}

	return &store{items: items}, nil
}

func (s *store) ListItems() {
	for _, item := range s.items {
		fmt.Println(item)
	}
}

func (s *store) CreateItem(name string) {
	item := &item{id: Guid(), name: name}
	s.items = append(s.items, *item)
}
