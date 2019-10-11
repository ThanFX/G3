package models

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
type Chunks struct {
	ID            uuid.UUID   `json:"id"`
	TerrainString string      `json:"-"`
	Chunk         interface{} `json:"chunk"`
}*/

var (
	Map []Chunk
)

func MapInitialize() {
	//Map = make([]Chunk, 25)
	rows, err := DB.Query("select chunk from map")
	if err != nil {
		log.Fatalf("Ошибка получения карты из БД: %s", err)
	}
	defer rows.Close()

	var ch string
	var chunk Chunk
	for rows.Next() {
		err = rows.Scan(&ch)
		if err != nil {
			log.Fatal("ошибка получения записи чанка: ", err)
		}
		err = json.Unmarshal([]byte(ch), &chunk)
		if err != nil {
			log.Fatal("ошибка парсинга данных чанка: ", err)
		}
		Map = append(Map, chunk)

		chunk.Terrains = nil
		chunk.Rivers = nil
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Карта успешно загружена")
}

func GetMap() []Chunk {
	return Map
}
