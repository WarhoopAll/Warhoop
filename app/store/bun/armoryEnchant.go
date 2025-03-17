package bun

import (
	"context"
	"github.com/uptrace/bun"
	"warhoop/app/model"
)

func (r *SaitRepo) GetEnchantDBCByIDs(ctx context.Context, ids []int32) (map[int32]int32, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	entry := &model.DBEnchantDBCSlice{}

	err := r.db.NewSelect().
		Model(entry).
		Where("ID IN (?)", bun.In(ids)).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	gemMap := make(map[int32]int32)
	for _, res := range *entry {
		if res.SrcItemID != 0 {
			gemMap[res.ID] = res.SrcItemID
		}
	}

	return gemMap, nil
}
