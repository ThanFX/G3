package models

import (
	"fmt"

	"github.com/ThanFX/G3/areas"

	uuid "github.com/satori/go.uuid"
)

var ChunkAreasInfo map[uuid.UUID]map[string]interface{}

func CreateTerrains() {
	ChunkAreasInfo = make(map[uuid.UUID]map[string]interface{})
	for _, m := range Map {
		ChunkAreasInfo[m.ID] = make(map[string]interface{})
		for _, t := range m.Terrains {
			//fmt.Println(t.Type)
			switch t.Type {
			case "forest":
				id := areas.CreateForest(m.ID, t.Size)
				ChunkAreasInfo[m.ID]["forest"] = areas.GetForestById(id)
			case "hill":
				id := areas.CreateHill(m.ID, t.Size)
				ChunkAreasInfo[m.ID]["hill"] = areas.GetHillsById(id)
			case "swamp":
				id := areas.CreateSwamp(m.ID, t.Size)
				ChunkAreasInfo[m.ID]["swamp"] = areas.GetSwampById(id)
			case "meadow":
				id := areas.CreateMeadow(m.ID, t.Size)
				ChunkAreasInfo[m.ID]["meadow"] = areas.GetMeadowById(id)
			case "lake":
				id := areas.CreateLake(m.ID, t.Size)
				ChunkAreasInfo[m.ID]["lake"] = areas.GetLakesById(id)
			}
		}
		var rs []areas.River
		for _, r := range m.Rivers {
			id := areas.CreateRiver(m.ID, r.Size, r.Bridge)
			rs = append(rs, areas.GetRiversById(id)[0])
		}
		ChunkAreasInfo[m.ID]["river"] = rs
	}
	//fmt.Println(ChunkAreasInfo)
}

func GetChunkTerrainsInfo(param string) map[string]interface{} {
	chunkId, err := uuid.FromString(param)
	if err != nil {
		fmt.Printf("При получении ID чанка %s произошла ошибка %s", param, err)
		return nil
	}
	return ChunkAreasInfo[chunkId]
}
