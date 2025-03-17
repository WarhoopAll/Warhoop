package web

import (
	"context"
	"warhoop/app/model"
	"warhoop/app/utils"
)

func (svc *WebService) GetCommentByNewsID(ctx context.Context, id int) (*model.CommentSlice, error) {
	entry, err := svc.store.SaitRepo.GetCommentsByNewsID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	comment := entry.ToWeb()
	return &comment, nil
}
func (svc *WebService) CreateComment(ctx context.Context, id int, entry *model.Comment) (*model.Comment, error) {
	entry.ProfileID = id
	res, err := svc.store.SaitRepo.CreateComment(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) DeleteComment(ctx context.Context, id int) error {
	err := svc.store.SaitRepo.DeleteComment(ctx, id)
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}

func (svc *WebService) UpdateComment(ctx context.Context, id int, entry *model.Comment) (*model.Comment, error) {
	comment, err := svc.GetCommentByID(ctx, entry.ID)
	if err != nil {
		return nil, err
	}

	if comment.Profile == nil || comment.Profile.AccountID != id {
		return nil, err
	}

	res, err := svc.store.SaitRepo.UpdateComment(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}

func (svc *WebService) GetCommentByID(ctx context.Context, id int) (*model.Comment, error) {
	res, err := svc.store.SaitRepo.GetCommentByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return res.ToWeb(), nil
}
