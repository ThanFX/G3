package models

import (
	uuid "github.com/satori/go.uuid"
)

type Chunk struct {
	ID       uuid.UUID `json:"id"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
	Terrains []Terrain `json:"terrains"`
	Rivers   []River   `json:"rivers"`
}

type Terrain struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

type River struct {
	Size   int    `json:"size"`
	From   string `json:"from"`
	To     string `json:"to"`
	Bridge bool   `json:"bridge"`
}

//var testStr string = `{"ID":"5260441b-22fd-4564-8b64-60973dfdce7b","X":5,"Y":-7,"Terrains":[{"Type":"forest","Size":1},{"Type":"meadow","Size":4}],"Rivers":[{"Size":1,"From":"W","To":"E","Bridge":true}]}`
//var ch Chunk

/*var tch = Chunk{
ID: uuid.Must(uuid.FromString("5260441b-22fd-4564-8b64-60973dfdce7b")),
X:  5,
Y:  -7,
Terrains: []Terrain{
	{
		Type: "forest",
		Size: 1},
	{
		Type: "meadow",
		Size: 4}},
Rivers: []River{
	{
		Size:   1,
		From:   "W",
		To:     "E",
		Bridge: true}}}


func CreateChunk() {
	//s, err := json.Marshal(tch)
	err := json.Unmarshal([]byte(testStr), &ch)
	if err != nil {
		fmt.Printf("Ошибка парсинга чанка %s", err)
	} else {
		fmt.Println(ch)
	}

}
*/
