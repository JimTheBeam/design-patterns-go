package main

import (
	"fmt"
	"strings"
)

func main() {
	product := new(Product)

	director := Director{&ConcreteBuilder{product}}
	director.Construct()

	// pizza with mozzarella, pepperoni and tomato sauce on thick dough
	fmt.Println(product.Bake())
}

// Builder provides a builder interface.
type Builder interface {
	AddDough(dough string)
	AddSauce(sauce string)
	AddFiling(filings ...string)
}

// Director implements a manager.
type Director struct {
	builder Builder
}

// Construct tells the builder what to do and in what order.
func (d *Director) Construct() {
	d.builder.AddDough("thick")
	d.builder.AddSauce("tomato")
	d.builder.AddFiling("mozzarella", "pepperoni")
}

// ConcreteBuilder implements Builder interface.
type ConcreteBuilder struct {
	product *Product
}

// AddDough builds dough of the pizza.
func (b *ConcreteBuilder) AddDough(dough string) {
	b.product.dough = dough
}

// AddSauce builds sauce of the pizza.
func (b *ConcreteBuilder) AddSauce(sauce string) {
	b.product.sauce = sauce
}

// AddFiling builds fillings of the pizza.
func (b *ConcreteBuilder) AddFiling(filings ...string) {
	b.product.fillings = filings
}

// Product implementation.
type Product struct {
	dough    string
	sauce    string
	fillings []string
}

// Bake returns product.
func (p *Product) Bake() string {
	return fmt.Sprintf(
		"pizza with %s and %s sauce on %s dough\n",
		strings.Join(p.fillings, ", "),
		p.sauce,
		p.dough,
	)
}
