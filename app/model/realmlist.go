package model

import "github.com/uptrace/bun"

type Realmlist struct {
	Id                   int     `json:"-"`
	Name                 string  `json:"name,omitempty"`
	Address              string  `json:"address,omitempty"`
	LocalAddress         string  `json:"-"`
	LocalSubnetMask      string  `json:"-"`
	Port                 int16   `json:"-"`
	Icon                 int8    `json:"-"`
	Flag                 int8    `json:"-"`
	Timezone             int8    `json:"-"`
	AllowedSecurityLevel int8    `json:"-"`
	Population           float64 `json:"-"`
	Gamebuild            int     `json:"-"`
}

type RealmlistSlice []Realmlist

type DBRealmlist struct {
	bun.BaseModel        `bun:"table:realmlist,alias:realmlist"`
	Id                   int     `bun:"id,pk"`
	Name                 string  `bun:"name"`
	Address              string  `bun:"address"`
	LocalAddress         string  `bun:"localAddress"`
	LocalSubnetMask      string  `bun:"localSubnetMask"`
	Port                 int16   `bun:"port"`
	Icon                 int8    `bun:"icon"`
	Flag                 int8    `bun:"flag"`
	Timezone             int8    `bun:"timezone"`
	AllowedSecurityLevel int8    `bun:"allowedSecurityLevel"`
	Population           float64 `bun:"population"`

	Gamebuild int `bun:"gamebuild"`
}

type DBRealmlistSlice []DBRealmlist

func (entry *Realmlist) ToDB() *DBRealmlist {
	if entry == nil {
		return nil
	}
	return &DBRealmlist{
		Id:                   entry.Id,
		Name:                 entry.Name,
		Address:              entry.Address,
		LocalAddress:         entry.LocalAddress,
		LocalSubnetMask:      entry.LocalSubnetMask,
		Port:                 entry.Port,
		Icon:                 entry.Icon,
		Flag:                 entry.Flag,
		Timezone:             entry.Timezone,
		AllowedSecurityLevel: entry.AllowedSecurityLevel,
		Population:           entry.Population,
		Gamebuild:            entry.Gamebuild,
	}
}

func (entry *DBRealmlist) ToWeb() *Realmlist {
	if entry == nil {
		return nil
	}

	return &Realmlist{
		Id:                   entry.Id,
		Name:                 entry.Name,
		Address:              entry.Address,
		LocalAddress:         entry.LocalAddress,
		LocalSubnetMask:      entry.LocalSubnetMask,
		Port:                 entry.Port,
		Icon:                 entry.Icon,
		Flag:                 entry.Flag,
		Timezone:             entry.Timezone,
		AllowedSecurityLevel: entry.AllowedSecurityLevel,
		Population:           entry.Population,
		Gamebuild:            entry.Gamebuild,
	}
}

func (data RealmlistSlice) ToDB() DBRealmlistSlice {
	result := make(DBRealmlistSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBRealmlistSlice) ToWeb() RealmlistSlice {
	result := make(RealmlistSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
