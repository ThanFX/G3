package models

import uuid "github.com/satori/go.uuid"

type InventoryItem struct {
	ItemID       uuid.UUID
	creationDate int
}

type PersonInventory struct {
	ID           uuid.UUID
	Name         string
	Weight       int
	Limit        int
	Quality      int
	CreationDate int
	ExpDate      int
}
