package areas

import (
	"strings"

	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type Meadow struct {
	libs.Area
}

type Meadows struct {
	Objects []Meadow
	InCh    chan string `json:"-"`
}

var M Meadows

func MeadowsStart() {
	M.InCh = make(chan string, 0)
	go M.meadowsListener()
}

func MeadowsNextDate() {
	M.InCh <- "next"
}

func CreateMeadow(chunkId uuid.UUID, size int) uuid.UUID {
	hcap, hmaxCap := libs.GetHuntingInitSize(size)
	fgcap, fgmaxCap := libs.GetFoodGatheringInitSize(size)
	m := Meadow{
		libs.Area{
			ID:      uuid.Must(uuid.NewV4()),
			Size:    size,
			ChunkID: chunkId,
			Masterships: []libs.AreaMastery{
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("hunting"),
					Capacity:    hcap,
					MaxCapacity: hmaxCap},
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("food_gathering"),
					Capacity:    fgcap,
					MaxCapacity: fgmaxCap}}}}
	M.Objects = append(M.Objects, m)
	return m.ID
}

func GetMeadows() []Meadow {
	return M.Objects
}

func GetMeadowById(id uuid.UUID) Meadow {
	var m Meadow
	for i := range M.Objects {
		if uuid.Equal(M.Objects[i].ID, id) {
			m = M.Objects[i]
			break
		}
	}
	return m
}

func (m *Meadows) meadowsListener() {
	for {
		com := <-m.InCh
		params := strings.Split(com, "|")
		switch params[0] {
		case "next":
			go m.setDayInc()
		}
	}
}

func (ms *Meadows) setDayInc() {
	for _, m := range ms.Objects {
		m.Area.SetDayIncCapacity()
	}
}
