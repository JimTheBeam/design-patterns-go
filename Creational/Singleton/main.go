package main

import (
	"fmt"
)

var instance *Singleton

func main() {
	instance1 := GetInstance()
	instance2 := GetInstance()

	fmt.Printf("is objects equal: %t\n", instance1 == instance2)
}

// Singleton implementation.
type Singleton struct {
	data string
}

// GetInstance returns singleton
func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{data: "singleton"}
	}

	return instance
}
