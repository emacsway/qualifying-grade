package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/member/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type GradeLogEntryExporter struct {
	EndorsedId      interfaces2.TenantMemberIdExporter
	EndorsedVersion uint
	AssignedGrade   interfaces.Exporter[uint8]
	Reason          interfaces.Exporter[string]
	CreatedAt       time.Time
}

func (ex *GradeLogEntryExporter) SetState(
	endorsedId interfaces2.TenantMemberIdExporter,
	endorsedVersion uint,
	assignedGrade interfaces.Exporter[uint8],
	reason interfaces.Exporter[string],
	createdAt time.Time,
) {
	ex.EndorsedId = endorsedId
	ex.EndorsedVersion = endorsedVersion
	ex.AssignedGrade = assignedGrade
	ex.Reason = reason
	ex.CreatedAt = createdAt
}

type GradeLogEntryState struct {
	EndorsedId      member.TenantMemberIdState
	EndorsedVersion uint
	AssignedGrade   uint8
	Reason          string
	CreatedAt       time.Time
}
