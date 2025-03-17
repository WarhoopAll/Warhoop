package web

import (
	"context"
	"warhoop/app/model"
	"warhoop/app/utils"
)

func (svc *WebService) CreateNews(ctx context.Context, id int, entry *model.News) (*model.News, error) {
	entry.ProfileID = id
	res, err := svc.store.SaitRepo.CreateNews(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) GetNewsByID(ctx context.Context, id int) (*model.News, error) {
	res, err := svc.store.SaitRepo.GetNewsByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) GetNewsSlice(ctx context.Context, limit, offset int) (*model.NewsSlice, int, error) {

	entry, total, err := svc.store.SaitRepo.GetNewsSlice(ctx, limit, offset)
	if err != nil {
		return nil, 0, utils.ErrDataBase
	}

	news := entry.ToWeb()

	return news, total, nil
}

func (svc *WebService) UpdateNews(ctx context.Context, entry *model.News) (*model.News, error) {
	res, err := svc.store.SaitRepo.UpdateNews(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) DeleteNews(ctx context.Context, id int) error {
	err := svc.store.SaitRepo.DeleteNews(ctx, id)
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}
