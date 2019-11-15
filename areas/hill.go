package areas

import (
	"strings"

	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type Hill struct {
	libs.Area
}

type Hills struct {
	Objects []Hill
	InCh    chan string `json:"-"`
}

var H Hills

func HillsStart() {
	H.InCh = make(chan string, 0)
	go H.hillsListener()
}

func HillsNextDate() {
	H.InCh <- "next"
}

func CreateHill(chunkId uuid.UUID, size int) uuid.UUID {
	hcap, hmaxCap := libs.GetHuntingInitSize(size)
	fgcap, fgmaxCap := libs.GetFoodGatheringInitSize(size)
	h := Hill{
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
	H.Objects = append(H.Objects, h)
	return h.ID
}

func GetHills() []Hill {
	return H.Objects
}

func GetHillsById(id uuid.UUID) Hill {
	var h Hill
	for i := range H.Objects {
		if uuid.Equal(H.Objects[i].ID, id) {
			h = H.Objects[i]
			break
		}
	}
	return h
}

func (h *Hills) hillsListener() {
	for {
		com := <-h.InCh
		params := strings.Split(com, "|")
		switch params[0] {
		case "next":
			go h.setDayInc()
		}
	}
}

func (hs *Hills) setDayInc() {
	for _, h := range hs.Objects {
		h.Area.SetDayIncCapacity()
	}
}
