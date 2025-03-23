package characters

import (
	"context"
	"warhoop/app/model/char"
	"warhoop/app/utils"
)

func (svc *CharService) GetArmoryCharactersSlice(ctx context.Context, limit, offset int) (*char.CharactersSlice, int, error) {
	entry, total, err := svc.store.CharRepo.GetArmoryCharactersSlice(ctx, limit, offset)
	if err != nil {
		return nil, 0, utils.ErrDataBase
	}

	characters := entry.ToWeb()

	return characters, total, nil
}
