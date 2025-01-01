package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model"
)

type CharRepo struct {
	db     *DB
	logger *log.Logger
}

// NewSaitRepo ...
func NewCharRepo(db *DB, logger *log.Logger) *CharRepo {
	return &CharRepo{
		db:     db,
		logger: logger,
	}
}

func (r *CharRepo) GetByID(ctx context.Context, id int) (*model.DBCharacters, error) {
	entry := &model.DBCharacters{}
	err := r.db.NewSelect().
		Model(entry).
		Where("guid = ?", id).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.CharRepo.GetByID",
			log.String("err", err.Error()),
			log.Int("id", id),
		)
		return nil, err
	}

	return entry, nil
}

func (r *CharRepo) GetByName(ctx context.Context, name string) (*model.DBCharacters, error) {
	entry := &model.DBCharacters{}
	err := r.db.NewSelect().
		Model(entry).
		Where("name = ?", name).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.CharRepo.GetByName",
			log.String("err", err.Error()),
			log.String("id", name),
		)
		return nil, err
	}
	return entry, nil
}

func (r *CharRepo) GetCharOnline(ctx context.Context) (int, error) {
	entry := &model.DBCharacters{}
	count, err := r.db.NewSelect().
		Model(entry).
		Where("online = ?", 1).
		Count(ctx)
	if err != nil {
		r.logger.Error("store.CharRepo.GetCharOnline",
			log.String("err", err.Error()),
		)
		return 0, err
	}
	return count, nil
}

func (r *CharRepo) GetCharTop10Kill(ctx context.Context) (*model.DBCharactersSlice, error) {
	entry := &model.DBCharactersSlice{}
	err := r.db.
		NewSelect().
		Model(entry).
		Where("totalkills > 0").
		OrderExpr("totalkills DESC").
		Limit(10).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.CharRepo.GetCharactersSlice",
			log.String("error", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}
	return entry, nil
}
