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

func DeleteStore() {
	file, err := os.OpenFile("store.csv", os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error deleting store file")
	}
	defer file.Close()
	file.Truncate(0)
}

func WriteToFile(data [][]string) {
	file, err := os.OpenFile("store.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening store file")
	}
	defer file.Close()
	file.Truncate(0)
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(data)
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
		status, _ := StringToStatus(record[4])
		items = append(items, item{id: record[0], createdAt: record[1], lastUpdated: record[2], name: record[3], status: status, description: record[5]})
	}

	return &store{items: items}, nil
}
