package main

import "fmt"

func main() {
	product := NewConcreteProduct("A")
	cloneProduct := product.Clone()

	fmt.Printf("product: %s, clone: %s\n", product.GetName(), cloneProduct.GetName())
}

// Prototyper provides a cloning interface.
type Prototyper interface {
	Clone() Prototyper
	GetName() string
}

// ConcreteProduct implements product "A"
type ConcreteProduct struct {
	name string // product name
}

// NewConcreteProduct is the Prototyper constructor.
func NewConcreteProduct(name string) Prototyper {
	return &ConcreteProduct{
		name: name,
	}
}

// GetName returns product name
func (p *ConcreteProduct) GetName() string {
	return p.name
}

// Clone returns a cloned object.
func (p *ConcreteProduct) Clone() Prototyper {
	return &ConcreteProduct{p.name}
}
