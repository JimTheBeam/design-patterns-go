package main

// BurtonFactory implements AbstractFactory interface.
type BurtonFactory struct {
}

// NewBurtonFactory is the BurtonFactory constructor.
func NewBurtonFactory() AbstractFactory {
	return &BurtonFactory{}
}

// CreateSnowboard implementation.
func (f *BurtonFactory) CreateSnowboard() AbstractSnowboard {
	return &Burton{length: 160}
}

// CreateBoots implementation.
func (f *BurtonFactory) CreateBoots() AbstractBoots {
	return &JonesBoots{size: 8}
}

// Burton implements AbstractSnowboard.
type Burton struct {
	length int
}

// GetLength returns length of snowboard.
func (s *Burton) GetLength() int {
	return s.length
}

// BurtonBoots implements AbstractBoots.
type BurtonBoots struct {
	size int
}

// GetSize returns size of the boots.
func (b *BurtonBoots) GetSize() int {
	return b.size
}
