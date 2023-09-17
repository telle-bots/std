package id

// Generator represents ID generator
type Generator interface {
	// NewID returns new uniq ID
	NewID() uint64
}
