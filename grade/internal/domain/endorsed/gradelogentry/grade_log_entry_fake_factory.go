package gradelogentry

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewGradeLogEntryFakeFactory() (*GradeLogEntryFakeFactory, error) {
	return &GradeLogEntryFakeFactory{
		EndorsedId:      1,
		EndorsedVersion: 2,
		AssignedGrade:   1,
		CreatedAt:       time.Now(),
	}, nil
}

type GradeLogEntryFakeFactory struct {
	EndorsedId      uint64
	EndorsedVersion uint
	AssignedGrade   uint8
	CreatedAt       time.Time
}

func (f GradeLogEntryFakeFactory) Create() (GradeLogEntry, error) {
	endorsedId, _ := external.NewMemberId(f.EndorsedId)
	assignedGrade, _ := shared.NewGrade(f.AssignedGrade)
	return NewGradeLogEntry(endorsedId, f.EndorsedVersion, assignedGrade, f.CreatedAt)
}
