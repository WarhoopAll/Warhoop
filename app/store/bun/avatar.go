package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model/nexus"
)

type NexusRepo struct {
	db     *DB
	logger *log.Logger
}

// NewSaitRepo ...
func NewNexusRepo(db *DB, logger *log.Logger) *NexusRepo {
	return &NexusRepo{
		db:     db,
		logger: logger,
	}
}

func (r *NexusRepo) ExistAvatar(ctx context.Context, accountID int) (bool, error) {
	exists, err := r.db.
		NewSelect().
		Model((*nexus.DBProfile)(nil)).
		Where("account_id = ?", accountID).
		Exists(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.ExistAvatar",
			log.String("err", err.Error()),
			log.Int("account_id", accountID),
		)
		return false, err
	}
	return exists, nil
}

func (r *NexusRepo) UpdateAvatar(ctx context.Context, entry *nexus.DBProfile) error {
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
			r.logger.Error("store.SaitRepo.UpdateAvatar",
				log.String("err", err.Error()),
				log.Object("entry", entry),
			)
			return err
		}
	} else {
		_, err = r.db.NewInsert().
			Model(entry).
			Exec(ctx)
		if err != nil {
			r.logger.Error("store.SaitRepo.UpdateAvatar",
				log.String("err", err.Error()),
				log.Object("entry", entry),
			)
			return err
		}
	}
	return nil
}
