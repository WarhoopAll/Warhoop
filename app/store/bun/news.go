package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model"
)

func (r *SaitRepo) CreateNews(ctx context.Context, entry *model.DBNews) (*model.DBNews, error) {
	_, err := r.db.
		NewInsert().
		Model(entry).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.CreateNews",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}

	entry, err = r.GetNewsByID(ctx, entry.ID)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *SaitRepo) GetNewsByID(ctx context.Context, id int) (*model.DBNews, error) {
	entry := &model.DBNews{}
	err := r.db.
		NewSelect().
		Model(entry).
		Relation("Profile").
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetNewsByID",
			log.String("err", err.Error()),
			log.Int("id", id),
		)
		return nil, err
	}

	comments, err := r.GetCommentsByNewsID(ctx, id)
	if err != nil {
		return nil, err
	}

	entry.Comments = comments

	return entry, nil
}

func (r *SaitRepo) GetNewsSlice(ctx context.Context, limit, offset int) (*model.DBNewsSlice, int, error) {
	var total int
	entry := &model.DBNewsSlice{}

	countErr := r.db.NewSelect().
		Model((*model.DBNews)(nil)).
		ColumnExpr("count(*)").
		Scan(ctx, &total)
	if countErr != nil {
		r.logger.Error("store.SaitRepo.GetNewsSlice - count",
			log.String("err", countErr.Error()),
		)
		return nil, 0, countErr
	}

	err := r.db.NewSelect().
		Model(entry).
		Relation("Profile").
		OrderExpr("created_at DESC").
		Limit(limit).
		Offset(offset).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetNewsSlice",
			log.String("err", err.Error()),
		)
		return nil, 0, err
	}

	return entry, total, nil
}

func (r *SaitRepo) UpdateNews(ctx context.Context, entry *model.DBNews) (*model.DBNews, error) {
	q := r.db.NewUpdate().Model(entry).Where("id = ?", entry.ID)

	if entry.Title != "" {
		q.Set("title = ?", entry.Title)
	}
	if entry.Text != "" {
		q.Set("text = ?", entry.Text)
	}
	if entry.ImageUrl != "" {
		q.Set("image_url = ?", entry.ImageUrl)
	}
	if entry.LikeCount != 0 {
		q.Set("like_count = ?", entry.LikeCount)
	}

	_, err := q.Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.UpdateNews",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}

	entry, err = r.GetNewsByID(ctx, entry.ID)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *SaitRepo) DeleteNews(ctx context.Context, id int) error {
	_, err := r.db.
		NewDelete().
		Model((*model.DBNews)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.DeleteNews",
			log.String("err", err.Error()),
			log.Int("id", id),
		)
		return err
	}
	return nil
}
