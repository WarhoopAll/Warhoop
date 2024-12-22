package model

import "github.com/uptrace/bun"

type Access struct {
	AccountID     int    `json:"-"`
	SecurityLevel int8   `json:"security_level,omitempty"`
	RealmID       int    `json:"realm_id,omitempty"`
	Comment       string `json:"comment,omitempty"`
}
type AccessSlice []Access

type DBAccess struct {
	bun.BaseModel `bun:"table:account_access,alias:account"`
	AccountID     int    `bun:"AccountID,pk" `
	SecurityLevel int8   `bun:"SecurityLevel" `
	RealmID       int    `bun:"RealmID,pk" `
	Comment       string `bun:"Comment,nullzero" `
}
type DBAccessSlice []DBAccess

// ToDB converts Access to DBAccess
func (entry *Access) ToDB() *DBAccess {
	if entry == nil {
		return nil
	}
	return &DBAccess{
		AccountID:     entry.AccountID,
		SecurityLevel: entry.SecurityLevel,
		RealmID:       entry.RealmID,
		Comment:       entry.Comment,
	}
}

// ToWeb converts DBAccess to Access
func (entry *DBAccess) ToWeb() *Access {
	if entry == nil {
		return nil
	}
	return &Access{
		AccountID:     entry.AccountID,
		SecurityLevel: entry.SecurityLevel,
		RealmID:       entry.RealmID,
		Comment:       entry.Comment,
	}
}

// ToDB converts AccessSlice to DBAccessSlice
func (data AccessSlice) ToDB() DBAccessSlice {
	result := make(DBAccessSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

// ToWeb converts DBAccessSlice to AccessSlice
func (data DBAccessSlice) ToWeb() AccessSlice {
	result := make(AccessSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
