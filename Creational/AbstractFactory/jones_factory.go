package main

// JonesFactory implements AbstractFactory interface.
type JonesFactory struct{}

// NewJonesFactory is the JonesFactory constructor.
func NewJonesFactory() AbstractFactory {
	return &JonesFactory{}
}

// CreateSnowboard implementation.
func (f *JonesFactory) CreateSnowboard() AbstractSnowboard {
	return &JonesSnowboard{length: 150}
}

// CreateBoots implementation.
func (f *JonesFactory) CreateBoots() AbstractBoots {
	return &JonesBoots{size: 9}
}

// JonesSnowboard implements AbstractSnowboard.
type JonesSnowboard struct {
	length int
}

// GetLength returns length of snowboard.
func (s *JonesSnowboard) GetLength() int {
	return s.length
}

// JonesBoots implements AbstractBoots.
type JonesBoots struct {
	size int
}

// GetSize returns size of the boots.
func (b *JonesBoots) GetSize() int {
	return b.size
}
