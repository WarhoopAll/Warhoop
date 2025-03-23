package nexus

import (
	"github.com/uptrace/bun"
	"time"
)

type Report struct {
	ID              int       `json:"id,omitempty"`
	Reporter        int       `json:"-"`
	Victim          int       `json:"victim,omitempty"`
	ReporterProfile *Profile  `json:"reporter_profile,omitempty"`
	VictimProfile   *Profile  `json:"victim_profile,omitempty"`
	Reason          string    `json:"reason,omitempty"`
	Status          uint8     `json:"status,omitempty"`
	ObjectType      int       `json:"object_type,omitempty"`
	ObjectID        int       `json:"object_id,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

type ReportSlice []Report

type DBReport struct {
	bun.BaseModel   `bun:"table:reports,alias:report"`
	ID              int        `bun:"id,pk,autoincrement"`
	Reporter        int        `bun:"reporter_id"`
	Victim          int        `bun:"victim_id"`
	Reason          string     `bun:"reason"`
	Status          uint8      `bun:"status"`
	ObjectType      int        `bun:"object_type"` // 1 = News, 2 = Comment, 3 = Profile
	ObjectID        int        `bun:"object_id"`
	CreatedAt       time.Time  `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt       time.Time  `bun:"updated_at,nullzero,default:current_timestamp on update current_timestamp"`
	ReporterProfile *DBProfile `bun:"rel:belongs-to,join:reporter_id=account_id"`
	VictimProfile   *DBProfile `bun:"rel:belongs-to,join:victim_id=account_id"`
}

type DBReportSlice []DBReport

// ToDB converts Report to DBReport
func (entry *Report) ToDB() *DBReport {
	if entry == nil {
		return nil
	}
	return &DBReport{
		ID:         entry.ID,
		Reporter:   entry.Reporter,
		ObjectType: entry.ObjectType,
		ObjectID:   entry.ObjectID,
		Victim:     entry.Victim,
		Reason:     entry.Reason,
		Status:     entry.Status,
		CreatedAt:  entry.CreatedAt,
		UpdatedAt:  entry.UpdatedAt,
	}
}

// ToWeb converts DBReport to Report
func (entry *DBReport) ToWeb() *Report {
	if entry == nil {
		return nil
	}
	return &Report{
		ID:              entry.ID,
		ReporterProfile: entry.ReporterProfile.ToWeb(),
		VictimProfile:   entry.VictimProfile.ToWeb(),
		ObjectType:      entry.ObjectType,
		ObjectID:        entry.ObjectID,
		Reason:          entry.Reason,
		Status:          entry.Status,
		CreatedAt:       entry.CreatedAt,
		UpdatedAt:       entry.UpdatedAt,
	}
}

// ToDB converts ReportsSlice to DBReportsSlice
func (data ReportSlice) ToDB() DBReportSlice {
	result := make(DBReportSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

// ToWeb converts DBReportsSlice to ReportsSlice
func (data DBReportSlice) ToWeb() ReportSlice {
	result := make(ReportSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
