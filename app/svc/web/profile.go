package web

import (
	"context"
	"warhoop/app/model/nexus"
	"warhoop/app/utils"
)

func (svc *WebService) UpdateAvatar(ctx context.Context, entry *nexus.Profile) (*nexus.Profile, error) {
	err := svc.store.NexusRepo.UpdateAvatar(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return entry, nil
}

func (svc *WebService) CreateProfile(ctx context.Context, entry *nexus.Profile) (*nexus.Profile, error) {
	_, err := svc.store.NexusRepo.CreateProfile(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return nil, nil
}
