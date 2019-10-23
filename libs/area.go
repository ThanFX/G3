package libs

import (
	uuid "github.com/satori/go.uuid"
)

type AreaMastery struct {
	Mastership  Mastery `json:"mastery"`
	Capacity    int     `json:"capacity"`
	MaxCapacity int     `json:"max_capacity"`
}

type Area struct {
	ID          uuid.UUID     `json:"id"`
	Size        int           `json:"size"`
	ChunkID     uuid.UUID     `json:"-"`
	Masterships []AreaMastery `json:"masterships"`
}
