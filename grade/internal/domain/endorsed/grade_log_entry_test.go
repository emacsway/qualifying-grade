package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"testing"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/stretchr/testify/assert"
)

func TestGradeLogEntryExport(t *testing.T) {
	f, err := NewGradeLogEntryFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	e, _ := f.Create()
	assert.Equal(t, GradeLogEntryState{
		EndorsedId: member.TenantMemberIdState{
			TenantId: f.EndorsedId.TenantId,
			MemberId: f.EndorsedId.MemberId,
		},
		EndorsedVersion: f.EndorsedVersion,
		AssignedGrade:   f.AssignedGrade,
		Reason:          f.Reason,
		CreatedAt:       f.CreatedAt,
	}, e.Export())
}

func TestRecognizerExportTo(t *testing.T) {
	var actualExporter GradeLogEntryExporter
	f, err := NewGradeLogEntryFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg, _ := f.Create()
	agg.ExportTo(&actualExporter)
	assert.Equal(t, GradeLogEntryExporter{
		EndorsedId:      member.NewTenantMemberIdExporter(f.EndorsedId.TenantId, f.EndorsedId.MemberId),
		EndorsedVersion: f.EndorsedVersion,
		AssignedGrade:   seedwork.NewUint8Exporter(f.AssignedGrade),
		Reason:          seedwork.NewStringExporter(f.Reason),
		CreatedAt:       f.CreatedAt,
	}, actualExporter)
}
