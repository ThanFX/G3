package libs

type Mastery struct {
	Name     string `json:"name"`
	NameID   string `json:"id"`
	MinValue int    `json:"-"`
	MaxValue int    `json:"-"`
}

var masterships map[string]Mastery = map[string]Mastery{
	"fishing": Mastery{
		Name:     "Рыбная ловля",
		NameID:   "fishing",
		MinValue: 1,
		MaxValue: 100},
	"hunting": Mastery{
		Name:     "Охота",
		NameID:   "hunting",
		MinValue: 1,
		MaxValue: 100},
	"food_gathering": Mastery{
		Name:     "Собирательство грибов и ягод",
		NameID:   "food_gathering",
		MinValue: 1,
		MaxValue: 100}}

func GetMasterships() map[string]Mastery {
	return masterships
}

func GetMasteryByName(name string) Mastery {
	return masterships[name]
}
