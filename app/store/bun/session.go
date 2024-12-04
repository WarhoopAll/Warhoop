package bun

import (
	"context"
	"database/sql"
	"grimoire/app/log"
	"grimoire/app/model"
	"grimoire/app/utils"
)

func (repo *SaitRepo) GetSession(ctx context.Context, entry *model.DBSession) (*model.DBSession, error) {
	err := repo.db.
		NewSelect().
		Model(entry).
		Where("token = ?", entry.Token).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			repo.logger.Warn("store.Serve.GetSession",
				log.String("err", err.Error()),
				log.String("token", entry.Token),
			)
			return nil, err
		}
		repo.logger.Error("store.Serve.GetSession",
			log.String("err", err.Error()),
			log.String("token", entry.Token),
		)
		return nil, err
	}

	return entry, nil
}

func (repo *SaitRepo) CreateSession(ctx context.Context, entry *model.DBSession) error {
	_, err := repo.db.
		NewInsert().
		Model(entry).
		Returning("*").
		Exec(ctx)

	if err != nil {
		repo.logger.Error("store.Serve.CreateSession",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return err
	}

	return nil
}

func (repo *SaitRepo) UpdateSession(ctx context.Context, entry *model.DBSession) error {
	_, err := repo.db.
		NewUpdate().
		Model(entry).
		Where("account_id = ? AND fingerprint = ?", entry.AccountID, entry.Finger).
		Set("token = ?", entry.Token).
		Returning("*").
		OmitZero().
		Exec(ctx)
	if err != nil {
		repo.logger.Error("store.Serve.UpdateSession",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return err
	}

	return nil
}

func (repo *SaitRepo) DeleteSession(ctx context.Context, id string) error {
	_, err := repo.db.
		NewDelete().
		Model((*model.DBSession)(nil)).
		Where("token = ?", id).
		Exec(ctx)
	if err != nil {
		repo.logger.Error("store.Serve.DeleteSession",
			log.String("err", err.Error()),
			log.String("id", id),
		)
		return err
	}

	return nil
}

func (repo *SaitRepo) ListSession(ctx context.Context, id uint) (*model.DBSessionSlice, error) {
	entry := &model.DBSessionSlice{}
	err := repo.db.
		NewSelect().
		Model(entry).
		Relation("Account").
		Where("account_id = ?", id).
		Scan(ctx)
	if err != nil {
		repo.logger.Error("store.Serve.ListSession",
			log.String("err", err.Error()),
			log.Uint("id", id),
		)
		return nil, err
	}
	return entry, nil
}

func (repo *SaitRepo) ExistSession(ctx context.Context, entry *model.DBSession) (bool, error) {
	exists, err := repo.db.
		NewSelect().
		Model(entry).
		Where("account_id = ? AND fingerprint =? AND expired_at > NOW()", entry.AccountID, entry.Finger).
		Exists(ctx)
	if err != nil {
		repo.logger.Error("store.Serve.HasSession",
			log.String("err", err.Error()),
		)
		return false, err
	}
	return exists, nil
}

func (repo *SaitRepo) UpdateOrCreateSession(ctx context.Context, entry *model.DBSession, newToken string) (*model.DBSession, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		repo.logger.Error("store.Serve.UpdateOrCreateSession",
			log.String("failed transaction", err.Error()),
		)
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	exists, err := tx.NewSelect().
		Model(entry).
		Where("account_id = ? AND token = ? AND fingerprint = ?", entry.AccountID, entry.Token, entry.Finger).
		For("UPDATE").
		Exists(ctx)
	if err != nil {
		repo.logger.Error("store.Serve.UpdateOrCreateSession",
			log.String("failed check session existence", err.Error()),
		)
		return nil, err
	}

	if !exists {
		return nil, utils.ErrDataBase
	}

	_, err = tx.NewUpdate().
		Model(entry).
		Where("account_id = ? AND token = ?", entry.AccountID, entry.Token).
		Set("token = ?", newToken).
		Set("updated_at = NOW()").
		Set("expired_at = ?", entry.ExpiredAt).
		Returning("*").
		Exec(ctx)
	if err != nil {
		repo.logger.Error("store.Serve.UpdateOrCreateSession",
			log.String("failed to update session", err.Error()),
		)
		return nil, err
	}

	return entry, nil
}
