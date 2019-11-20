package models

import (
	"fmt"

	"github.com/ThanFX/G3/areas"

	uuid "github.com/satori/go.uuid"
)

type AreaMastery struct {
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
	ChunkMasteryInfo map[uuid.UUID]map[string][]AreaMastery
	ChunkAreasInfo   map[uuid.UUID]map[string]interface{}
	ChunkAreasInfoEx map[uuid.UUID]AreaInfo
)

func CreateTerrains() {
	ChunkMasteryInfo = make(map[uuid.UUID]map[string][]AreaMastery)
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
		//fmt.Println(ai)

		ChunkMasteryInfo[m.ID] = make(map[string][]AreaMastery)
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Forest.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Forest.Masterships {
				var am AreaMastery
				am.Name = "forest"
				am.Size = ChunkAreasInfoEx[m.ID].Forest.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Forest.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Hill.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Hill.Masterships {
				var am AreaMastery
				am.Name = "hill"
				am.Size = ChunkAreasInfoEx[m.ID].Hill.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Hill.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Swamp.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Swamp.Masterships {
				var am AreaMastery
				am.Name = "swamp"
				am.Size = ChunkAreasInfoEx[m.ID].Swamp.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Swamp.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Meadow.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Meadow.Masterships {
				var am AreaMastery
				am.Name = "meadow"
				am.Size = ChunkAreasInfoEx[m.ID].Meadow.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Meadow.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Lake.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Lake.Masterships {
				var am AreaMastery
				am.Name = "lake"
				am.Size = ChunkAreasInfoEx[m.ID].Lake.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Lake.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if len(ChunkAreasInfoEx[m.ID].Rivers) > 0 && !uuid.Equal(ChunkAreasInfoEx[m.ID].Rivers[0].ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Rivers {
				for _, vr := range v.Masterships {
					var am AreaMastery
					am.Name = "river"
					am.Size = v.Size
					am.AreaID = v.ID
					ChunkMasteryInfo[m.ID][vr.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][vr.Mastership.NameID], am)

				}

			}

		}
	}
	//fmt.Println(ChunkMasteryInfo)
}

func GetChunkTerrainsInfo(param string) map[string]interface{} {
	chunkId, err := uuid.FromString(param)
	if err != nil {
		fmt.Printf("При получении ID чанка %s произошла ошибка %s", param, err)
		return nil
	}
	return ChunkAreasInfo[chunkId]
}

func GetChunckAreasMastery(chunkId uuid.UUID) map[string][]AreaMastery {
	return ChunkMasteryInfo[chunkId]
}
