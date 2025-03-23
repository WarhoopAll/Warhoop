package nexus

import "github.com/uptrace/bun"

type Zone struct {
	ID     int32  `json:"-"`
	LangEn string `json:"zone_lang_en,omitempty"`
	LangRu string `json:"zone_lang_ru,omitempty"`
}

type DBZone struct {
	bun.BaseModel `bun:"table:sait.zone"`
	ID            int32  `bun:"ID,pk"`
	LangEn        string `bun:"AreaName_Lang_enUS"`
	LangRu        string `bun:"AreaName_Lang_ruRU"`
}

func (entry *DBZone) ToWeb() *Zone {
	if entry == nil {
		return nil
	}
	return &Zone{
		ID:     entry.ID,
		LangEn: entry.LangEn,
		LangRu: entry.LangRu,
	}
}
