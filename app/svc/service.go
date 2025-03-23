package svc

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"warhoop/app/model/auth"
	"warhoop/app/model/char"
	"warhoop/app/model/nexus"
)

type Auth interface {
	// Account
	GetByID(context.Context, int) (*auth.Account, error)
	// Auth
	SignIn(context.Context, *auth.Account) (*auth.Account, error)
	SignUp(context.Context, *auth.Account) (*auth.Account, error)
	// Access
	ExistsAccess(context.Context, int) (bool, error)
	AccessByID(context.Context, int) (*auth.Access, error)
	// Uptime
	GetUptime(context.Context) (*auth.Uptime, error)
	CheckOnlineStatus(string, int16) bool
}

type Characters interface {
	// Characters
	GetByID(context.Context, int) (*char.Characters, error)
	GetByName(context.Context, string) (*char.Characters, error)
	GetTop10Kill(context.Context) ([]map[string]interface{}, error)
	GetOnlineCount(context.Context) (int, error)
	GetOnlineSlice(context.Context) ([]map[string]interface{}, error)
	// Armory
	GetArmoryCharactersSlice(ctx context.Context, limit, offset int) (*char.CharactersSlice, int, error)
}

type Web interface {
	// Profile
	CreateProfile(context.Context, *nexus.Profile) (*nexus.Profile, error)
	UpdateAvatar(context.Context, *nexus.Profile) (*nexus.Profile, error)
	// Session
	CreateCookie(string) *fiber.Cookie
	GenerateAccessToken(int) (string, error)
	GetSession(context.Context, *nexus.Session) (*nexus.Session, error)
	DeleteSession(context.Context, string) (*fiber.Cookie, error)
	ExistSession(context.Context, *nexus.Session) (bool, error)
	CreateSession(context.Context, *nexus.Session) error
	UpdateSession(context.Context, *nexus.Session) error
	UpdateOrCreateSession(context.Context, *nexus.Session, string) error
	HandleSession(context.Context, *nexus.Session) (string, error)
	// News
	CreateNews(context.Context, int, *nexus.News) (*nexus.News, error)
	GetNewsSlice(context.Context, int, int) (*nexus.NewsSlice, int, error)
	GetNewsByID(context.Context, int) (*nexus.News, error)
	UpdateNews(context.Context, *nexus.News) (*nexus.News, error)
	DeleteNews(context.Context, int) error
	// Comments
	CreateComment(context.Context, int, *nexus.Comment) (*nexus.Comment, error)
	DeleteComment(context.Context, int) error
	UpdateComment(context.Context, int, *nexus.Comment) (*nexus.Comment, error)
	GetCommentByID(context.Context, int) (*nexus.Comment, error)
	GetCommentByNewsID(context.Context, int) (*nexus.CommentSlice, error)
	// Like
	ToggleReaction(context.Context, *nexus.Reaction) (*nexus.Reaction, error)
	// Report
	CreateReport(context.Context, int, *nexus.Report) (*nexus.Report, error)
	GetReportByID(context.Context, int) (*nexus.Report, error)
	GetReports(context.Context, int, int) (*nexus.ReportSlice, error)
	DeleteReportByID(context.Context, int) error
}

type Soap interface {
	ExecuteCommand(string) (string, error)
}
