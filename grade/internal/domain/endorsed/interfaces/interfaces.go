package interfaces

import (
	interfaces4 "github.com/emacsway/qualifying-grade/grade/internal/domain/member/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type EndorsedExporter interface {
	SetState(
		id interfaces4.TenantMemberIdExporter,
		grade interfaces.Exporter[uint8],
		receivedEndorsements []EndorsementExporter,
		gradeLogEntries []GradeLogEntryExporter,
		version uint,
		createdAt time.Time,
	)
}

type GradeLogEntryExporter interface {
	SetState(
		endorsedId interfaces4.TenantMemberIdExporter,
		endorsedVersion uint,
		assignedGrade interfaces.Exporter[uint8],
		reason interfaces.Exporter[string],
		createdAt time.Time,
	)
}

type EndorsementExporter interface {
	SetState(
		recognizerId interfaces4.TenantMemberIdExporter,
		recognizerGrade interfaces.Exporter[uint8],
		recognizerVersion uint,
		endorsedId interfaces4.TenantMemberIdExporter,
		endorsedGrade interfaces.Exporter[uint8],
		endorsedVersion uint,
		artifactId interfaces.Exporter[uint64],
		createdAt time.Time,
	)
}
