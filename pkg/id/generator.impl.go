package id

import (
	"time"

	"github.com/sony/sonyflake"
)

type generator struct {
	flake *sonyflake.Sonyflake
}

func NewGenerator() (Generator, error) {
	flake, err := sonyflake.New(sonyflake.Settings{
		StartTime: time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		return nil, err
	}
	return &generator{
		flake: flake,
	}, nil
}

func (g *generator) NewID() uint64 {
	id, _ := g.flake.NextID()
	return id
}
