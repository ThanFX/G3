package areas

import (
	"strings"

	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type Lake struct {
	libs.Area
}

type Lakes struct {
	Objects []Lake
	InCh    chan string `json:"-"`
}

var (
	L Lakes
)

func LakesStart() {
	L.InCh = make(chan string, 0)
	go L.lakesListener()
}

func LakesNextDate() {
	L.InCh <- "next"
}

func CreateLake(chunkId uuid.UUID, size int) uuid.UUID {
	hcap, hmaxCap := libs.GetHuntingInitSize(size)
	fcap, fmaxCap := libs.GetFishingInitSize(size)
	l := Lake{
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
					Mastership:  libs.GetMasteryByName("fishing"),
					Capacity:    fcap,
					MaxCapacity: fmaxCap}}}}
	L.Objects = append(L.Objects, l)
	return l.ID
}

func GetLakes() []Lake {
	return L.Objects
}

func GetLakesById(id uuid.UUID) []Lake {
	var l []Lake
	for i := range L.Objects {
		if uuid.Equal(L.Objects[i].ID, id) {
			l = append(l, L.Objects[i])
		}
	}
	return l
}

func (l *Lakes) lakesListener() {
	for {
		com := <-l.InCh
		params := strings.Split(com, "|")
		switch params[0] {
		case "next":
			go l.setDayInc()
		}
	}
}

func (ls *Lakes) setDayInc() {
	for _, l := range ls.Objects {
		l.Area.SetDayIncCapacity()
	}
}
