package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model"
)

type AuthRepo struct {
	db     *DB
	logger *log.Logger
	saitr  *SaitRepo
}

// NewAuthr ...
func NewAuthRepo(db *DB, logger *log.Logger, saitr *SaitRepo) *AuthRepo {
	return &AuthRepo{
		db:     db,
		logger: logger,
		saitr:  saitr,
	}
}

func (r *AuthRepo) GetByID(ctx context.Context, id int) (*model.DBAccount, error) {
	entry := &model.DBAccount{}
	err := r.db.
		NewSelect().
		Model(entry).
		Relation("Access").
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		r.logger.Error("store.AuthRepo.GetByUID",
			log.String("err", err.Error()),
			log.Int("id", id),
		)
		return nil, err
	}

	avatar, err := r.saitr.GetProfile(ctx, entry.ID)
	if err != nil {
		return nil, err
	}
	entry.Profile = avatar

	return entry, nil
}

func (r *AuthRepo) GetByUsername(ctx context.Context, username string) (*model.DBAccount, error) {
	entry := &model.DBAccount{}
	err := r.db.
		NewSelect().
		Model(entry).
		Relation("Access").
		Where("username = ?", username).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.AuthRepo.GetByUsername",
			log.String("err", err.Error()),
			log.String("username", username),
		)
		return nil, err
	}

	avatar, err := r.saitr.GetProfile(ctx, entry.ID)
	if err != nil {
		return nil, err
	}
	entry.Profile = avatar

	return entry, nil
}

func (r *AuthRepo) ExistsEmail(ctx context.Context, email string) (bool, error) {
	exist, err := r.db.
		NewSelect().
		Model((*model.DBAccount)(nil)).
		Where("email = ?", email).
		Exists(ctx)

	if err != nil {
		r.logger.Error("store.AuthRepo.ExistsEmail",
			log.String("err", err.Error()),
			log.String("email", email),
		)
		return false, err
	}
	return exist, nil
}

func (r *AuthRepo) ExistsUsername(ctx context.Context, username string) (bool, error) {
	exist, err := r.db.
		NewSelect().
		Model((*model.DBAccount)(nil)).
		Where("username = ?", username).
		Exists(ctx)

	if err != nil {
		r.logger.Error("store.AuthRepo.ExistsUsername",
			log.String("err", err.Error()),
			log.String("username", username),
		)
		return false, err
	}
	return exist, nil
}

func (r *AuthRepo) Create(ctx context.Context, entry *model.DBAccount) (*model.DBAccount, error) {
	_, err := r.db.
		NewInsert().
		Model(entry).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.AuthRepo.Create",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}
	return entry, nil
}
