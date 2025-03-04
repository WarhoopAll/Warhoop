package model

import "github.com/uptrace/bun"

type Map struct {
	ID     int32  `json:"-"`
	LangEn string `json:"map_lang_en,omitempty"`
	LangRU string `json:"map_lang_ru,omitempty"`
}

type DBMap struct {
	bun.BaseModel `bun:"table:sait.map"`
	ID            int32  `bun:"ID,pk"`
	LangEn        string `bun:"MapName_Lang_enUS"`
	LangRU        string `bun:"MapName_Lang_ruRU"`
}

func (entry *DBMap) ToWeb() *Map {
	if entry == nil {
		return nil
	}
	return &Map{
		ID:     entry.ID,
		LangEn: entry.LangEn,
		LangRU: entry.LangRU,
	}
}
