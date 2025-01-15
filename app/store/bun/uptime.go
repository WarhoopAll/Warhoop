package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model"
)

func (r *AuthRepo) GetUptimeByID(ctx context.Context, id int) (*model.DBUptime, error) {
	entry := &model.DBUptime{}
	err := r.db.NewSelect().
		Model(entry).
		Where("realmid = ?", id).
		Relation("Realm").
		Order("starttime DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.AuthRepo.GetUptimeByID",
			log.String("err", err.Error()),
			log.Int("realmid", id),
		)
	}
	return entry, err
}
