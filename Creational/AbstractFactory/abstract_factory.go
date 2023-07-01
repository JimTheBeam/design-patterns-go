package main

// AbstractFactory provides an interface for creating families of related objects.
type AbstractFactory interface {
	CreateSnowboard() AbstractSnowboard
	CreateBoots() AbstractBoots
}

// AbstractSnowboard provides snowboard interface.
type AbstractSnowboard interface {
	GetLength() int
}

// AbstractBoots provides boots interface.
type AbstractBoots interface {
	GetSize() int
}
