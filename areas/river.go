package areas

import (
	"strings"

	"github.com/ThanFX/G3/libs"
	uuid "github.com/satori/go.uuid"
)

type River struct {
	libs.Area
	IsBribge bool
}

type Rivers struct {
	Objects []River
	InCh    chan string `json:"-"`
}

var R Rivers

func RiversStart() {
	R.InCh = make(chan string, 0)
	go R.riversListener()
}

func RiversNextDate() {
	R.InCh <- "next"
}

func CreateRiver(chunkId uuid.UUID, size int, isBridge bool) uuid.UUID {
	cap, maxCap := libs.GetFishingInitSize(size)
	r := River{
		libs.Area{
			ID:      uuid.Must(uuid.NewV4()),
			Size:    size,
			ChunkID: chunkId,
			Masterships: []libs.AreaMastery{
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("fishing"),
					Capacity:    cap,
					MaxCapacity: maxCap}}},
		isBridge}
	R.Objects = append(R.Objects, r)
	return r.ID
}

func GetRivers() []River {
	return R.Objects
}

func GetRiversById(id uuid.UUID) []River {
	var r []River
	for i := range R.Objects {
		if uuid.Equal(R.Objects[i].ID, id) {
			r = append(r, R.Objects[i])
		}
	}
	return r
}

func (r *Rivers) riversListener() {
	for {
		com := <-r.InCh
		params := strings.Split(com, "|")
		switch params[0] {
		case "next":
			go r.setDayInc()
		}
	}
}

func (rs *Rivers) setDayInc() {
	for _, r := range rs.Objects {
		cap := r.Area.GetFishingCap()
		newCap := libs.GetFishingDayInc(cap, r.Size)
		r.Area.SetFishingCap(newCap)
	}
}
