package svc

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"warhoop/app/model"
)

type Auth interface {
	// Account
	GetByID(context.Context, int) (*model.Account, error)
	// Auth
	SignIn(context.Context, *model.Account) (*model.Account, error)
	SignUp(context.Context, *model.Account) (*model.Account, error)
	// Access
	ExistsAccess(context.Context, int) (bool, error)
	AccessByID(context.Context, int) (*model.Access, error)
	// Uptime
	GetUptime(context.Context) (*model.Uptime, error)
	CheckOnlineStatus(string, int16) bool
}

type Characters interface {
	GetByID(context.Context, int) (*model.Characters, error)
	GetByName(context.Context, string) (*model.Characters, error)
	GetTop10Kill(context.Context) ([]map[string]interface{}, error)
	GetOnlineCount(context.Context) (int, error)
	GetOnlineSlice(context.Context) ([]map[string]interface{}, error)
}

type Web interface {
	// Profile
	CreateProfile(context.Context, *model.Profile) (*model.Profile, error)
	UpdateAvatar(context.Context, *model.Profile) (*model.Profile, error)
	// Session
	CreateCookie(string) *fiber.Cookie
	GenerateAccessToken(int) (string, error)
	GetSession(context.Context, *model.Session) (*model.Session, error)
	DeleteSession(context.Context, string) (*fiber.Cookie, error)
	ExistSession(context.Context, *model.Session) (bool, error)
	CreateSession(context.Context, *model.Session) error
	UpdateSession(context.Context, *model.Session) error
	UpdateOrCreateSession(context.Context, *model.Session, string) error
	HandleSession(context.Context, *model.Session) (string, error)
	// News
	CreateNews(context.Context, *model.News) (*model.News, error)
	GetNewsSlice(context.Context, int, int) (*model.NewsSlice, int, error)
	GetNewsByID(context.Context, int) (*model.News, error)
	UpdateNews(context.Context, *model.News) (*model.News, error)
	DeleteNews(context.Context, int) error
	// Comments
	CreateComment(context.Context, *model.Comment) (*model.Comment, error)
	DeleteComment(context.Context, int) error
	UpdateComment(context.Context, *model.Comment) (*model.Comment, error)
	GetCommentByID(context.Context, int) (*model.Comment, error)
	GetCommentByNewsID(context.Context, int) (*model.CommentSlice, error)
	// Like
	ToggleReaction(context.Context, *model.Reaction) (*model.Reaction, error)
}

type Soap interface {
	ExecuteCommand(command string) (string, error)
}
