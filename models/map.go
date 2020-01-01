package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ThanFX/G3/libs"
)

/*
type Chunks struct {
	ID            uuid.UUID   `json:"id"`
	TerrainString string      `json:"-"`
	Chunk         interface{} `json:"chunk"`
}*/

var (
	Map []libs.Chunk
	//DB  *sql.DB
)

func MapInitialize() {
	//Map = make([]Chunk, 25)
	rows, err := DB.Query("select chunk from map")
	if err != nil {
		log.Fatalf("Ошибка получения карты из БД: %s", err)
	}
	defer rows.Close()

	var ch string
	var chunk libs.Chunk
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

func GetMap() []libs.Chunk {
	return Map
}
