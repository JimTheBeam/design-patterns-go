package main

import "fmt"

func main() {
	adapter := NewAdapter(&Adaptee{})

	req := adapter.Request()

	fmt.Printf("request: %s", req)
}

// Target provides an interface with which the system should work.
type Target interface {
	Request() string
}

// Adaptee implements system to be adapted.
type Adaptee struct{}

// NewAdapter is the Adapter constructor.
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee: adaptee}
}

// SpecificRequest implementation.
func (a *Adaptee) SpecificRequest() string {
	return "Specific Request"
}

// Adapter implements Target interface and is an adapter.
type Adapter struct {
	adaptee *Adaptee
}

// Request is an adaptive method.
func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}
