package main

import "fmt"

func main() {
	component := &ConcreteComponent{}
	fmt.Println(component.Operation()) // component

	decorator := &ConcreteDecorator{component}
	fmt.Println(decorator.Operation()) // ***component***
}

// Component provides an interface for a decorator and component.
type Component interface {
	Operation() string
}

// ConcreteComponent implements a component.
type ConcreteComponent struct {
}

// Operation implementation.
func (c *ConcreteComponent) Operation() string {
	return "component"
}

// ConcreteDecorator implements a decorator.
type ConcreteDecorator struct {
	component Component
}

// Operation wraps operation of component
func (d *ConcreteDecorator) Operation() string {
	return "***" + d.component.Operation() + "***"
}
