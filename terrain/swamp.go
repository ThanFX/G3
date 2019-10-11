package terrain

import (
	uuid "github.com/satori/go.uuid"
)

type Swamp struct {
	Terrain
}

type Swamps struct {
	Objects []Swamp
	InCh    chan string `json:"-"`
}

var S Swamps

func CreateSwamp(chunkId uuid.UUID, size int) {
	s := Swamp{
		Terrain{uuid.Must(uuid.NewV4()),
			size,
			chunkId}}
	S.Objects = append(S.Objects, s)
}

func GetSwamps() []Swamp {
	return S.Objects
}