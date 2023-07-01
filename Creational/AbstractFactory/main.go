package main

import (
	"fmt"
)

func main() {
	jonesFactory := NewJonesFactory()
	printFactoryProductsInfo(jonesFactory)

	burtonFactory := NewBurtonFactory()
	printFactoryProductsInfo(burtonFactory)
}

// printFactoryProductsInfo - client, uses a factory.
func printFactoryProductsInfo(factory AbstractFactory) {
	snowboard := factory.CreateSnowboard()
	boots := factory.CreateBoots()

	fmt.Printf("Snowboard lenght: %d; Boots size: %d\n", snowboard.GetLength(), boots.GetSize())
}
