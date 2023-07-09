package main

import (
	"fmt"
)

func main() {
	var files = []struct {
		filename string
		width    int
		height   int
		opacity  float64
	}{
		{"cat.jpg", 100, 100, 0.95},
		{"cat.jpg", 200, 200, 0.75},
		{"dog.jpg", 300, 300, 0.50},
	}

	var factory = new(FlyweightFactory)

	for _, f := range files {
		flyweight := factory.GetFlyweight(f.filename)
		result := flyweight.Draw(f.width, f.height, f.opacity)
		fmt.Println(result)
	}
}

// Flyweighter interface
type Flyweighter interface {
	Draw(width, height int, opacity float64) string
}

// FlyweightFactory implements a factory.
type FlyweightFactory struct {
	pool map[string]Flyweighter
}

// GetFlyweight creates or returns a suitable Flyweighter by state.
// If a suitable flyweighter is in pool, then returns it.
func (f *FlyweightFactory) GetFlyweight(filename string) Flyweighter {
	if f.pool == nil {
		f.pool = make(map[string]Flyweighter)
	}
	if _, ok := f.pool[filename]; !ok {
		f.pool[filename] = &ConcreteFlyweight{filename: filename}
	}
	return f.pool[filename]
}

// ConcreteFlyweight implements a Flyweighter interface.
type ConcreteFlyweight struct {
	filename string // internal state
}

// Draw draws image. Args width, height and opacity is external state.
func (f *ConcreteFlyweight) Draw(width, height int, opacity float64) string {
	return fmt.Sprintf("draw image: %s, width: %d, height: %d, opacity: %.2f", f.filename, width, height, opacity)
}
