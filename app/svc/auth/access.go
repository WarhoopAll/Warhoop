package auth

import (
	"context"
	"warhoop/app/model/auth"
	"warhoop/app/utils"
)

func (svc *AuthService) ExistsAccess(ctx context.Context, id int) (bool, error) {
	result, err := svc.store.AuthRepo.ExistsAccess(ctx, id)
	if err != nil {
		return false, utils.ErrDataBase
	}
	return result, nil
}

func (svc *AuthService) AccessByID(ctx context.Context, id int) (*auth.Access, error) {
	result, err := svc.store.AuthRepo.GetAccessByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return result.ToWeb(), nil
}
