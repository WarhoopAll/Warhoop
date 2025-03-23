package char

import (
	"github.com/uptrace/bun"
	"warhoop/app/model/nexus"
)

type ItemInstance struct {
	Guid         int            `json:"-"`
	ItemEntry    int32          `json:"-"`
	Enchantments string         `json:"enchantments,omitempty"`
	Socket       string         `json:"socket,omitempty"`
	Gems         string         `json:"gems,omitempty"`
	ItemDBC      *nexus.ItemDBC `json:"item_display,omitempty"`
}

type ItemInstanceSlice []ItemInstance

type DBItemInstance struct {
	bun.BaseModel `bun:"table:item_instance,alias:item_instance"`
	Guid          int              `bun:"guid,pk"`
	ItemEntry     int32            `bun:"itemEntry"`
	Enchantments  string           `bun:"enchantments"`
	Socket        string           `bun:"-"`
	Gems          string           `bun:"-"`
	ItemDBC       *nexus.DBItemDBC `bun:"rel:belongs-to,join:itemEntry=ItemID"`
}

type DBItemInstanceSlice []DBItemInstance

func (entry *ItemInstance) ToDB() *DBItemInstance {
	if entry == nil {
		return nil
	}
	return &DBItemInstance{
		Guid:         entry.Guid,
		Enchantments: entry.Enchantments,
		Socket:       entry.Socket,
		Gems:         entry.Gems,
	}
}

func (entry *DBItemInstance) ToWeb() *ItemInstance {
	if entry == nil {
		return nil
	}
	return &ItemInstance{
		Guid:         entry.Guid,
		Enchantments: entry.Enchantments,
		Socket:       entry.Socket,
		Gems:         entry.Gems,
		ItemDBC:      entry.ItemDBC.ToWeb(),
	}
}

func (data ItemInstanceSlice) ToDB() DBItemInstanceSlice {
	result := make(DBItemInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBItemInstanceSlice) ToWeb() ItemInstanceSlice {
	result := make(ItemInstanceSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
