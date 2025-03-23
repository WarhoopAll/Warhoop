package nexus

import (
	"github.com/uptrace/bun"
	"time"
)

type Reaction struct {
	ID           int       `json:"id,omitempty"`
	UserID       int       `json:"user_id,omitempty"`
	ObjectType   int       `json:"object_type,omitempty"`
	ObjectID     int       `json:"object_id,omitempty"`
	ReactionType int       `json:"reaction_type,omitempty"`
	CreatedAt    time.Time `json:"-"`
}

type ReactionSlice []Reaction

type DBReaction struct {
	bun.BaseModel `bun:"table:Reactions,alias:Reactions"`
	ID            int       `bun:"id,pk,autoincrement"`
	UserID        int       `bun:"user_id"`
	ObjectType    int       `bun:"object_type"` // 1 = News, 2 = Comment, 3 = Profile
	ObjectID      int       `bun:"object_id"`
	ReactionType  int       `bun:"reaction_type"` // 1 = Like, 2 = Dislike
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
}

type DBReactionSlice []DBReaction

// ToDB converts Reaction to DBReaction
func (entry *Reaction) ToDB() *DBReaction {
	if entry == nil {
		return nil
	}
	return &DBReaction{
		ID:           entry.ID,
		UserID:       entry.UserID,
		ObjectType:   entry.ObjectType,
		ObjectID:     entry.ObjectID,
		ReactionType: entry.ReactionType,
		CreatedAt:    entry.CreatedAt,
	}
}

// ToWeb converts DBReaction to Reaction
func (entry *DBReaction) ToWeb() *Reaction {
	if entry == nil {
		return nil
	}
	return &Reaction{
		ID:           entry.ID,
		UserID:       entry.UserID,
		ObjectType:   entry.ObjectType,
		ObjectID:     entry.ObjectID,
		ReactionType: entry.ReactionType,
		CreatedAt:    entry.CreatedAt,
	}
}

// ToDB converts ReactionSlice to DBReactionSlice
func (data ReactionSlice) ToDB() DBReactionSlice {
	result := make(DBReactionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

// ToWeb converts DBReactionSlice to ReactionSlice
func (data DBReactionSlice) ToWeb() ReactionSlice {
	result := make(ReactionSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
