package bun

import (
	"context"
	"grimoire/app/log"
	"grimoire/app/model"
)

func (r *SaitRepo) ExistReaction(ctx context.Context, entry *model.DBReaction) (bool, error) {
	exists, err := r.db.NewSelect().
		Model(entry).
		Where("user_id = ? AND object_type = ? AND object_id = ?", entry.UserID, entry.ObjectType, entry.ObjectID).
		Exists(ctx)
	if err != nil {
		r.logger.Error("Failed to check like existence",
			log.Int("user_id", entry.UserID),
			log.Int("object_type", entry.ObjectType),
			log.Int("object_id", entry.ObjectID),
			log.String("error", err.Error()))
		return false, err
	}
	return exists, nil
}

func (r *SaitRepo) CreateReaction(ctx context.Context, entry *model.DBReaction) error {
	_, err := r.db.NewInsert().
		Model(entry).
		Exec(ctx)
	if err != nil {
		r.logger.Error("Failed to add like",
			log.Int("user_id", entry.UserID),
			log.Int("object_type", entry.ObjectType),
			log.Int("object_id", entry.ObjectID),
			log.String("error", err.Error()))
		return err
	}

	return nil
}

func (r *SaitRepo) DeleteReaction(ctx context.Context, entry *model.DBReaction) error {
	_, err := r.db.NewDelete().
		Model((*model.DBReaction)(nil)).
		Where("user_id = ? AND object_type = ? AND object_id = ?", entry.UserID, entry.ObjectType, entry.ObjectID).
		Exec(ctx)
	if err != nil {
		r.logger.Error("Failed to remove like",
			log.Int("user_id", entry.UserID),
			log.Int("object_type", entry.ObjectType),
			log.Int("object_id", entry.ObjectID),
			log.String("error", err.Error()))
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
