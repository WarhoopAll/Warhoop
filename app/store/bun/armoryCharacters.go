package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model/char"
)

func (r *CharRepo) GetArmoryCharactersSlice(ctx context.Context, limit, offset int) (*char.DBCharactersSlice, int, error) {
	var total int
	entry := &char.DBCharactersSlice{}

	countErr := r.db.NewSelect().
		Model((*char.DBCharacters)(nil)).
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
