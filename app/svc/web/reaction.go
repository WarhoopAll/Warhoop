package web

import (
	"context"
	"warhoop/app/model/nexus"
	"warhoop/app/utils"
)

func (s *WebService) ToggleReaction(ctx context.Context, entry *nexus.Reaction) (*nexus.Reaction, error) {
	res, err := s.store.NexusRepo.ToggleReaction(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}
