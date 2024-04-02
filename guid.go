package main

import (
	uuid "github.com/google/uuid"
)

func Guid() string {
	return uuid.New().String()
}
