package terrain

import (
	uuid "github.com/satori/go.uuid"
)

type Forest struct {
	Terrain
}

type Forests struct {
	Objects []Forest
	InCh    chan string `json:"-"`
}

var F Forests

func CreateForest(chunkId uuid.UUID, size int) {
	//fmt.Println(chunkId)
	f := Forest{
		Terrain{uuid.Must(uuid.NewV4()),
			size,
			chunkId}}
	F.Objects = append(F.Objects, f)
}

func GetForests() []Forest {
	return F.Objects
}
