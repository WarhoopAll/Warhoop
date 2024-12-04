package web

import (
	"context"
	"grimoire/app/model"
)

func (svc WebService) GetNewsSlice(ctx context.Context) (*model.NewsSlice, error) {
	entry, err := svc.store.SaitRepo.GetNewsSlice(ctx)
	if err != nil {
		return nil, err
	}
	news := entry.ToWeb()
	return &news, nil
}

func (svc WebService) GetNewsByID(ctx context.Context, entry *model.News) (*model.News, error) {
	res, err := svc.store.SaitRepo.GetNewsByID(ctx, entry.ToDB())
	if err != nil {
		return nil, err
	}
	return res.ToWeb(), nil
}
