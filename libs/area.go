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

func (a *Area) getMasteryByName(mastery string) *AreaMastery {
	var am *AreaMastery
	for _, m := range a.Masterships {
		if m.Mastership.NameID == mastery {
			am = &m
			break
		}
	}
	return am
}

func (a *Area) GetFishingCap() int {
	am := a.getMasteryByName("fishing")
	return am.Capacity
}

func (a *Area) GetHuntingCap() int {
	am := a.getMasteryByName("hunting")
	return am.Capacity
}

func (a *Area) GetFoodGatheringCap() int {
	am := a.getMasteryByName("food_gathering")
	return am.Capacity
}

func (a *Area) SetFishingCap(cap int) {
	for i, m := range a.Masterships {
		if m.Mastership.NameID == "fishing" {
			a.Masterships[i].Capacity = cap
			break
		}
	}
}

func (a *Area) SetHuntingCap(cap int) {
	for i, m := range a.Masterships {
		if m.Mastership.NameID == "hunting" {
			a.Masterships[i].Capacity = cap
			break
		}
	}
}

func (a *Area) SetDayIncCapacity() {
	for i, m := range a.Masterships {
		switch m.Mastership.NameID {
		case "fishing":
			cap := a.GetFishingCap()
			a.Masterships[i].Capacity = GetFishingDayInc(cap, a.Size)
		case "hunting":
			cap := a.GetHuntingCap()
			a.Masterships[i].Capacity = GetHuntingDayInc(cap, a.Size)
		case "food_gathering":
			cap := a.GetFoodGatheringCap()
			a.Masterships[i].Capacity = GetFoodGatheringDayInc(cap, a.Size)
		}
	}
}
