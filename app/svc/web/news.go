package web

import (
	"context"
	"warhoop/app/model/nexus"
	"warhoop/app/utils"
)

func (svc *WebService) CreateNews(ctx context.Context, id int, entry *nexus.News) (*nexus.News, error) {
	entry.ProfileID = id
	res, err := svc.store.NexusRepo.CreateNews(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) GetNewsByID(ctx context.Context, id int) (*nexus.News, error) {
	res, err := svc.store.NexusRepo.GetNewsByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) GetNewsSlice(ctx context.Context, limit, offset int) (*nexus.NewsSlice, int, error) {

	entry, total, err := svc.store.NexusRepo.GetNewsSlice(ctx, limit, offset)
	if err != nil {
		return nil, 0, utils.ErrDataBase
	}

	news := entry.ToWeb()

	return news, total, nil
}

func (svc *WebService) UpdateNews(ctx context.Context, entry *nexus.News) (*nexus.News, error) {
	res, err := svc.store.NexusRepo.UpdateNews(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) DeleteNews(ctx context.Context, id int) error {
	err := svc.store.NexusRepo.DeleteNews(ctx, id)
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}
