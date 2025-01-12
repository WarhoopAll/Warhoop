package web

import (
	"context"
	"warhoop/app/model"
	"warhoop/app/utils"
)

func (svc WebService) CreateReport(ctx context.Context, id int, entry *model.Report) (*model.Report, error) {
	entry.Reporter = id
	result, err := svc.store.SaitRepo.CreateReport(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return result.ToWeb(), nil
}

func (svc WebService) GetReportByID(ctx context.Context, id int) (*model.Report, error) {
	result, err := svc.store.SaitRepo.GetReportByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return result.ToWeb(), nil
}

func (svc WebService) GetReports(ctx context.Context, limit, offset int) (*model.ReportSlice, error) {
	result, err := svc.store.SaitRepo.GetReports(ctx, limit, offset)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	reports := result.ToWeb()
	return &reports, nil
}

func (svc WebService) DeleteReportByID(ctx context.Context, id int) error {
	err := svc.store.SaitRepo.DeleteReportByID(ctx, id)
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}
