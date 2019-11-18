package models

import (
	"fmt"

	"github.com/ThanFX/G3/areas"

	uuid "github.com/satori/go.uuid"
)

type AreaMasrery struct {
	Name   string
	AreaID uuid.UUID
	Size   int
}

type AreaInfo struct {
	Forest areas.Forest
	Hill   areas.Hill
	Swamp  areas.Swamp
	Meadow areas.Meadow
	Lake   areas.Lake
	Rivers []areas.River
}

var (
	ChunkMasteryInfo map[uuid.UUID]map[string]interface{}
	ChunkAreasInfo   map[uuid.UUID]map[string]interface{}
	ChunkAreasInfoEx map[uuid.UUID]AreaInfo
)

func CreateTerrains() {
	ChunkMasteryInfo = make(map[uuid.UUID]map[string]interface{})
	ChunkAreasInfo = make(map[uuid.UUID]map[string]interface{})
	ChunkAreasInfoEx = make(map[uuid.UUID]AreaInfo)
	for _, m := range Map {
		var ai AreaInfo
		ChunkAreasInfo[m.ID] = make(map[string]interface{})
		for _, t := range m.Terrains {
			switch t.Type {
			case "forest":
				id := areas.CreateForest(m.ID, t.Size)
				ai.Forest = areas.GetForestById(id)
				ChunkAreasInfo[m.ID]["forest"] = ai.Forest
			case "hill":
				id := areas.CreateHill(m.ID, t.Size)
				ai.Hill = areas.GetHillsById(id)
				ChunkAreasInfo[m.ID]["hill"] = ai.Hill
			case "swamp":
				id := areas.CreateSwamp(m.ID, t.Size)
				ai.Swamp = areas.GetSwampById(id)
				ChunkAreasInfo[m.ID]["swamp"] = ai.Swamp
			case "meadow":
				id := areas.CreateMeadow(m.ID, t.Size)
				ai.Meadow = areas.GetMeadowById(id)
				ChunkAreasInfo[m.ID]["meadow"] = ai.Meadow
			case "lake":
				id := areas.CreateLake(m.ID, t.Size)
				ai.Lake = areas.GetLakesById(id)
				ChunkAreasInfo[m.ID]["lakes"] = ai.Lake
			}
		}
		var rs []areas.River
		for _, r := range m.Rivers {
			id := areas.CreateRiver(m.ID, r.Size, r.Bridge)
			rs = append(rs, areas.GetRiversById(id)[0])
		}
		if len(rs) > 0 {
			ai.Rivers = rs
			ChunkAreasInfo[m.ID]["rivers"] = rs
		}

		/*
			-- Озера пока не в отдельной структуре идут
			var ls []areas.Lake
			for _, l := range m.Lakes {
				id := areas.CreateLake(m.ID, l.Size)
				ls = append(ls, areas.GetLakesById(id)[0])
			}
			if len(ls) > 0 {
				ChunkAreasInfo[m.ID]["lakes"] = ls
			}
		*/

		ChunkAreasInfoEx[m.ID] = ai
		fmt.Println(ai)

		/*
			ChunkMasteryInfo[m.ID] = make(map[string]interface{})
			for k, v := range ChunkAreasInfo[m.ID] {
				fmt.Println(k)
				fmt.Println(v)
			}
		*/
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

func GetChunckAreasMastery(chunkId uuid.UUID, mastery string) {
	//fmt.Println(ChunkMasteryInfo[chunkId])
}
