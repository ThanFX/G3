package libs

import (
	uuid "github.com/satori/go.uuid"
)

type AreaMastery struct {
	Mastership  Mastery `json:"mastery"`
	Capacity    int     `json:"capacity"`
	MaxCapacity int     `json:"max_capacity"`
}

type Area struct {
	ID          uuid.UUID     `json:"id"`
	Size        int           `json:"size"`
	ChunkID     uuid.UUID     `json:"-"`
	Masterships []AreaMastery `json:"masterships"`
}

func (a *Area) getMasteryByName(mastery string) AreaMastery {
	var am AreaMastery
	for _, m := range a.Masterships {
		if m.Mastership.Name == mastery {
			am = m
			break
		}
	}
	return am
}

func (a *Area) GetLakeFishingCap() (int, int) {
	am := a.getMasteryByName("fishing")
	return am.Capacity, am.MaxCapacity
}

func (a *Area) SetLakeFishingCap(cap, maxCap int) {
	am := a.getMasteryByName("fishing")
	am.Capacity = cap
	am.MaxCapacity = maxCap
}
