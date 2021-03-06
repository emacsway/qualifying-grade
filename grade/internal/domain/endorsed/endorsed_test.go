package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndorsedExport(t *testing.T) {
	ef := NewEndorsedFakeFactory()
	rf := recognizer.NewRecognizerFakeFactory()
	for i := 0; i < 4; i++ {
		err := ef.ReceiveEndorsement(rf)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
	agg, err := ef.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, EndorsedState{
		Id: member.TenantMemberIdState{
			TenantId: ef.Id.TenantId,
			MemberId: ef.Id.MemberId,
		},
		Grade: ef.Grade + 1,
		ReceivedEndorsements: []endorsement.EndorsementState{
			{
				RecognizerId: member.TenantMemberIdState{
					TenantId: rf.Id.TenantId,
					MemberId: rf.Id.MemberId,
				},
				RecognizerGrade:   rf.Grade,
				RecognizerVersion: 0,
				EndorsedId: member.TenantMemberIdState{
					TenantId: ef.Id.TenantId,
					MemberId: ef.Id.MemberId,
				},
				EndorsedGrade:   ef.Grade,
				EndorsedVersion: 0,
				ArtifactId:      ef.ReceivedEndorsements[0].ArtifactId,
				CreatedAt:       ef.ReceivedEndorsements[0].CreatedAt,
			},
			{
				RecognizerId: member.TenantMemberIdState{
					TenantId: rf.Id.TenantId,
					MemberId: rf.Id.MemberId,
				},
				RecognizerGrade:   rf.Grade,
				RecognizerVersion: 0,
				EndorsedId: member.TenantMemberIdState{
					TenantId: ef.Id.TenantId,
					MemberId: ef.Id.MemberId,
				},
				EndorsedGrade:   ef.Grade,
				EndorsedVersion: 1,
				ArtifactId:      ef.ReceivedEndorsements[1].ArtifactId,
				CreatedAt:       ef.ReceivedEndorsements[1].CreatedAt,
			},
			{
				RecognizerId: member.TenantMemberIdState{
					TenantId: rf.Id.TenantId,
					MemberId: rf.Id.MemberId,
				},
				RecognizerGrade:   rf.Grade,
				RecognizerVersion: 0,
				EndorsedId: member.TenantMemberIdState{
					TenantId: ef.Id.TenantId,
					MemberId: ef.Id.MemberId,
				},
				EndorsedGrade:   ef.Grade,
				EndorsedVersion: 2,
				ArtifactId:      ef.ReceivedEndorsements[2].ArtifactId,
				CreatedAt:       ef.ReceivedEndorsements[2].CreatedAt,
			},
			{
				RecognizerId: member.TenantMemberIdState{
					TenantId: rf.Id.TenantId,
					MemberId: rf.Id.MemberId,
				},
				RecognizerGrade:   rf.Grade,
				RecognizerVersion: 0,
				EndorsedId: member.TenantMemberIdState{
					TenantId: ef.Id.TenantId,
					MemberId: ef.Id.MemberId,
				},
				EndorsedGrade:   ef.Grade + 1,
				EndorsedVersion: 3,
				ArtifactId:      ef.ReceivedEndorsements[3].ArtifactId,
				CreatedAt:       ef.ReceivedEndorsements[3].CreatedAt,
			},
		},
		GradeLogEntries: []gradelogentry.GradeLogEntryState{
			{
				EndorsedId: member.TenantMemberIdState{
					TenantId: ef.Id.TenantId,
					MemberId: ef.Id.MemberId,
				},
				EndorsedVersion: 2,
				AssignedGrade:   ef.Grade + 1,
				Reason:          "Endorsement count is achieved",
				CreatedAt:       ef.ReceivedEndorsements[2].CreatedAt,
			},
		},
		Version:   4,
		CreatedAt: ef.CreatedAt,
	}, agg.Export())
}

func TestEndorsedExportTo(t *testing.T) {
	var actualExporter EndorsedExporter
	ef := NewEndorsedFakeFactory()
	rf := recognizer.NewRecognizerFakeFactory()
	for i := 0; i < 4; i++ {
		err := ef.ReceiveEndorsement(rf)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
	agg, err := ef.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.ExportTo(&actualExporter)
	assert.Equal(t, EndorsedExporter{
		Id:    member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
		Grade: seedwork.NewUint8Exporter(ef.Grade + 1),
		ReceivedEndorsements: []endorsement.EndorsementExporterSetter{
			&endorsement.EndorsementExporter{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.NewUint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedGrade:     seedwork.NewUint8Exporter(ef.Grade),
				EndorsedVersion:   0,
				ArtifactId:        seedwork.NewUint64Exporter(ef.ReceivedEndorsements[0].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[0].CreatedAt,
			},
			&endorsement.EndorsementExporter{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.NewUint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedGrade:     seedwork.NewUint8Exporter(ef.Grade),
				EndorsedVersion:   1,
				ArtifactId:        seedwork.NewUint64Exporter(ef.ReceivedEndorsements[1].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[1].CreatedAt,
			},
			&endorsement.EndorsementExporter{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.NewUint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedGrade:     seedwork.NewUint8Exporter(ef.Grade),
				EndorsedVersion:   2,
				ArtifactId:        seedwork.NewUint64Exporter(ef.ReceivedEndorsements[2].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[2].CreatedAt,
			},
			&endorsement.EndorsementExporter{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.NewUint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedGrade:     seedwork.NewUint8Exporter(ef.Grade + 1),
				EndorsedVersion:   3,
				ArtifactId:        seedwork.NewUint64Exporter(ef.ReceivedEndorsements[3].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[3].CreatedAt,
			},
		},
		GradeLogEntries: []gradelogentry.GradeLogEntryExporterSetter{
			&gradelogentry.GradeLogEntryExporter{
				EndorsedId:      member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedVersion: 2,
				AssignedGrade:   seedwork.NewUint8Exporter(ef.Grade + 1),
				Reason:          seedwork.NewStringExporter("Endorsement count is achieved"),
				CreatedAt:       ef.ReceivedEndorsements[2].CreatedAt,
			},
		},
		Version:   4,
		CreatedAt: ef.CreatedAt,
	}, actualExporter)
}
