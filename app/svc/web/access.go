package web

import (
	"context"
	"grimoire/app/model"
	"grimoire/app/utils"
)

func (svc *WebService) ExistsAccess(ctx context.Context, id int) (bool, error) {
	result, err := svc.store.AuthRepo.ExistsAccess(ctx, id)
	if err != nil {
		return false, utils.ErrDataBase
	}
	return result, nil
}

func (svc *WebService) AccessByID(ctx context.Context, id int) (*model.Access, error) {
	result, err := svc.store.AuthRepo.GetAccessByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return result.ToWeb(), nil
}
