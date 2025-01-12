package model

import (
	"github.com/uptrace/bun"
	"time"
)

type News struct {
	ID            int           `json:"id,omitempty"`
	Title         string        `json:"title,omitempty"`
	Text          string        `json:"text,omitempty"`
	CreatedAt     time.Time     `json:"created_at,omitempty"`
	UpdatedAt     time.Time     `json:"updated_at,omitempty"`
	Comment       *CommentSlice `json:"comments,omitempty"`
	ProfileID     int           `json:"-"`
	Profile       *Profile      `json:"profile,omitempty"`
	ImageUrl      string        `json:"image_url"`
	LikeCount     int           `json:"like_count,omitempty"`
	DisLikeCount  int           `json:"dislike_count,omitempty"`
	CommentsCount int           `json:"comments_count"`
}

type NewsSlice []News

type DBNews struct {
	bun.BaseModel `bun:"table:news,alias:news"`
	ID            int             `bun:"id,pk,autoincrement"`
	Title         string          `bun:"title"`
	Text          string          `bun:"text" `
	CreatedAt     time.Time       `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt     time.Time       `bun:"updated_at,nullzero,default:current_timestamp on update current_timestamp"`
	ProfileID     int             `bun:"profile_id"`
	Profile       *DBProfile      `bun:"rel:belongs-to,join:profile_id=account_id"`
	Comments      *DBCommentSlice `bun:"-"`
	ImageUrl      string          `bun:"image_url"`
	LikeCount     int             `bun:"like_count"`
	DisLikeCount  int             `bun:"dislike_count"`
	CommentsCount int             `bun:"comments_count"`
}

type DBNewsSlice []DBNews

// ToDB converts News to DBNews
func (entry *News) ToDB() *DBNews {
	if entry == nil {
		return nil
	}

	return &DBNews{
		ID:            entry.ID,
		Title:         entry.Title,
		Text:          entry.Text,
		CreatedAt:     entry.CreatedAt,
		UpdatedAt:     entry.UpdatedAt,
		ProfileID:     entry.ProfileID,
		ImageUrl:      entry.ImageUrl,
		LikeCount:     entry.LikeCount,
		DisLikeCount:  entry.DisLikeCount,
		CommentsCount: entry.CommentsCount,
	}
}

// ToWeb converts DBNews to News
func (entry *DBNews) ToWeb() *News {
	if entry == nil {
		return nil
	}

	var comments *CommentSlice
	if entry.Comments != nil {
		temp := entry.Comments.ToWeb()
		comments = &temp
	}

	return &News{
		ID:            entry.ID,
		Title:         entry.Title,
		Text:          entry.Text,
		CreatedAt:     entry.CreatedAt,
		UpdatedAt:     entry.UpdatedAt,
		Profile:       entry.Profile.ToWeb(),
		Comment:       comments,
		ImageUrl:      entry.ImageUrl,
		LikeCount:     entry.LikeCount,
		DisLikeCount:  entry.DisLikeCount,
		CommentsCount: entry.CommentsCount,
	}
}

// ToDB converts NewsSlice to DBNewsSlice
func (data NewsSlice) ToDB() DBNewsSlice {
	result := make(DBNewsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

// ToWeb converts DBNewsSlice to NewsSlice
func (data DBNewsSlice) ToWeb() NewsSlice {
	result := make(NewsSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
