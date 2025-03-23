package bun

import (
	"context"
	"database/sql"
	"warhoop/app/log"
	"warhoop/app/model/nexus"
	"warhoop/app/utils"
)

func (r *NexusRepo) GetSession(ctx context.Context, entry *nexus.DBSession) (*nexus.DBSession, error) {
	err := r.db.
		NewSelect().
		Model(entry).
		Where("token = ?", entry.Token).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Warn("store.SaitRepo.GetSession",
				log.String("err", err.Error()),
				log.Object("entry", entry),
			)
			return nil, err
		}
		r.logger.Error("store.SaitRepo.GetSession",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}

	return entry, nil
}

func (r *NexusRepo) CreateSession(ctx context.Context, entry *nexus.DBSession) error {
	_, err := r.db.
		NewInsert().
		Model(entry).
		Returning("*").
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.CreateSession",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return err
	}

	return nil
}

func (r *NexusRepo) UpdateSession(ctx context.Context, entry *nexus.DBSession) error {
	_, err := r.db.
		NewUpdate().
		Model(entry).
		Where("account_id = ? AND fingerprint = ?", entry.AccountID, entry.Finger).
		Set("token = ?", entry.Token).
		Returning("*").
		OmitZero().
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.UpdateSession",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return err
	}

	return nil
}

func (r *NexusRepo) DeleteSession(ctx context.Context, id string) error {
	_, err := r.db.
		NewDelete().
		Model((*nexus.DBSession)(nil)).
		Where("token = ?", id).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.DeleteSession",
			log.String("err", err.Error()),
			log.String("token", id),
		)
		return err
	}

	return nil
}

func (r *NexusRepo) ListSession(ctx context.Context, id uint) (*nexus.DBSessionSlice, error) {
	entry := &nexus.DBSessionSlice{}
	err := r.db.
		NewSelect().
		Model(entry).
		Relation("Account").
		Where("account_id = ?", id).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.ListSession",
			log.String("err", err.Error()),
			log.Uint("account_id", id),
		)
		return nil, err
	}
	return entry, nil
}

func (r *NexusRepo) ExistSession(ctx context.Context, entry *nexus.DBSession) (bool, error) {
	exists, err := r.db.
		NewSelect().
		Model(entry).
		Where("account_id = ? AND fingerprint =? AND expired_at > NOW()", entry.AccountID, entry.Finger).
		Exists(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.HasSession",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return false, err
	}
	return exists, nil
}

func (r *NexusRepo) UpdateOrCreateSession(ctx context.Context, entry *nexus.DBSession, newToken string) (*nexus.DBSession, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error("store.SaitRepo.UpdateOrCreateSession",
			log.String("err", err.Error()),
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
		r.logger.Error("store.SaitRepo.UpdateOrCreateSession - select",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}

	if !exists {
		return nil, utils.ErrNoRows
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
		r.logger.Error("store.SaitRepo.UpdateOrCreateSession - update",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}

	return entry, nil
}
