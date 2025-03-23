package nexus

import "github.com/uptrace/bun"

type EnchantDBC struct {
	ID        int32 `json:"id"`
	SrcItemID int32 `json:"src_item_id"`
}

type EnchantDBCSlice []EnchantDBC

type DBEnchantDBC struct {
	bun.BaseModel `bun:"table:sait.armory_enchant_dbc,alias:armory"`
	ID            int32 `bun:"ID,pk"`
	SrcItemID     int32 `bun:"src_item_id"`
}

type DBEnchantDBCSlice []DBEnchantDBC

func (entry *EnchantDBC) ToDB() *DBEnchantDBC {
	if entry == nil {
		return nil
	}
	return &DBEnchantDBC{
		ID:        entry.ID,
		SrcItemID: entry.SrcItemID,
	}
}

func (entry *DBEnchantDBC) ToWeb() *EnchantDBC {
	if entry == nil {
		return nil
	}
	return &EnchantDBC{
		ID:        entry.ID,
		SrcItemID: entry.SrcItemID,
	}
}

func (data EnchantDBCSlice) ToDB() DBEnchantDBCSlice {
	result := make(DBEnchantDBCSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBEnchantDBCSlice) ToWeb() EnchantDBCSlice {
	result := make(EnchantDBCSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
