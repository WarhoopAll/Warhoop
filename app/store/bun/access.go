package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model/auth"
)

func (r *AuthRepo) ExistsAccess(ctx context.Context, id int) (bool, error) {
	exist, err := r.db.
		NewSelect().
		Model((*auth.DBAccess)(nil)).
		Where("AccountID = ?", id).
		Exists(ctx)

	if err != nil {
		r.logger.Error("store.AuthRepo.ExistsAccess",
			log.String("err", err.Error()),
			log.Int("AccountID", id),
		)
		return false, err
	}
	return exist, nil
}

func (r *AuthRepo) GetAccessByID(ctx context.Context, id int) (*auth.DBAccess, error) {
	entry := &auth.DBAccess{}
	err := r.db.
		NewSelect().
		Model(entry).
		Where("AccountID = ?", id).
		Scan(ctx)

	if err != nil {
		r.logger.Error("store.SaitRepo.GetAccessByID",
			log.String("err", err.Error()),
			log.Int("AccountID", id),
		)
		return nil, err
	}
	return entry, nil
}
