package web

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"warhoop/app/model/nexus"
	"warhoop/app/utils"
)

func (svc *WebService) GetSession(ctx context.Context, entry *nexus.Session) (*nexus.Session, error) {
	session, err := svc.store.NexusRepo.GetSession(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return session.ToWeb(), nil
}

func (svc *WebService) DeleteSession(ctx context.Context, id string) (*fiber.Cookie, error) {
	err := svc.store.NexusRepo.DeleteSession(ctx, id)
	if err != nil {
		return nil, utils.ErrBadSession
	}

	cookie := DeleteCookie()
	return cookie, nil
}

func (svc *WebService) ExistSession(ctx context.Context, entry *nexus.Session) (bool, error) {
	exist, err := svc.store.NexusRepo.ExistSession(ctx, entry.ToDB())
	if err != nil {
		return !exist, utils.ErrDataBase
	}
	return exist, nil
}

func (svc *WebService) CreateSession(ctx context.Context, entry *nexus.Session) error {
	err := svc.store.NexusRepo.CreateSession(ctx, entry.ToDB())
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}

func (svc *WebService) UpdateSession(ctx context.Context, entry *nexus.Session) error {
	err := svc.store.NexusRepo.UpdateSession(ctx, entry.ToDB())
	if err != nil {
		return err
	}
	return nil
}

func (svc *WebService) UpdateOrCreateSession(ctx context.Context, entry *nexus.Session, newToken string) error {
	_, err := svc.store.NexusRepo.UpdateOrCreateSession(ctx, entry.ToDB(), newToken)
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}

func (svc *WebService) HandleSession(ctx context.Context, session *nexus.Session) (string, error) {
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
