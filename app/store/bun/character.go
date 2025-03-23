package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model/char"
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

func (r *CharRepo) GetByID(ctx context.Context, id int) (*char.DBCharacters, error) {
	entry := &char.DBCharacters{}
	err := r.db.NewSelect().
		Model(entry).
		Where("guid = ?", id).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.CharRepo.GetByID",
			log.String("err", err.Error()),
			log.Int("guid", id),
		)
		return nil, err
	}

	return entry, nil
}

func (r *CharRepo) GetByName(ctx context.Context, name string) (*char.DBCharacters, error) {
	entry := &char.DBCharacters{}
	err := r.db.NewSelect().
		Model(entry).
		Relation("Stats").
		Relation("Inventory.ItemInstance.ItemDBC").
		Where("LOWER(name) = LOWER(?)", name).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.CharRepo.GetByName",
			log.String("err", err.Error()),
			log.String("name", name),
		)
		return nil, err
	}
	return entry, nil
}

func (r *CharRepo) GetOnlineCount(ctx context.Context) (int, error) {
	entry := &char.DBCharacters{}
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

func (r *CharRepo) GetTop10Kill(ctx context.Context) (*char.DBCharactersSlice, error) {
	entry := &char.DBCharactersSlice{}
	err := r.db.
		NewSelect().
		Model(entry).
		Where("totalkills > 0").
		OrderExpr("totalkills DESC").
		Limit(10).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.CharRepo.GetTop10Kill",
			log.String("err", err.Error()),
		)
		return nil, err
	}
	return entry, nil
}

func (r *CharRepo) GetOnlineSlice(ctx context.Context) (*char.DBCharactersSlice, error) {
	entry := &char.DBCharactersSlice{}
	err := r.db.NewSelect().
		Model(entry).
		Relation("Maps").
		Relation("Zones").
		Where("online = ?", 1).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.CharRepo.GetOnlineSlice",
			log.String("err", err.Error()),
		)
		return nil, err
	}
	return entry, nil
}
