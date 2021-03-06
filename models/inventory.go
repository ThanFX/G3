package models

import uuid "github.com/satori/go.uuid"

type InventoryItem struct {
	ItemID       uuid.UUID
	creationDate int
}

type PersonInventory struct {
	ID           uuid.UUID `json:"id"`
	ItemID       int       `json:"-"`
	Name         string    `json:"name"`
	Weight       int       `json:"weight"`
	Limit        int       `json:"limit"`
	Quality      int       `json:"quality"`
	CreationDate int       `json:"creation_date"`
	ExpDate      int       `json:"expiration_date"`
	Category     string    `json:"category"`
	IsCountable  bool      `json:"is_countable"`
}
