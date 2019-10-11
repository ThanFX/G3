package terrain

import (
	uuid "github.com/satori/go.uuid"
)

type Terrain struct {
	ID       uuid.UUID
	Size     int
	ChunckID uuid.UUID
}
