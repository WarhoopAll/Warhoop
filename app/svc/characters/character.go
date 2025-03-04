package characters

import (
	"context"
	"log"
	"time"
	"warhoop/app/model"
	"warhoop/app/utils"
)

func (svc *CharService) GetByID(ctx context.Context, id int) (*model.Characters, error) {
	result, err := svc.store.CharRepo.GetByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	return result.ToWeb(), nil
}

func (svc *CharService) GetByName(ctx context.Context, name string) (*model.Characters, error) {
	var cachedData model.Characters
	cacheKey := "GetByName:" + name

	err := svc.redisCache.Get(ctx, cacheKey, &cachedData)
	if err == nil {
		return &cachedData, nil
	}

	result, err := svc.store.CharRepo.GetByName(ctx, name)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	go func() {
		cacheErr := svc.redisCache.Set(ctx, cacheKey, result, time.Minute)
		if cacheErr != nil {
			log.Printf("err: %v", cacheErr)
		}
	}()

	return result.ToWeb(), nil
}

func (svc *CharService) GetTop10Kill(ctx context.Context) ([]map[string]interface{}, error) {
	cacheKey := "GetTop10Kill"

	var cachedData []map[string]interface{}
	err := svc.redisCache.Get(ctx, cacheKey, &cachedData)
	if err == nil {
		return cachedData, nil
	}

	result, err := svc.store.CharRepo.GetTop10Kill(ctx)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	transformed := make([]map[string]interface{}, 0, len(*result))
	for _, char := range *result {
		transformed = append(transformed, map[string]interface{}{
			"name":       char.Name,
			"race":       char.Race,
			"class":      char.Class,
			"guid":       char.Guid,
			"totalkills": char.TotalKills,
		})
	}

	go func() {
		cacheErr := svc.redisCache.Set(ctx, cacheKey, transformed, 10*time.Minute)
		if cacheErr != nil {
			log.Printf("err: %v", cacheErr)
		}
	}()

	return transformed, nil
}

func (svc *CharService) GetOnlineCount(ctx context.Context) (int, error) {
	count, err := svc.store.CharRepo.GetOnlineCount(ctx)
	if err != nil {
		return 0, utils.ErrDataBase
	}
	return count, err
}

func (svc *CharService) GetOnlineSlice(ctx context.Context) ([]map[string]interface{}, error) {
	cacheKey := "GetOnlineSlice"

	var cachedData []map[string]interface{}
	err := svc.redisCache.Get(ctx, cacheKey, &cachedData)
	if err == nil {
		return cachedData, nil
	}

	result, err := svc.store.CharRepo.GetOnlineSlice(ctx)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	transformed := make([]map[string]interface{}, 0, len(*result))
	for _, char := range *result {
		transformed = append(transformed, map[string]interface{}{
			"name":   char.Name,
			"level":  char.Level,
			"race":   char.Race,
			"class":  char.Class,
			"gender": char.Gender,
			"map":    char.Maps.ToWeb(),
			"zone":   char.Zones.ToWeb(),
		})
	}

	go func() {
		cacheErr := svc.redisCache.Set(ctx, cacheKey, transformed, 10*time.Minute)
		if cacheErr != nil {
			log.Printf("err: %v", cacheErr)
		}
	}()

	return transformed, err
}
