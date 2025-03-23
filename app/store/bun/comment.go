package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model/nexus"
)

func (r *NexusRepo) CreateComment(ctx context.Context, entry *nexus.DBComment) (*nexus.DBComment, error) {
	_, err := r.db.
		NewInsert().
		Model(entry).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.CreateComment",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}

	entry, err = r.GetCommentByID(ctx, entry.ID)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *NexusRepo) GetCommentsByNewsID(ctx context.Context, id int) (*nexus.DBCommentSlice, error) {
	entry := &nexus.DBCommentSlice{}
	err := r.db.
		NewSelect().
		Model(entry).
		Relation("Profile").
		Where("news_id = ?", id).
		OrderExpr("like_count DESC, created_at DESC").
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetCommentsByNewsID",
			log.String("err", err.Error()),
			log.Int("news_id", id),
		)
		return nil, err
	}
	return entry, nil
}

func (r *NexusRepo) GetCommentByID(ctx context.Context, id int) (*nexus.DBComment, error) {
	entry := &nexus.DBComment{}
	err := r.db.
		NewSelect().
		Model(entry).
		Relation("Profile").
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetCommentByID",
			log.String("err", err.Error()),
			log.Int("id", id),
		)
		return nil, err
	}
	return entry, nil
}

func (r *NexusRepo) DeleteComment(ctx context.Context, id int) error {
	_, err := r.db.
		NewDelete().
		Model((*nexus.DBComment)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.DeleteComment",
			log.String("err", err.Error()),
			log.Int("id", id),
		)
		return err
	}
	return nil
}

func (r *NexusRepo) UpdateComment(ctx context.Context, entry *nexus.DBComment) (*nexus.DBComment, error) {
	_, err := r.db.
		NewUpdate().
		Model(entry).
		Set("text = ?", entry.Text).
		Where("id = ?", entry.ID).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.UpdateComment",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}

	entry, err = r.GetCommentByID(ctx, entry.ID)
	if err != nil {
		return nil, err
	}
	return entry, nil
}
