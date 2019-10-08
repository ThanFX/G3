package models

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Chunk struct {
	ID       uuid.UUID
	X        int
	Y        int
	Terrains []Terrain
	Rivers   []River
}

type Terrain struct {
	Type string
	Size int
}

type River struct {
	Size   int
	From   string
	To     string
	Bridge bool
}

var testStr string = `{"x":5,"y":-7,"id":"5260441b-22fd-4564-8b64-60973dfdce7b","terrains":[{"terrain":{"type":"forest","size":1}},{"terrain":{"type":"meadow","size":4}}],"rivers":[{"river":{"size":1,"bridge":true,"frow":"W","to":"E"}}]}`
var ch Chunk

func CreateChunk() {
	err := json.Unmarshal([]byte(testStr), &ch)
	if err != nil {
		fmt.Printf("Ошибка парсинга чанка %s", err)
	} else {
		fmt.Println(ch)
	}

}
