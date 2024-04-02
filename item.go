package main

import "fmt"

type Item struct {
	name string
}

func (i *Item) Create(name string) {
	i.name = name
	fmt.Println("Item created", i.name)
}
