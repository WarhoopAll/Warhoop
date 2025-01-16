package bun

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/model"
)

func (r *SaitRepo) CreateReport(ctx context.Context, entry *model.DBReport) (*model.DBReport, error) {
	_, err := r.db.NewInsert().
		Model(entry).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.CreateReport",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}
	return entry, nil
}

func (r *SaitRepo) GetReportByID(ctx context.Context, id int) (*model.DBReport, error) {
	entry := &model.DBReport{}
	err := r.db.NewSelect().
		Model(entry).
		Relation("ReporterProfile").
		Relation("VictimProfile").
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetReportByID",
			log.String("err", err.Error()),
			log.Int("id", id),
		)
		return nil, err
	}
	return entry, nil
}

func (r *SaitRepo) UpdateReport(ctx context.Context, entry *model.DBReport) (*model.DBReport, error) {
	_, err := r.db.NewUpdate().
		Model(entry).
		Where("id = ?", entry.ID).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.UpdateReport",
			log.String("err", err.Error()),
			log.Object("entry", entry),
		)
		return nil, err
	}
	return entry, nil
}

func (r *SaitRepo) GetReports(ctx context.Context, limit, offset int) (*model.DBReportSlice, error) {
	entry := &model.DBReportSlice{}
	err := r.db.NewSelect().
		Model(entry).
		Relation("ReporterProfile").
		Relation("VictimProfile").
		Limit(limit).
		Offset(offset).
		Scan(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.GetReports",
			log.String("err", err.Error()),
		)
		return nil, err
	}
	return entry, nil
}

func (r *SaitRepo) DeleteReportByID(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().
		Model((*model.DBReport)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		r.logger.Error("store.SaitRepo.DeleteReportByID",
			log.String("err", err.Error()),
			log.Int("id", id),
		)
		return err
	}
	return nil
}
