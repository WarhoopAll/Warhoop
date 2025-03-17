package web

import (
	"context"
	"warhoop/app/model"
	"warhoop/app/utils"
)

func (svc *WebService) UpdateAvatar(ctx context.Context, entry *model.Profile) (*model.Profile, error) {
	err := svc.store.SaitRepo.UpdateAvatar(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return entry, nil
}

func (svc *WebService) CreateProfile(ctx context.Context, entry *model.Profile) (*model.Profile, error) {
	_, err := svc.store.SaitRepo.CreateProfile(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return nil, nil
}
