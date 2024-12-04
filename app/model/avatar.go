package model

import "github.com/uptrace/bun"

type Profile struct {
	AccountID int    `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar,omitempty" validate:"required,avatar"`
	Rank      string `json:"rank"`
}

type ProfileSlice []Profile

type DBProfile struct {
	bun.BaseModel `bun:"table:profile,alias:site"`
	AccountID     int    `bun:"account_id,pk"`
	Name          string `bun:"name"`
	Avatar        string `bun:"avatar"`
	Rank          string `bun:"rank"`
}

type DBProfileSlice []DBProfile

// ToDB converts Profile to DBProfile
func (entry *Profile) ToDB() *DBProfile {
	if entry == nil {
		return nil
	}
	return &DBProfile{
		AccountID: entry.AccountID,
		Name:      entry.Name,
		Avatar:    entry.Avatar,
		Rank:      entry.Rank,
	}
}

// ToWeb converts DBProfile to Profile
func (entry *DBProfile) ToWeb() *Profile {
	if entry == nil {
		return nil
	}
	return &Profile{
		AccountID: entry.AccountID,
		Name:      entry.Name,
		Avatar:    entry.Avatar,
		Rank:      entry.Rank,
	}
}

// ToDB converts ProfileSlice to DBProfileSlice
func (data ProfileSlice) ToDB() DBProfileSlice {
	result := make(DBProfileSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

// ToWeb converts DBProfileSlice to ProfileSlice
func (data DBProfileSlice) ToWeb() ProfileSlice {
	result := make(ProfileSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
