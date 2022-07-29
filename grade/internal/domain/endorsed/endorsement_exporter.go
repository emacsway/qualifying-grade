package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/member/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type EndorsementExporter struct {
	RecognizerId      interfaces2.TenantMemberIdExporter
	RecognizerGrade   interfaces.Exporter[uint8]
	RecognizerVersion uint
	EndorsedId        interfaces2.TenantMemberIdExporter
	EndorsedGrade     interfaces.Exporter[uint8]
	EndorsedVersion   uint
	ArtifactId        interfaces.Exporter[uint64]
	CreatedAt         time.Time
}

func (ex *EndorsementExporter) SetState(
	recognizerId interfaces2.TenantMemberIdExporter,
	recognizerGrade interfaces.Exporter[uint8],
	recognizerVersion uint,
	endorsedId interfaces2.TenantMemberIdExporter,
	endorsedGrade interfaces.Exporter[uint8],
	endorsedVersion uint,
	artifactId interfaces.Exporter[uint64],
	createdAt time.Time,
) {
	ex.RecognizerId = recognizerId
	ex.RecognizerGrade = recognizerGrade
	ex.RecognizerVersion = recognizerVersion
	ex.EndorsedId = endorsedId
	ex.EndorsedGrade = endorsedGrade
	ex.EndorsedVersion = endorsedVersion
	ex.ArtifactId = artifactId
	ex.CreatedAt = createdAt
}

type EndorsementState struct {
	RecognizerId      member.TenantMemberIdState
	RecognizerGrade   uint8
	RecognizerVersion uint
	EndorsedId        member.TenantMemberIdState
	EndorsedGrade     uint8
	EndorsedVersion   uint
	ArtifactId        uint64
	CreatedAt         time.Time
}
