package models

import (
	"fmt"

	"github.com/ThanFX/G3/terrain"
	uuid "github.com/satori/go.uuid"
)

var ChunkTerrains map[uuid.UUID]map[string][]uuid.UUID

func CreateTerrains() {
	ChunkTerrains = make(map[uuid.UUID]map[string][]uuid.UUID)
	for _, m := range Map {
		ChunkTerrains[m.ID] = make(map[string][]uuid.UUID)
		for _, t := range m.Terrains {
			//fmt.Println(t.Type)
			switch t.Type {
			case "forest":
				id := terrain.CreateForest(m.ID, t.Size)
				ChunkTerrains[m.ID]["forest"] = append(ChunkTerrains[m.ID]["forest"], id)
			case "hill":
				id := terrain.CreateHill(m.ID, t.Size)
				ChunkTerrains[m.ID]["hill"] = append(ChunkTerrains[m.ID]["hill"], id)
			case "swamp":
				id := terrain.CreateSwamp(m.ID, t.Size)
				ChunkTerrains[m.ID]["swamp"] = append(ChunkTerrains[m.ID]["swamp"], id)
			case "meadow":
				id := terrain.CreateMeadow(m.ID, t.Size)
				ChunkTerrains[m.ID]["meadow"] = append(ChunkTerrains[m.ID]["meadow"], id)
			case "lake":
				id := terrain.CreateLake(m.ID, t.Size)
				ChunkTerrains[m.ID]["lake"] = append(ChunkTerrains[m.ID]["lake"], id)
			}
		}
		for _, r := range m.Rivers {
			id := terrain.CreateRiver(m.ID, r.Size, r.Bridge)
			ChunkTerrains[m.ID]["river"] = append(ChunkTerrains[m.ID]["river"], id)
		}
	}
	fmt.Println(ChunkTerrains)
}

func GetChunkTerrainsInfo(param string) map[string][]uuid.UUID {
	chunkId, err := uuid.FromString(param)
	if err != nil {
		fmt.Printf("При получении ID чанка %s произошла ошибка %s", param, err)
		return nil
	}
	return ChunkTerrains[chunkId]
}
