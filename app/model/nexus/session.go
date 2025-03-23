package nexus

import (
	"github.com/uptrace/bun"
	"time"
)

type Session struct {
	ID        uint      `json:"-"`
	Token     string    `json:"-"`
	AccountID int       `json:"-"`
	IPs       []string  `json:"ips,omitempty"`
	Finger    string    `json:"finger"`
	ExpiredAt time.Time `json:"-"`
	LoginedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// SessionSlice
type SessionSlice []Session

type DBSession struct {
	bun.BaseModel `bun:"table:sessions,alias:sessions"`
	ID            uint      `bun:"id,pk,notnull,autoincrement"`
	Token         string    `bun:"token,notnull"`
	AccountID     int       `bun:"account_id,notnull"`
	IPs           []string  `bun:"ips,array"`
	Finger        string    `bun:"fingerprint"`
	ExpiredAt     time.Time `bun:"expired_at"`
	LoginedAt     time.Time `bun:"logined_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at"`
}

// DBSessionSlice
type DBSessionSlice []DBSession

// ToDB converts Session to DBSession
func (entry *Session) ToDB() *DBSession {
	if entry == nil {
		return nil
	}
	return &DBSession{
		ID:        entry.ID,
		Token:     entry.Token,
		AccountID: entry.AccountID,
		IPs:       entry.IPs,
		Finger:    entry.Finger,
		ExpiredAt: entry.ExpiredAt,
		LoginedAt: entry.LoginedAt,
		UpdatedAt: entry.UpdatedAt,
	}
}

// ToWeb converts DBSession to Session
func (entry *DBSession) ToWeb() *Session {
	if entry == nil {
		return nil
	}
	return &Session{
		ID:        entry.ID,
		Token:     entry.Token,
		AccountID: entry.AccountID,
		IPs:       entry.IPs,
		Finger:    entry.Finger,
		ExpiredAt: entry.ExpiredAt,
		LoginedAt: entry.LoginedAt,
		UpdatedAt: entry.UpdatedAt,
	}
}

// ToDB converts SessionSlice to DBSessionSlice
func (data SessionSlice) ToDB() DBSessionSlice {
	result := make(DBSessionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

// ToWeb converts DBSessionSlice to SessionSlice
func (data DBSessionSlice) ToWeb() SessionSlice {
	result := make(SessionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
