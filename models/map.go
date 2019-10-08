package models

import (
	"encoding/json"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
)

type Chunks struct {
	ID            uuid.UUID   `json:"id"`
	TerrainString string      `json:"-"`
	Chunk         interface{} `json:"chunk"`
}

var (
	Map []Chunks
)

func MapInitialize() {
	//Map = make([]Chunk, 25)
	rows, err := DB.Query("select * from map")
	if err != nil {
		log.Fatalf("Ошибка получения карты из БД: %s", err)
	}
	defer rows.Close()

	var ch Chunks
	for rows.Next() {
		err = rows.Scan(&ch.ID, &ch.TerrainString)
		if err != nil {
			log.Fatal("ошибка парсинга записи чанка: ", err)
		}
		err = json.Unmarshal([]byte(ch.TerrainString), &ch.Chunk)
		if err != nil {
			log.Fatal("ошибка парсинга данных чанка: ", err)
		}
		Map = append(Map, ch)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Карта успешно загружена")
}

func GetMap() []Chunks {
	return Map
}
