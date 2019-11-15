package areas

import (
	"strings"

	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type Swamp struct {
	libs.Area
}

type Swamps struct {
	Objects []Swamp
	InCh    chan string `json:"-"`
}

var S Swamps

func SwampsStart() {
	S.InCh = make(chan string, 0)
	go S.swampsListener()
}

func SwampsNextDate() {
	S.InCh <- "next"
}

func CreateSwamp(chunkId uuid.UUID, size int) uuid.UUID {
	fgcap, fgmaxCap := libs.GetFoodGatheringInitSize(size)
	s := Swamp{
		libs.Area{
			ID:      uuid.Must(uuid.NewV4()),
			Size:    size,
			ChunkID: chunkId,
			Masterships: []libs.AreaMastery{
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("food_gathering"),
					Capacity:    fgcap,
					MaxCapacity: fgmaxCap}}}}
	S.Objects = append(S.Objects, s)
	return s.ID
}

func GetSwamps() []Swamp {
	return S.Objects
}

func GetSwampById(id uuid.UUID) Swamp {
	var s Swamp
	for i := range S.Objects {
		if uuid.Equal(S.Objects[i].ID, id) {
			s = S.Objects[i]
			break
		}
	}
	return s
}

func (s *Swamps) swampsListener() {
	for {
		com := <-s.InCh
		params := strings.Split(com, "|")
		switch params[0] {
		case "next":
			go s.setDayInc()
		}
	}
}

func (ss *Swamps) setDayInc() {
	for _, s := range ss.Objects {
		s.Area.SetDayIncCapacity()
	}
}
