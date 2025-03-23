package nexus

import "github.com/uptrace/bun"

type ItemDBC struct {
	ItemID            int32  `json:"item_id,omitempty"`
	ItemDisplayInfoID int32  `json:"item_display_info_id,omitempty"`
	InventoryType     int32  `json:"inventory_type,omitempty"`
	InventoryIcon0    string `json:"inventory_icon_0,omitempty"`
	Quality           int    `json:"quality,omitempty"`
	ItemLevel         int    `json:"item_level,omitempty"`
	En                string `json:"en"`
	Ru                string `json:"ru"`
}

type ItemDBCSlice []ItemDBC

type DBItemDBC struct {
	bun.BaseModel     `bun:"table:sait.armory_item_dbc,alias:armory"`
	ItemID            int32  `bun:"ItemID,pk"`
	ItemDisplayInfoID int32  `bun:"ItemDisplayInfoID"`
	InventoryType     int32  `bun:"InventoryType" `
	InventoryIcon0    string `bun:"InventoryIcon_1"`
	Quality           int    `bun:"Quality" `
	ItemLevel         int    `bun:"ItemLevel"`
	En                string `bun:"en"`
	Ru                string `bun:"ru"`
}

type DBItemDBCSlice []DBItemDBC

func (entry *ItemDBC) ToDB() *DBItemDBC {
	if entry == nil {
		return nil
	}
	return &DBItemDBC{
		ItemID:            entry.ItemID,
		ItemDisplayInfoID: entry.ItemDisplayInfoID,
		InventoryType:     entry.InventoryType,
		InventoryIcon0:    entry.InventoryIcon0,
		Quality:           entry.Quality,
		ItemLevel:         entry.ItemLevel,
		En:                entry.En,
		Ru:                entry.Ru,
	}
}

func (entry *DBItemDBC) ToWeb() *ItemDBC {
	if entry == nil {
		return nil
	}
	return &ItemDBC{
		ItemID:            entry.ItemID,
		ItemDisplayInfoID: entry.ItemDisplayInfoID,
		InventoryType:     entry.InventoryType,
		InventoryIcon0:    entry.InventoryIcon0,
		Quality:           entry.Quality,
		ItemLevel:         entry.ItemLevel,
		En:                entry.En,
		Ru:                entry.Ru,
	}
}

func (data ItemDBCSlice) ToDB() DBItemDBCSlice {
	result := make(DBItemDBCSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemDBCSlice) ToWeb() ItemDBCSlice {
	result := make(ItemDBCSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
