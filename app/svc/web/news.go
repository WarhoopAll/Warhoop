package web

import (
	"context"
	"warhoop/app/model"
)

func (svc WebService) CreateNews(ctx context.Context, entry *model.News) (*model.News, error) {
	res, err := svc.store.SaitRepo.CreateNews(ctx, entry.ToDB())
	if err != nil {
		return nil, err
	}
	return res.ToWeb(), nil
}

func (svc WebService) GetNewsByID(ctx context.Context, id int) (*model.News, error) {
	res, err := svc.store.SaitRepo.GetNewsByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.ToWeb(), nil
}

func (svc WebService) GetNewsSlice(ctx context.Context) (*model.NewsSlice, error) {
	entry, err := svc.store.SaitRepo.GetNewsSlice(ctx)
	if err != nil {
		return nil, err
	}
	news := entry.ToWeb()
	return &news, nil
}

func (svc WebService) UpdateNews(ctx context.Context, entry *model.News) error {
	err := svc.store.SaitRepo.UpdateNews(ctx, entry.ToDB())
	if err != nil {
		return err
	}
	return nil
}

func (svc WebService) DeleteNews(ctx context.Context, id int) error {
	err := svc.store.SaitRepo.DeleteNews(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
