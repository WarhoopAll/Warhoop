package bun

import (
	"context"
	"grimoire/app/log"
	"grimoire/app/model"
)

func (r *SaitRepo) GetNewsSlice(ctx context.Context) (*model.DBNewsSlice, error) {
	var entry model.DBNewsSlice
	err := r.db.
		NewSelect().
		Model(&entry).
		Relation("Profile").
		OrderExpr("created_at DESC").
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetNewsSlice",
			log.String("error", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}
	return &entry, nil
}

func (r *SaitRepo) GetNewsByID(ctx context.Context, entry *model.DBNews) (*model.DBNews, error) {
	err := r.db.
		NewSelect().
		Model(entry).
		Relation("Profile").
		Where("news.id = ?", entry.ID).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetNewsByID",
			log.String("error", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}

	var comments model.DBCommentSlice
	err = r.db.
		NewSelect().
		Model(&comments).
		Where("news_id = ?", entry.ID).
		Relation("Profile").
		OrderExpr("like_count DESC, created_at DESC").
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetComments",
			log.String("error", err.Error()),
			log.Int("news_id", entry.ID),
		)
		return nil, err
	}

	entry.Comments = comments

	return entry, nil
}
