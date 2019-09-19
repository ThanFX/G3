package models

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

type Object struct {
	ID          int
	Level       int
	ParentID    int
	Name        string
	IsCountable bool
	IsLiquid    bool
	Limit       int
}

type Item struct {
	UUID         uuid.UUID
	Name         string
	Weight       int
	Limit        int
	Quality      int
	CreationDate int
	ExpDate      int
	Object
}

var (
	Items    []Item
	itemPool = sync.Pool{
		New: func() interface{} {
			return new(Item)
		},
	}
)

// Берём память из пула
func getItemPool() (item *Item) {
	mem := itemPool.Get()
	if mem != nil {
		item = mem.(*Item)
	}
	return
}

// Чистим экземпляр предмета перед возвратом в пул.
//Если хотим события, как в "Мире дикого Запада", то не вызываем эту функцию ))
func (item *Item) Reset() {
	item.UUID = uuid.Nil
	item.Name = ""
	item.Weight = 0
	item.Limit = 0
	item.Quality = 0
	item.CreationDate = 0
	item.ExpDate = 0
	item.Object.ID = 0
	item.Object.Level = 0
	item.Object.ParentID = 0
	item.Object.Name = ""
	item.Object.IsCountable = false
	item.Object.IsLiquid = false
	item.Object.Limit = 0
}

// Возвращаем ненужный экземпляр предмета в пул аллоцированной памяти
func putItemToPool(item Item) {
	item.Reset()
	itemPool.Put(item)
}
