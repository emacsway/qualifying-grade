package endorsed

import (
	"errors"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

var (
	ErrAlreadyEndorsed = errors.New(
		"this artifact has already been endorsed by the recogniser",
	)
)

func NewEndorsed(
	id member.TenantMemberId,
	createdAt time.Time,
) (*Endorsed, error) {
	versioned, err := seedwork.NewVersionedAggregate(0)
	if err != nil {
		return nil, err
	}
	eventive, err := seedwork.NewEventiveEntity()
	if err != nil {
		return nil, err
	}
	return &Endorsed{
		id:                 id,
		grade:              shared.WithoutGrade,
		VersionedAggregate: versioned,
		EventiveEntity:     eventive,
		createdAt:          createdAt,
	}, nil
}

type Endorsed struct {
	id                   member.TenantMemberId
	grade                shared.Grade
	receivedEndorsements []Endorsement
	gradeLogEntries      []GradeLogEntry
	createdAt            time.Time
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}

func (e *Endorsed) ReceiveEndorsement(r recognizer.Recognizer, aId artifact.ArtifactId, t time.Time) error {
	err := e.canReceiveEndorsement(r, aId)
	if err != nil {
		return err
	}
	ent, err := NewEndorsement(
		r.GetId(), r.GetGrade(), r.GetVersion(),
		e.id, e.grade, e.GetVersion(),
		aId, t,
	)
	if err != nil {
		return err
	}
	e.receivedEndorsements = append(e.receivedEndorsements, ent)
	err = e.actualizeGrade(t)
	if err != nil {
		return err
	}
	return nil
}

func (e Endorsed) canReceiveEndorsement(r recognizer.Recognizer, aId artifact.ArtifactId) error {
	err := r.CanCompleteEndorsement()
	if err != nil {
		return err
	}
	return e.canBeEndorsed(r, aId)
}

func (e Endorsed) canBeEndorsed(r recognizer.Recognizer, aId artifact.ArtifactId) error {
	for _, ent := range e.receivedEndorsements {
		if ent.IsEndorsedBy(r.GetId(), aId) {
			return ErrAlreadyEndorsed
		}
	}
	return CanEndorse(r.GetId(), r.GetGrade(), e.id, e.grade)
}

func (e Endorsed) CanBeginEndorsement(r recognizer.Recognizer, aId artifact.ArtifactId) error {
	err := r.CanReserveEndorsement()
	if err != nil {
		return err
	}
	return e.canBeEndorsed(r, aId)
}

func (e *Endorsed) actualizeGrade(t time.Time) error {
	if e.grade.NextGradeAchieved(e.getReceivedEndorsementCount()) {
		nextGrade, err := e.grade.Next()
		if err != nil {
			return err
		}
		reason, err := NewReason("Endorsement count is achieved")
		if err != nil {
			return err
		}
		return e.setGrade(nextGrade, reason, t)
	}
	return nil
}
func (e Endorsed) getReceivedEndorsementCount() uint {
	var counter uint
	for _, v := range e.receivedEndorsements {
		if v.GetEndorsedGrade() == e.grade {
			counter += uint(v.GetWeight())
		}
	}
	return counter
}

func (e *Endorsed) setGrade(g shared.Grade, reason Reason, t time.Time) error {
	gle, err := NewGradeLogEntry(
		e.id, e.GetVersion(), g, reason, t,
	)
	if err != nil {
		return err
	}
	e.gradeLogEntries = append(e.gradeLogEntries, gle)
	e.grade = g
	return nil
}

func (e *Endorsed) DecreaseGrade(reason Reason, t time.Time) error {
	previousGrade, err := e.grade.Next()
	if err != nil {
		return err
	}
	return e.setGrade(previousGrade, reason, t)
}

func (e Endorsed) ExportTo(ex interfaces.EndorsedExporter) {
	var id member.TenantMemberIdExporter
	var grade seedwork.Uint8Exporter
	var receivedEndorsements []interfaces.EndorsementExporter
	var gradeLogEntries []interfaces.GradeLogEntryExporter

	for _, v := range e.receivedEndorsements {
		re := &EndorsementExporter{}
		v.ExportTo(re)
		receivedEndorsements = append(receivedEndorsements, re)
	}

	for _, v := range e.gradeLogEntries {
		gle := &GradeLogEntryExporter{}
		v.ExportTo(gle)
		gradeLogEntries = append(gradeLogEntries, gle)
	}

	e.id.ExportTo(&id)
	e.grade.ExportTo(&grade)
	ex.SetState(
		&id, &grade, receivedEndorsements, gradeLogEntries, e.GetVersion(), e.createdAt,
	)
}

func (e Endorsed) Export() EndorsedState {
	var receivedEndorsements []EndorsementState
	var gradeLogEntries []GradeLogEntryState
	for _, v := range e.receivedEndorsements {
		receivedEndorsements = append(receivedEndorsements, v.Export())
	}
	for _, v := range e.gradeLogEntries {
		gradeLogEntries = append(gradeLogEntries, v.Export())
	}
	return EndorsedState{
		Id:                   e.id.Export(),
		Grade:                e.grade.Export(),
		ReceivedEndorsements: receivedEndorsements,
		GradeLogEntries:      gradeLogEntries,
		Version:              e.GetVersion(),
		CreatedAt:            e.createdAt,
	}
}
