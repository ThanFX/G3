package terrain

import (
	uuid "github.com/satori/go.uuid"
)

type Meadow struct {
	Terrain
}

type Meadows struct {
	Objects []Meadow
	InCh    chan string `json:"-"`
}

var M Meadows

func CreateMeadow(chunkId uuid.UUID, size int) uuid.UUID {
	m := Meadow{
		Terrain{uuid.Must(uuid.NewV4()),
			size,
			chunkId}}
	M.Objects = append(M.Objects, m)
	return m.ID
}

func GetMeadows() []Meadow {
	return M.Objects
}
