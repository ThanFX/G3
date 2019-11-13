package areas

import (
	"strings"

	"github.com/ThanFX/G3/libs"

	uuid "github.com/satori/go.uuid"
)

type Forest struct {
	libs.Area
}

type Forests struct {
	Objects []Forest
	InCh    chan string `json:"-"`
}

var F Forests

func ForestsStart() {
	F.InCh = make(chan string, 0)
	go F.forestsListener()
}

func ForestsNextDate() {
	F.InCh <- "next"
}

func CreateForest(chunkId uuid.UUID, size int) uuid.UUID {
	cap, maxCap := libs.GetHuntingInitSize(size)
	f := Forest{
		libs.Area{
			ID:      uuid.Must(uuid.NewV4()),
			Size:    size,
			ChunkID: chunkId,
			Masterships: []libs.AreaMastery{
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("hunting"),
					Capacity:    cap,
					MaxCapacity: maxCap},
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("food_gathering"),
					Capacity:    0,
					MaxCapacity: 0}}}}
	F.Objects = append(F.Objects, f)
	return f.ID
}

func GetForests() []Forest {
	return F.Objects
}

func GetForestById(id uuid.UUID) Forest {
	var f Forest
	for i := range F.Objects {
		if uuid.Equal(F.Objects[i].ID, id) {
			f = F.Objects[i]
			break
		}
	}
	return f
}

func (f *Forests) forestsListener() {
	for {
		com := <-f.InCh
		params := strings.Split(com, "|")
		switch params[0] {
		case "next":
			go f.setDayInc()
		}
	}
}

func (fs *Forests) setDayInc() {
	for _, f := range fs.Objects {
		cap := f.Area.GetHuntingCap()
		newCap := libs.GetHuntingDayInc(cap, f.Size)
		f.Area.SetHuntingCap(newCap)
	}
}
