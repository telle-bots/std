package id

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

// generator represents SonyFlake ID generator
type generator struct {
	flake *sonyflake.Sonyflake
}

// NewGenerator provides generator
func NewGenerator() (Generator, error) {
	flake, err := sonyflake.New(sonyflake.Settings{
		StartTime: time.Date(2023, 4, 28, 6, 17, 42, 23, time.UTC),
	})
	if err != nil {
		return nil, fmt.Errorf("id generator: %w", err)
	}
	return &generator{
		flake: flake,
	}, nil
}

// NewID generates new flake ID
func (g *generator) NewID() uint64 {
	id, _ := g.flake.NextID()
	return id
}
