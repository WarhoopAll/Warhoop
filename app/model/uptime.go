package model

import "github.com/uptrace/bun"

type Uptime struct {
	Realmid    int        `json:"-"`
	Starttime  int        `json:"starttime,omitempty"`
	Uptime     int        `json:"uptime,omitempty"`
	Maxplayers int16      `json:"maxplayers,omitempty"`
	Realm      *Realmlist `json:"realm,omitempty"`
	Revision   string     `json:"revision,omitempty"`
	CharOnline int        `json:"charOnline"`
}

type UptimeSlice []Uptime

type DBUptime struct {
	bun.BaseModel `bun:"table:uptime,alias:uptime"`
	Realmid       int          `bun:"realmid,pk"`
	Realm         *DBRealmlist `bun:"rel:belongs-to,join:realmid=id"`
	Starttime     int          `bun:"starttime"`
	Uptime        int          `bun:"uptime"`
	Maxplayers    int16        `bun:"maxplayers"`
	Revision      string       `bun:"revision"`
}

type DBUptimeSlice []DBUptime

func (entry *Uptime) ToDB() *DBUptime {
	if entry == nil {
		return nil
	}
	return &DBUptime{
		Realmid:    entry.Realmid,
		Starttime:  entry.Starttime,
		Uptime:     entry.Uptime,
		Maxplayers: entry.Maxplayers,
		Revision:   entry.Revision,
	}
}

func (entry *DBUptime) ToWeb() *Uptime {
	if entry == nil {
		return nil
	}
	return &Uptime{
		Realmid:    entry.Realmid,
		Starttime:  entry.Starttime,
		Uptime:     entry.Uptime,
		Maxplayers: entry.Maxplayers,
		Revision:   entry.Revision,
		Realm:      entry.Realm.ToWeb(),
	}
}

func (data UptimeSlice) ToDB() DBUptimeSlice {
	result := make(DBUptimeSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBUptimeSlice) ToWeb() UptimeSlice {
	result := make(UptimeSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
