package areas

import (
	"database/sql"
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
	L  Lakes
	DB *sql.DB
)

func LakesStart() {
	go L.lakesListener()
	//libs.ReadFishCatalog()
}

func LakesNextDate() {
	L.InCh <- "next"
}

func CreateLake(chunkId uuid.UUID, size int) uuid.UUID {
	cap, maxCap := libs.GetFishingInitSize(size)
	l := Lake{
		libs.Area{
			ID:      uuid.Must(uuid.NewV4()),
			Size:    size,
			ChunkID: chunkId,
			Masterships: []libs.AreaMastery{
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("hunting"),
					Capacity:    0,
					MaxCapacity: 0},
				libs.AreaMastery{
					Mastership:  libs.GetMasteryByName("fishing"),
					Capacity:    cap,
					MaxCapacity: maxCap}}}}
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
			//case "fishing":
			//	go l.calcFishingResult(params[1], params[2])
		}
	}
}

func (ls *Lakes) setDayInc() {
	for _, l := range ls.Objects {
		cap, maxCap := l.Area.GetLakeFishingCap()
		dayInc := int((maxCap - cap) / 100)
		if dayInc < 1 {
			dayInc = 1
		}
		cap += dayInc
		if cap > maxCap {
			dayInc -= (cap - maxCap)
			cap = maxCap
		}
		l.Area.SetLakeFishingCap(cap, maxCap)
		//NewEvent(
		//	fmt.Sprintf("В озере %s за день родилось %s рыбы. Всего сейчас %s рыбы.", strconv.Itoa(l.ID), strconv.Itoa(l.DayInc), strconv.Itoa(l.Capacity)))

	}
}
