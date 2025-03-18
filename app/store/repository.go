package store

import (
	"context"
	"warhoop/app/model"
)

type AuthRepo interface {
	// Account
	GetByID(context.Context, int) (*model.DBAccount, error)
	GetByUsername(context.Context, string) (*model.DBAccount, error)
	ExistsEmail(context.Context, string) (bool, error)
	ExistsUsername(context.Context, string) (bool, error)
	Create(context.Context, *model.DBAccount) (*model.DBAccount, error)
	// Access
	ExistsAccess(context.Context, int) (bool, error)
	GetAccessByID(context.Context, int) (*model.DBAccess, error)
	// Uptime
	GetUptimeByID(context.Context, int) (*model.DBUptime, error)
}

type CharRepo interface {
	// Characters
	GetByID(context.Context, int) (*model.DBCharacters, error)
	GetByName(context.Context, string) (*model.DBCharacters, error)
	GetTop10Kill(context.Context) (*model.DBCharactersSlice, error)
	GetOnlineCount(context.Context) (int, error)
	GetOnlineSlice(context.Context) (*model.DBCharactersSlice, error)
	// Armory

	GetArmoryCharactersSlice(context.Context, int, int) (*model.DBCharactersSlice, int, error)
}

type SaitRepo interface {
	// Profile
	CreateProfile(context.Context, *model.DBProfile) (*model.DBProfile, error)
	GetProfile(context.Context, int) (*model.DBProfile, error)
	// Avatar
	UpdateAvatar(context.Context, *model.DBProfile) error
	ExistAvatar(context.Context, int) (bool, error)
	// Session
	GetSession(context.Context, *model.DBSession) (*model.DBSession, error)
	CreateSession(context.Context, *model.DBSession) error
	UpdateSession(context.Context, *model.DBSession) error
	DeleteSession(context.Context, string) error
	ListSession(context.Context, uint) (*model.DBSessionSlice, error)
	UpdateOrCreateSession(context.Context, *model.DBSession, string) (*model.DBSession, error)
	// News
	CreateNews(context.Context, *model.DBNews) (*model.DBNews, error)
	GetNewsByID(context.Context, int) (*model.DBNews, error)
	GetNewsSlice(context.Context, int, int) (*model.DBNewsSlice, int, error)
	UpdateNews(context.Context, *model.DBNews) (*model.DBNews, error)
	DeleteNews(context.Context, int) error
	// Comments
	CreateComment(context.Context, *model.DBComment) (*model.DBComment, error)
	GetCommentByID(context.Context, int) (*model.DBComment, error)
	DeleteComment(context.Context, int) error
	UpdateComment(context.Context, *model.DBComment) (*model.DBComment, error)
	// Like
	ToggleReaction(context.Context, *model.DBReaction) (*model.DBReaction, error)
	DeleteReaction(context.Context, *model.DBReaction) error
	ExistReaction(context.Context, *model.DBReaction) (bool, error)
	// Report Methods
	CreateReport(context.Context, *model.DBReport) (*model.DBReport, error)
	GetReportByID(context.Context, int) (*model.DBReport, error)
	UpdateReport(context.Context, *model.DBReport) (*model.DBReport, error)
	GetReports(context.Context, int, int) (*model.DBReportSlice, error)
	DeleteReportByID(context.Context, int) error
	//Enchant DBC
	GetEnchantDBCByIDs(context.Context, []int32) (map[int32]int32, error)
}
