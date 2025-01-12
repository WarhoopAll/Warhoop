package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model"
)

type SaitRepo struct {
	db     *DB
	logger *log.Logger
}

// NewSaitRepo ...
func NewSaitRepo(db *DB, logger *log.Logger) *SaitRepo {
	return &SaitRepo{
		db:     db,
		logger: logger,
	}
}

func (r *SaitRepo) ExistAvatar(ctx context.Context, accountID int) (bool, error) {
	exists, err := r.db.
		NewSelect().
		Model((*model.DBProfile)(nil)).
		Where("account_id = ?", accountID).
		Exists(ctx)

	if err != nil {
		r.logger.Error("store.Account.IfExists",
			log.String("error", err.Error()),
			log.Int("account_id", accountID),
		)
		return false, err
	}
	return exists, nil
}

func (r *SaitRepo) UpdateAvatar(ctx context.Context, entry *model.DBProfile) error {
	exists, err := r.ExistAvatar(ctx, entry.AccountID)
	if err != nil {
		return err
	}

	if exists {
		_, err = r.db.
			NewUpdate().
			Model(entry).
			Set("avatar = ?", entry.Avatar).
			Where("account_id = ?", entry.AccountID).
			Exec(ctx)
		if err != nil {
			r.logger.Error("store.Account.UpdateAvatar",
				log.String("error", err.Error()),
				log.Int("account_id", entry.AccountID),
				log.String("Avatar", entry.Avatar),
			)
			return err
		}
	} else {
		_, err = r.db.NewInsert().
			Model(entry).
			Exec(ctx)
		if err != nil {
			r.logger.Error("store.Account.InsertAvatar",
				log.String("error", err.Error()),
				log.Int("account_id", entry.AccountID),
				log.String("Avatar", entry.Avatar),
			)
			return err
		}
	}
	return nil
}
