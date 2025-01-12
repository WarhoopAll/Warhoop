package bun

import (
	"context"
	"errors"
	"warhoop/app/log"
	"warhoop/app/model"
	"warhoop/app/utils"
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

func (r *SaitRepo) GetProfile(ctx context.Context, id int) (*model.DBProfile, error) {
	avatar := &model.DBProfile{}
	err := r.db.
		NewSelect().
		Model(avatar).
		Where("account_id = ?", id).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, utils.ErrNoRows) {
			return nil, nil
		}
		r.logger.Error("store.Account.GetAvatarIfExists",
			log.String("error", err.Error()),
			log.Int("account_id", id),
		)
		return nil, err
	}
	return avatar, nil
}
