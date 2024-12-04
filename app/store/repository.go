package store

import (
	"context"
	"grimoire/app/model"
)

type AuthRepo interface {
	// Account
	GetByID(context.Context, int) (*model.DBAccount, error)
	GetByUsername(context.Context, string) (*model.DBAccount, error)
	ExistsEmail(context.Context, string) (bool, error)
	ExistsUsername(context.Context, string) (bool, error)
	Create(context.Context, *model.DBAccount) (*model.DBAccount, error)
}

type CharRepo interface {
	// Characters
	GetByID(context.Context, int) (*model.DBCharacters, error)
	GetByName(context.Context, string) (*model.DBCharacters, error)
	GetCharTop10Kill(context.Context) (*model.DBCharactersSlice, error)
}

type SaitRepo interface {
	// Avatar
	GetAvatar(ctx context.Context, id int) (*model.DBProfile, error)
	UpdateAvatar(context.Context, *model.DBProfile) error
	ExistAvatar(context.Context, int) (bool, error)
	// Session
	GetSession(context.Context, *model.DBSession) (*model.DBSession, error)
	CreateSession(context.Context, *model.DBSession) error
	UpdateSession(context.Context, *model.DBSession) error
	DeleteSession(context.Context, string) error
	ListSession(context.Context, uint) (model.DBSessionSlice, error)
	UpdateOrCreateSession(context.Context, *model.DBSession, string) (*model.DBSession, error)
	// News
	GetNewsByID(context.Context, *model.DBNews) (*model.DBNews, error)
	GetNewsSlice(context.Context) (*model.DBNewsSlice, error)
	// Comments
	AddComment(context.Context, *model.DBComment) (*model.DBComment, error)
	GetCommentByID(context.Context, int) (*model.DBComment, error)
	DeleteComment(context.Context, int) error
	UpdateComment(context.Context, *model.DBComment) error
}
