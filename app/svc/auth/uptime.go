package auth

import (
	"context"
	"warhoop/app/model"
	"warhoop/app/utils"
)

func (s *AuthService) GetUptimeByID(ctx context.Context, id int) (*model.Uptime, error) {
	entry, err := s.store.AuthRepo.GetUptimeByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return entry.ToWeb(), nil
}
