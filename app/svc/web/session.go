package web

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"grimoire/app/model"
	"grimoire/app/utils"
)

func (svc *WebService) GetSession(ctx context.Context, entry *model.Session) (*model.Session, error) {
	session, err := svc.store.SaitRepo.GetSession(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return session.ToWeb(), nil
}

func (svc WebService) DeleteSession(ctx context.Context, id string) (*fiber.Cookie, error) {
	err := svc.store.SaitRepo.DeleteSession(ctx, id)
	if err != nil {
		return nil, utils.ErrBadSession
	}

	cookie := DeleteCookie()
	return cookie, nil
}

func (svc WebService) ExistSession(ctx context.Context, entry *model.Session) (bool, error) {
	exist, err := svc.store.SaitRepo.ExistSession(ctx, entry.ToDB())
	if err != nil {
		return !exist, utils.ErrDataBase
	}
	return exist, nil
}

func (svc WebService) CreateSession(ctx context.Context, entry *model.Session) error {
	err := svc.store.SaitRepo.CreateSession(ctx, entry.ToDB())
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}

func (svc WebService) UpdateSession(ctx context.Context, entry *model.Session) error {
	err := svc.store.SaitRepo.UpdateSession(ctx, entry.ToDB())
	if err != nil {
		return err
	}
	return nil
}

func (svc WebService) UpdateOrCreateSession(ctx context.Context, entry *model.Session, newToken string) error {
	_, err := svc.store.SaitRepo.UpdateOrCreateSession(ctx, entry.ToDB(), newToken)
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}

func (svc *WebService) HandleSession(ctx context.Context, session *model.Session) (string, error) {
	exists, err := svc.ExistSession(ctx, session)
	if err != nil {
		return "", err
	}

	if exists {
		err = svc.UpdateSession(ctx, session)
	} else {
		err = svc.CreateSession(ctx, session)
	}

	if err != nil {
		return "", err
	}
	return session.Token, nil
}
