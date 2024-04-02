package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

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
	WriteToFile(s.convertItemsToRecords())
}

func WriteToFile(data [][]string) {
	file, err := os.OpenFile("store.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening store file")
	}
	defer file.Close()
	file.WriteString("")
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(data)

	// for _, record := range data {
	// 	err := writer.Write(record)
	// 	if err != nil {
	// 		fmt.Println("Error writing to store file")
	// 	}
	// }
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
