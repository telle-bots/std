package id

type Generator interface {
	NewID() uint64
}
