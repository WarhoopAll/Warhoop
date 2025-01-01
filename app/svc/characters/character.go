package characters

import (
	"context"
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
	result, err := svc.store.CharRepo.GetByName(ctx, name)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	return result.ToWeb(), nil
}

func (svc *CharService) GetTop10Kill(ctx context.Context) ([]map[string]interface{}, error) {
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

	return transformed, nil
}

func (svc *CharService) GetOnlineCount(ctx context.Context) (int, error) {
	count, err := svc.store.CharRepo.GetOnlineCount(ctx)
	if err != nil {
		return 0, utils.ErrDataBase
	}
	return count, err
}
