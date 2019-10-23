package libs

type Mastery struct {
	Name     string `json:"name"`
	MinValue int    `json:"-"`
	MaxValue int    `json:"-"`
}

var masterships map[string]Mastery = map[string]Mastery{
	"fishing": Mastery{
		Name:     "Рыбная ловля",
		MinValue: 1,
		MaxValue: 100},
	"hunting": Mastery{
		Name:     "Охота",
		MinValue: 1,
		MaxValue: 100},
	"food_gathering": Mastery{
		Name:     "Собирательство грибов и ягод",
		MinValue: 1,
		MaxValue: 100}}

func GetMasterships() map[string]Mastery {
	return masterships
}

func GetMasteryByName(name string) Mastery {
	return masterships[name]
}
