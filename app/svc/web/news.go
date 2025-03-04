package web

import (
	"context"
	"fmt"
	"time"
	"warhoop/app/model"
	"warhoop/app/utils"
)

func (svc WebService) CreateNews(ctx context.Context, id int, entry *model.News) (*model.News, error) {
	entry.ProfileID = id
	res, err := svc.store.SaitRepo.CreateNews(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc WebService) GetNewsByID(ctx context.Context, id int) (*model.News, error) {
	res, err := svc.store.SaitRepo.GetNewsByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) GetNewsSlice(ctx context.Context, limit, offset int) (*model.NewsSlice, int, error) {
	cacheKey := fmt.Sprintf("news_slice:%d:%d", limit, offset)

	var cachedData struct {
		News  model.NewsSlice `json:"news"`
		Total int             `json:"total"`
	}
	err := svc.redisCache.Get(ctx, cacheKey, &cachedData)
	if err == nil {
		return &cachedData.News, cachedData.Total, nil
	}

	entry, total, err := svc.store.SaitRepo.GetNewsSlice(ctx, limit, offset)
	if err != nil {
		return nil, 0, utils.ErrDataBase
	}

	news := entry.ToWeb()

	_ = svc.redisCache.Set(ctx, cacheKey, struct {
		News  model.NewsSlice `json:"news"`
		Total int             `json:"total"`
	}{news, total}, 30*time.Minute)

	return &news, total, nil
}

func (svc WebService) UpdateNews(ctx context.Context, entry *model.News) (*model.News, error) {
	res, err := svc.store.SaitRepo.UpdateNews(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc WebService) DeleteNews(ctx context.Context, id int) error {
	err := svc.store.SaitRepo.DeleteNews(ctx, id)
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}
