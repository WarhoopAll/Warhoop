package bun

import (
	"context"
	"grimoire/app/log"
	"grimoire/app/model"
)

func (r *SaitRepo) GetCommentByNewsID(ctx context.Context, id int) (*model.DBCommentSlice, error) {
	entry := &model.DBCommentSlice{}
	err := r.db.
		NewSelect().
		Model(entry).
		Where("news_id = ?", id).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetNewsSlice",
			log.String("error", err.Error()),
			log.Int("id_news", id),
			log.Object("entries", entry),
		)
		return nil, err
	}
	return entry, nil
}

func (r *SaitRepo) CreateComment(ctx context.Context, entry *model.DBComment) (*model.DBComment, error) {
	_, err := r.db.
		NewInsert().
		Model(entry).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.AddComment",
			log.String("error", err.Error()),
			log.Int("comment_id", entry.NewsID),
			log.Object("entry", entry),
		)
		return nil, err
	}

	err = r.db.
		NewSelect().
		Model(entry).
		Relation("Profile").
		Where("id = ?", entry.ID).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.AddComment (SELECT)",
			log.String("error", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}
	return entry, nil
}

func (r *SaitRepo) DeleteComment(ctx context.Context, id int) error {
	_, err := r.db.
		NewDelete().
		Model((*model.DBComment)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetNewsSlice",
			log.String("error", err.Error()),
			log.Int("comment_id", id),
		)
		return err
	}
	return nil
}

func (r *SaitRepo) UpdateComment(ctx context.Context, entry *model.DBComment) error {
	_, err := r.db.
		NewUpdate().
		Model(entry).
		Column("text").
		Where("id = ?", entry.ID).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetNewsSlice",
			log.String("error", err.Error()),
			log.Int("comment_id", entry.ID),
		)
		return err
	}
	return nil
}

func (r *SaitRepo) GetCommentByID(ctx context.Context, id int) (*model.DBComment, error) {
	entry := &model.DBComment{}
	err := r.db.
		NewSelect().
		Model(entry).
		Relation("Profile").
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetNewsSlice",
			log.String("error", err.Error()),
			log.Int("comment_id", id),
			log.Object("entries", entry),
		)
		return nil, err
	}
	return entry, nil
}
