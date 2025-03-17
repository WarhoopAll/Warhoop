package model

import "github.com/uptrace/bun"

type CharacterInventory struct {
	Guid         int           `json:"-"`
	Bag          int           `json:"bag,omitempty"`
	Slot         int8          `json:"slot"`
	Item         int           `json:"-"`
	ItemInstance *ItemInstance `json:"item_instance,omitempty"`
}

type CharacterInventorySlice []CharacterInventory

type DBCharacterInventory struct {
	bun.BaseModel `bun:"table:character_inventory,alias:character_inventory"`
	Guid          int             `bun:"guid,pk"`
	Bag           int             `bun:"bag"`
	Slot          int8            `bun:"slot"`
	Item          int             `bun:"item"`
	ItemInstance  *DBItemInstance `bun:"rel:belongs-to,join:item=guid"`
}

type DBCharacterInventorySlice []DBCharacterInventory

func (entry *CharacterInventory) ToDB() *DBCharacterInventory {
	if entry == nil {
		return nil
	}
	return &DBCharacterInventory{
		Guid: entry.Guid,
		Bag:  entry.Bag,
		Slot: entry.Slot,
		Item: entry.Item,
	}
}

func (entry *DBCharacterInventory) ToWeb() *CharacterInventory {
	if entry == nil {
		return nil
	}
	return &CharacterInventory{
		Guid:         entry.Guid,
		Bag:          entry.Bag,
		Slot:         entry.Slot,
		Item:         entry.Item,
		ItemInstance: entry.ItemInstance.ToWeb(),
	}
}

func (data CharacterInventorySlice) ToDB() DBCharacterInventorySlice {
	result := make(DBCharacterInventorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCharacterInventorySlice) ToWeb() CharacterInventorySlice {
	result := make(CharacterInventorySlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
