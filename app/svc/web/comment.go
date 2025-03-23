package web

import (
	"context"
	"warhoop/app/model/nexus"
	"warhoop/app/utils"
)

func (svc *WebService) GetCommentByNewsID(ctx context.Context, id int) (*nexus.CommentSlice, error) {
	entry, err := svc.store.NexusRepo.GetCommentsByNewsID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	comment := entry.ToWeb()
	return &comment, nil
}
func (svc *WebService) CreateComment(ctx context.Context, id int, entry *nexus.Comment) (*nexus.Comment, error) {
	entry.ProfileID = id
	res, err := svc.store.NexusRepo.CreateComment(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) DeleteComment(ctx context.Context, id int) error {
	err := svc.store.NexusRepo.DeleteComment(ctx, id)
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}

func (svc *WebService) UpdateComment(ctx context.Context, id int, entry *nexus.Comment) (*nexus.Comment, error) {
	comment, err := svc.GetCommentByID(ctx, entry.ID)
	if err != nil {
		return nil, err
	}

	if comment.Profile == nil || comment.Profile.AccountID != id {
		return nil, err
	}

	res, err := svc.store.NexusRepo.UpdateComment(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) GetCommentByID(ctx context.Context, id int) (*nexus.Comment, error) {
	res, err := svc.store.NexusRepo.GetCommentByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}
