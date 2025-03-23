package web

import (
	"context"
	"warhoop/app/model/nexus"
	"warhoop/app/utils"
)

func (svc *WebService) CreateReport(ctx context.Context, id int, entry *nexus.Report) (*nexus.Report, error) {
	entry.Reporter = id
	result, err := svc.store.NexusRepo.CreateReport(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return result.ToWeb(), nil
}

func (svc *WebService) GetReportByID(ctx context.Context, id int) (*nexus.Report, error) {
	result, err := svc.store.NexusRepo.GetReportByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	return result.ToWeb(), nil
}

func (svc *WebService) GetReports(ctx context.Context, limit, offset int) (*nexus.ReportSlice, error) {
	result, err := svc.store.NexusRepo.GetReports(ctx, limit, offset)
	if err != nil {
		return nil, utils.ErrDataBase
	}
	reports := result.ToWeb()
	return &reports, nil
}

func (svc *WebService) DeleteReportByID(ctx context.Context, id int) error {
	err := svc.store.NexusRepo.DeleteReportByID(ctx, id)
	if err != nil {
		return utils.ErrDataBase
	}
	return nil
}
