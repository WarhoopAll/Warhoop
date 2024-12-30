package web

import (
	"context"
	"warhoop/app/model"
	"warhoop/app/utils"
)

func (s *WebService) ToggleReaction(ctx context.Context, entry *model.Reaction) (*model.Reaction, error) {
	res, err := s.store.SaitRepo.ToggleReaction(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}
