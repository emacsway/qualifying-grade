package recognizer

import (
	"errors"
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

var (
	ErrNoEndorsementAvailable         = errors.New("no endorsement is available")
	ErrNoEndorsementReservation       = errors.New("no endorsement is reserved")
	ErrEndorsementReservationExceeded = errors.New("endorsement reservation exceeded")
)

func NewRecognizer(
	id member.TenantMemberId,
	createdAt time.Time,
) (*Recognizer, error) {
	availableCount, err := NewEndorsementCount(YearlyEndorsementCount)
	if err != nil {
		return nil, err
	}
	pendingCount, err := NewEndorsementCount(0)
	if err != nil {
		return nil, err
	}
	versioned, err := seedwork.NewVersionedAggregate(0)
	if err != nil {
		return nil, err
	}
	eventive, err := seedwork.NewEventiveEntity()
	if err != nil {
		return nil, err
	}
	return &Recognizer{
		id:                        id,
		grade:                     shared.WithoutGrade,
		availableEndorsementCount: availableCount,
		pendingEndorsementCount:   pendingCount,
		createdAt:                 createdAt,
		VersionedAggregate:        versioned,
		EventiveEntity:            eventive,
	}, nil
}

type Recognizer struct {
	id                        member.TenantMemberId
	grade                     shared.Grade
	availableEndorsementCount EndorsementCount
	pendingEndorsementCount   EndorsementCount
	createdAt                 time.Time
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}

func (r Recognizer) GetId() member.TenantMemberId {
	return r.id
}

func (r Recognizer) GetGrade() shared.Grade {
	return r.grade
}

func (r *Recognizer) SetGrade(g shared.Grade) error {
	r.grade = g
	return nil
}

func (r Recognizer) CanReserveEndorsement() error {
	if !(r.availableEndorsementCount > r.pendingEndorsementCount) {
		return ErrEndorsementReservationExceeded
	}
	return nil
}

func (r Recognizer) CanCompleteEndorsement() error {
	if r.pendingEndorsementCount == 0 {
		return ErrNoEndorsementReservation
	}
	if r.availableEndorsementCount == 0 {
		return ErrNoEndorsementAvailable
	}
	if r.availableEndorsementCount < r.pendingEndorsementCount {
		return ErrEndorsementReservationExceeded
	}
	return nil
}

func (r *Recognizer) ReserveEndorsement() error {
	err := r.CanReserveEndorsement()
	if err != nil {
		return err
	}
	r.pendingEndorsementCount += 1
	return nil
}

func (r *Recognizer) ReleaseEndorsementReservation() error {
	if r.pendingEndorsementCount == 0 {
		return ErrNoEndorsementReservation
	}
	r.pendingEndorsementCount -= 1
	return nil
}

func (r *Recognizer) CompleteEndorsement() error {
	err := r.CanCompleteEndorsement()
	if err != nil {
		return err
	}
	r.availableEndorsementCount -= 1
	r.pendingEndorsementCount -= 1
	return nil
}

func (r Recognizer) ExportTo(ex RecognizerExporterSetter) {
	var id member.TenantMemberIdExporter
	var grade, availableEndorsementCount, pendingEndorsementCount seedwork.Uint8Exporter

	r.id.ExportTo(&id)
	r.grade.ExportTo(&grade)
	r.availableEndorsementCount.ExportTo(&availableEndorsementCount)
	r.pendingEndorsementCount.ExportTo(&pendingEndorsementCount)
	ex.SetState(
		&id, &grade, &availableEndorsementCount, &pendingEndorsementCount, r.GetVersion(), r.createdAt,
	)
}

func (r Recognizer) Export() RecognizerState {
	return RecognizerState{
		Id:                        r.id.Export(),
		Grade:                     r.grade.Export(),
		AvailableEndorsementCount: r.availableEndorsementCount.Export(),
		PendingEndorsementCount:   r.pendingEndorsementCount.Export(),
		Version:                   r.GetVersion(),
		CreatedAt:                 r.createdAt,
	}
}

type RecognizerExporterSetter interface {
	SetState(
		id member.TenantMemberIdExporterSetter,
		grade seedwork.ExporterSetter[uint8],
		availableEndorsementCount seedwork.ExporterSetter[uint8],
		pendingEndorsementCount seedwork.ExporterSetter[uint8],
		version uint,
		createdAt time.Time,
	)
}
