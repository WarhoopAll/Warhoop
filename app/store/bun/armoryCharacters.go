package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model"
)

func (r *CharRepo) GetArmoryCharactersSlice(ctx context.Context, limit, offset int) (*model.DBCharactersSlice, int, error) {
	var total int
	entry := &model.DBCharactersSlice{}

	countErr := r.db.NewSelect().
		Model((*model.DBCharacters)(nil)).
		ColumnExpr("count(*)").
		Scan(ctx, &total)
	if countErr != nil {
		r.logger.Error("store.SaitRepo.GetNewsSlice - count",
			log.String("err", countErr.Error()),
		)
		return nil, 0, countErr
	}

	err := r.db.NewSelect().
		Model(entry).
		OrderExpr("level DESC").
		Limit(limit).
		Offset(offset).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetNewsSlice",
			log.String("err", err.Error()),
		)
		return nil, 0, err
	}

	return entry, total, nil
}
