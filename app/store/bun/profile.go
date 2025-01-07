package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model"
)

func (r *SaitRepo) CreateProfile(ctx context.Context, entry *model.DBProfile) (*model.DBProfile, error) {
	_, err := r.db.
		NewInsert().
		Model(entry).
		Column("account_id").
		Column("name").
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.Profile.Create",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}
	return nil, nil
}
