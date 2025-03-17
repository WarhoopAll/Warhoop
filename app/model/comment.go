package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Comment struct {
	ID           int       `json:"id,omitempty"`
	NewsID       int       `json:"news_id,omitempty"`
	Text         string    `json:"text,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	ProfileID    int       `json:"-"`
	Profile      *Profile  `json:"profile,omitempty"`
	LikeCount    int       `json:"like_count,omitempty"`
	DislikeCount int       `json:"dislike_count,omitempty"`
}

type CommentSlice []Comment

type DBComment struct {
	bun.BaseModel `bun:"table:news_comments,alias:news"`
	ID            int        `bun:"id,pk,autoincrement" `
	NewsID        int        `bun:"news_id"`
	Text          string     `bun:"text"`
	CreatedAt     time.Time  `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:"updated_at,nullzero,default:current_timestamp on update current_timestamp"`
	ProfileID     int        `bun:"profile_id,notnull"`
	Profile       *DBProfile `bun:"rel:belongs-to,join:profile_id=account_id"`
	LikeCount     int        `bun:"like_count"`
	DislikeCount  int        `bun:"dislike_count"`
}

type DBCommentSlice []DBComment

// ToDB converts Comment to DBComment
func (entry *Comment) ToDB() *DBComment {
	if entry == nil {
		return nil
	}

	return &DBComment{
		ID:           entry.ID,
		NewsID:       entry.NewsID,
		Text:         entry.Text,
		CreatedAt:    entry.CreatedAt,
		ProfileID:    entry.ProfileID,
		UpdatedAt:    entry.UpdatedAt,
		LikeCount:    entry.LikeCount,
		DislikeCount: entry.DislikeCount,
	}
}

// ToWeb converts DBComment to Comment
func (entry *DBComment) ToWeb() *Comment {
	if entry == nil {
		return nil
	}

	return &Comment{
		ID:           entry.ID,
		NewsID:       entry.NewsID,
		Text:         entry.Text,
		CreatedAt:    entry.CreatedAt,
		UpdatedAt:    entry.UpdatedAt,
		Profile:      entry.Profile.ToWeb(),
		LikeCount:    entry.LikeCount,
		DislikeCount: entry.DislikeCount,
	}
}

// ToDB converts CommentSlice to DBCommentSlice
func (data CommentSlice) ToDB() DBCommentSlice {
	result := make(DBCommentSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

// ToWeb converts DBCommentSlice to CommentSlice
func (data DBCommentSlice) ToWeb() CommentSlice {
	result := make(CommentSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
