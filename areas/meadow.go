package areas

import (
	uuid "github.com/satori/go.uuid"
)

type Meadow struct {
	Area
}

type Meadows struct {
	Objects []Meadow
	InCh    chan string `json:"-"`
}

var M Meadows

func CreateMeadow(chunkId uuid.UUID, size int) uuid.UUID {
	m := Meadow{
		Area{uuid.Must(uuid.NewV4()),
			size,
			chunkId,
			[]AreaMastery{
				AreaMastery{GetMasteryByName("hunting"), 0, 0},
				AreaMastery{GetMasteryByName("food_gathering"), 0, 0}}}}
	M.Objects = append(M.Objects, m)
	return m.ID
}

func GetMeadows() []Meadow {
	return M.Objects
}

func GetMeadowById(id uuid.UUID) []Meadow {
	var m []Meadow
	for i := range M.Objects {
		if uuid.Equal(M.Objects[i].ID, id) {
			m = append(m, M.Objects[i])
		}
	}
	return m
}
