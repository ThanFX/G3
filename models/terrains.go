package models

import (
	"github.com/ThanFX/G3/terrain"
)

func CreateTerrains() {
	for _, m := range Map {
		for _, t := range m.Terrains {
			//fmt.Println(t.Type)
			switch t.Type {
			case "forest":
				terrain.CreateForest(m.ID, t.Size)
			case "hill":
				terrain.CreateHill(m.ID, t.Size)
			case "swamp":
				terrain.CreateSwamp(m.ID, t.Size)
			case "meadow":
				terrain.CreateMeadow(m.ID, t.Size)
			case "lake":
				terrain.CreateLake(m.ID, t.Size)
			}
		}
		for _, r := range m.Rivers {
			terrain.CreateRiver(m.ID, r.Size, r.Bridge)
		}
	}
}
