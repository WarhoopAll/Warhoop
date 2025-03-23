package store

import (
	"context"
	"warhoop/app/model/auth"
	"warhoop/app/model/char"
	"warhoop/app/model/nexus"
)

type AuthRepo interface {
	// Account
	GetByID(context.Context, int) (*auth.DBAccount, error)
	GetByUsername(context.Context, string) (*auth.DBAccount, error)
	ExistsEmail(context.Context, string) (bool, error)
	ExistsUsername(context.Context, string) (bool, error)
	Create(context.Context, *auth.DBAccount) (*auth.DBAccount, error)
	// Access
	ExistsAccess(context.Context, int) (bool, error)
	GetAccessByID(context.Context, int) (*auth.DBAccess, error)
	// Uptime
	GetUptimeByID(context.Context, int) (*auth.DBUptime, error)
}

type CharRepo interface {
	// Characters
	GetByID(context.Context, int) (*char.DBCharacters, error)
	GetByName(context.Context, string) (*char.DBCharacters, error)
	GetTop10Kill(context.Context) (*char.DBCharactersSlice, error)
	GetOnlineCount(context.Context) (int, error)
	GetOnlineSlice(context.Context) (*char.DBCharactersSlice, error)
	// Armory

	GetArmoryCharactersSlice(context.Context, int, int) (*char.DBCharactersSlice, int, error)
}

type NexusRepo interface {
	// Profile
	CreateProfile(context.Context, *nexus.DBProfile) (*nexus.DBProfile, error)
	GetProfile(context.Context, int) (*nexus.DBProfile, error)
	// Avatar
	UpdateAvatar(context.Context, *nexus.DBProfile) error
	ExistAvatar(context.Context, int) (bool, error)
	// Session
	GetSession(context.Context, *nexus.DBSession) (*nexus.DBSession, error)
	CreateSession(context.Context, *nexus.DBSession) error
	UpdateSession(context.Context, *nexus.DBSession) error
	DeleteSession(context.Context, string) error
	ListSession(context.Context, uint) (*nexus.DBSessionSlice, error)
	UpdateOrCreateSession(context.Context, *nexus.DBSession, string) (*nexus.DBSession, error)
	// News
	CreateNews(context.Context, *nexus.DBNews) (*nexus.DBNews, error)
	GetNewsByID(context.Context, int) (*nexus.DBNews, error)
	GetNewsSlice(context.Context, int, int) (*nexus.DBNewsSlice, int, error)
	UpdateNews(context.Context, *nexus.DBNews) (*nexus.DBNews, error)
	DeleteNews(context.Context, int) error
	// Comments
	CreateComment(context.Context, *nexus.DBComment) (*nexus.DBComment, error)
	GetCommentByID(context.Context, int) (*nexus.DBComment, error)
	DeleteComment(context.Context, int) error
	UpdateComment(context.Context, *nexus.DBComment) (*nexus.DBComment, error)
	// Like
	ToggleReaction(context.Context, *nexus.DBReaction) (*nexus.DBReaction, error)
	DeleteReaction(context.Context, *nexus.DBReaction) error
	ExistReaction(context.Context, *nexus.DBReaction) (bool, error)
	// Report Methods
	CreateReport(context.Context, *nexus.DBReport) (*nexus.DBReport, error)
	GetReportByID(context.Context, int) (*nexus.DBReport, error)
	UpdateReport(context.Context, *nexus.DBReport) (*nexus.DBReport, error)
	GetReports(context.Context, int, int) (*nexus.DBReportSlice, error)
	DeleteReportByID(context.Context, int) error
	//Enchant DBC
	GetEnchantDBCByIDs(context.Context, []int32) (map[int32]int32, error)
}
