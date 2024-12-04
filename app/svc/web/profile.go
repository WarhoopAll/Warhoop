package web

import (
	"context"
	"grimoire/app/model"
	"grimoire/app/utils"
)

func (svc WebService) UpdateAvatar(ctx context.Context, entry *model.Profile) (*model.Profile, error) {
	err := svc.store.SaitRepo.UpdateAvatar(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return entry, nil
}
