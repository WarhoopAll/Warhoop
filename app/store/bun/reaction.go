package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model"
)

func (r *SaitRepo) ExistReaction(ctx context.Context, entry *model.DBReaction) (bool, error) {
	exists, err := r.db.NewSelect().
		Model(entry).
		Where("user_id = ? AND object_type = ? AND object_id = ? AND reaction_type = ?", entry.UserID, entry.ObjectType, entry.ObjectID, entry.ReactionType).
		Exists(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.ExistReaction",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return false, err
	}
	return exists, nil
}

func (r *SaitRepo) CreateReaction(ctx context.Context, entry *model.DBReaction) error {
	_, err := r.db.NewInsert().
		Model(entry).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.CreateReaction",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return err
	}

	return nil
}

func (r *SaitRepo) DeleteReaction(ctx context.Context, entry *model.DBReaction) error {
	_, err := r.db.NewDelete().
		Model((*model.DBReaction)(nil)).
		Where("user_id = ? AND object_type = ? AND object_id = ? AND reaction_type = ?", entry.UserID, entry.ObjectType, entry.ObjectID, entry.ReactionType).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.DeleteReaction",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return err
	}
	return nil
}

func (r *SaitRepo) ToggleReaction(ctx context.Context, entry *model.DBReaction) (*model.DBReaction, error) {
	exists, err := r.ExistReaction(ctx, entry)
	if err != nil {
		return nil, err
	}

	if exists {
		err = r.DeleteReaction(ctx, entry)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		err = r.CreateReaction(ctx, entry)
		if err != nil {
			return nil, err
		}
	}
	return entry, nil

}
