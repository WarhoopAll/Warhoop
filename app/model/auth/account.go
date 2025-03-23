package auth

import (
	"github.com/uptrace/bun"
	"warhoop/app/model/nexus"
)

type Account struct {
	ID       int            `json:"-"`
	Username string         `json:"username,omitempty" validate:"required,min=4,max=20"`
	Password string         `json:"password,omitempty" validate:"required,min=6,max=30"`
	Salt     []byte         `json:"-"`
	Verifier []byte         `json:"-"`
	Access   *Access        `json:"access,omitempty"`
	Profile  *nexus.Profile `json:"profile,omitempty" validate:"omitempty,eq=0"`
	Email    string         `json:"email,omitempty" validate:"omitempty,email"`
	LastIP   string         `json:"last_ip,omitempty" validate:"omitempty,eq=0"`
}

type AccountSlice []Account

type DBAccount struct {
	bun.BaseModel `bun:"table:account,alias:account"`
	ID            int              `bun:"id,notnull,unique,pk,autoincrement"`
	Username      string           `bun:"username,notnull"`
	Salt          []byte           `bun:"salt,notnull"`
	Verifier      []byte           `bun:"verifier,notnull"`
	Access        *DBAccess        `bun:"rel:belongs-to,join:id=AccountID"`
	Profile       *nexus.DBProfile `bun:"-"`
	Email         string           `bun:"email,notnull"`
	LastIP        string           `bun:"last_ip"`
}

type DBAccountSlice []DBAccount

// ToDB converts Account to DBAccount
func (entry *Account) ToDB() *DBAccount {
	if entry == nil {
		return nil
	}
	return &DBAccount{
		ID:       entry.ID,
		Username: entry.Username,
		Salt:     entry.Salt,
		Verifier: entry.Verifier,
		Email:    entry.Email,
		LastIP:   entry.LastIP,
	}
}

// ToWeb converts DBAccount to Account
func (entry *DBAccount) ToWeb() *Account {
	if entry == nil {
		return nil
	}
	return &Account{
		ID:       entry.ID,
		Username: entry.Username,
		Salt:     entry.Salt,
		Verifier: entry.Verifier,
		Email:    entry.Email,
		LastIP:   entry.LastIP,
		Access:   entry.Access.ToWeb(),
		Profile:  entry.Profile.ToWeb(),
	}
}

// ToDB converts AccountSlice to DBAccountSlice
func (data AccountSlice) ToDB() DBAccountSlice {
	result := make(DBAccountSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

// ToWeb converts DBAccountSlice to AccountSlice
func (data DBAccountSlice) ToWeb() AccountSlice {
	result := make(AccountSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
