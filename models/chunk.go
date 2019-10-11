package models

import (
	uuid "github.com/satori/go.uuid"
)

type Chunk struct {
	ID       uuid.UUID      `json:"id"`
	X        int            `json:"x"`
	Y        int            `json:"y"`
	Terrains []TerrainChunk `json:"terrains"`
	Rivers   []RiverChunk   `json:"rivers"`
}

type TerrainChunk struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

type RiverChunk struct {
	Size   int    `json:"size"`
	From   string `json:"from"`
	To     string `json:"to"`
	Bridge bool   `json:"bridge"`
}
