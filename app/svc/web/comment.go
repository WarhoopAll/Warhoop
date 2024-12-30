package web

import (
	"context"
	"warhoop/app/model"
)

func (svc WebService) GetCommentByNewsID(ctx context.Context, id int) (*model.CommentSlice, error) {
	entry, err := svc.store.SaitRepo.GetCommentByNewsID(ctx, id)
	if err != nil {
		return nil, err
	}
	comment := entry.ToWeb()
	return &comment, nil
}
func (svc WebService) CreateComment(ctx context.Context, entry *model.Comment) (*model.Comment, error) {
	res, err := svc.store.SaitRepo.CreateComment(ctx, entry.ToDB())
	if err != nil {
		return nil, err
	}
	return res.ToWeb(), nil
}

func (svc WebService) DeleteComment(ctx context.Context, id int) error {
	err := svc.store.SaitRepo.DeleteComment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc WebService) UpdateComment(ctx context.Context, entry *model.Comment) error {
	err := svc.store.SaitRepo.UpdateComment(ctx, entry.ToDB())
	if err != nil {
		return err
	}
	return nil
}

func (svc WebService) GetCommentByID(ctx context.Context, id int) (*model.Comment, error) {
	res, err := svc.store.SaitRepo.GetCommentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.ToWeb(), nil
}
