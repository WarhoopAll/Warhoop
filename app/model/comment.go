package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	NewsID    int       `json:"news_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    int       `json:"author,omitempty"`
	Profile   *Profile  `json:"profile,omitempty"`
	LikeCount int       `json:"like_count"`
}

type CommentSlice []Comment

type DBComment struct {
	bun.BaseModel `bun:"table:news_comments,alias:news"`
	ID            int        `bun:"id,pk,autoincrement" `
	NewsID        int        `bun:"news_id"`
	Text          string     `bun:"text"`
	CreatedAt     time.Time  `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:"updated_at,nullzero,default:current_timestamp on update current_timestamp"`
	Author        int        `bun:"author,notnull"`
	Profile       *DBProfile `bun:"rel:belongs-to,join:author=account_id"`
	LikeCount     int        `bun:"like_count"`
}

type DBCommentSlice []DBComment

// ToDB converts Comment to DBComment
func (entry *Comment) ToDB() *DBComment {
	if entry == nil {
		return nil
	}
	return &DBComment{
		ID:        entry.ID,
		NewsID:    entry.NewsID,
		Text:      entry.Text,
		CreatedAt: entry.CreatedAt,
		Author:    entry.Author,
		UpdatedAt: entry.UpdatedAt,
		LikeCount: entry.LikeCount,
	}
}

// ToWeb converts DBComment to Comment
func (entry *DBComment) ToWeb() *Comment {
	if entry == nil {
		return nil
	}
	return &Comment{
		ID:        entry.ID,
		NewsID:    entry.NewsID,
		Text:      entry.Text,
		CreatedAt: entry.CreatedAt,
		UpdatedAt: entry.UpdatedAt,
		Profile:   entry.Profile.ToWeb(),
		LikeCount: entry.LikeCount,
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
