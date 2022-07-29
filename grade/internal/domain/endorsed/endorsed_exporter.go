package endorsed

import (
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	interfaces4 "github.com/emacsway/qualifying-grade/grade/internal/domain/member/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type EndorsedExporter struct {
	Id                   interfaces4.TenantMemberIdExporter
	Grade                interfaces.Exporter[uint8]
	ReceivedEndorsements []interfaces2.EndorsementExporter
	GradeLogEntries      []interfaces2.GradeLogEntryExporter
	Version              uint
	CreatedAt            time.Time
}

func (ex *EndorsedExporter) SetState(
	id interfaces4.TenantMemberIdExporter,
	grade interfaces.Exporter[uint8],
	receivedEndorsements []interfaces2.EndorsementExporter,
	gradeLogEntries []interfaces2.GradeLogEntryExporter,
	version uint,
	createdAt time.Time,
) {
	ex.Id = id
	ex.Grade = grade
	ex.ReceivedEndorsements = receivedEndorsements
	ex.GradeLogEntries = gradeLogEntries
	ex.Version = version
	ex.CreatedAt = createdAt
}

type EndorsedState struct {
	Id                   member.TenantMemberIdState
	Grade                uint8
	ReceivedEndorsements []EndorsementState
	GradeLogEntries      []GradeLogEntryState
	Version              uint
	CreatedAt            time.Time
}
