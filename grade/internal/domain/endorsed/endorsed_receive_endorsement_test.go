package endorsed

import (
	"fmt"
	recognizer2 "github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
)

func TestEndorsedReceiveEndorsement(t *testing.T) {
	cases := []struct {
		RecogniserId    uint64
		RecognizerGrade uint8
		EndorsedId      uint64
		EndorsedGrade   uint8
		ExpectedError   error
	}{
		{1, 0, 2, 0, nil},
		{1, 1, 2, 0, nil},
		{1, 0, 2, 1, ErrLowerGradeEndorses},
		{3, 0, 3, 0, ErrEndorsementOneself},
	}
	ef, err := NewEndorsedFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	rf, err := recognizer.NewRecognizerFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			ef.Id.MemberId = c.EndorsedId
			ef.Grade = c.EndorsedGrade
			rf.Id.MemberId = c.RecogniserId
			rf.Grade = c.RecognizerGrade
			e, err := ef.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			r, err := rf.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			artifactId, err := artifact.NewArtifactId(ef.CurrentArtifactId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = r.ReserveEndorsement()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = e.ReceiveEndorsement(*r, artifactId, time.Now())
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}

func TestEndorsedCanCompleteEndorsement(t *testing.T) {
	cases := []struct {
		Prepare       func(*recognizer.Recognizer) error
		ExpectedError error
	}{
		{func(r *recognizer.Recognizer) error {
			return r.ReserveEndorsement()
		}, nil},
		{func(r *recognizer.Recognizer) error {
			return nil
		}, recognizer.ErrNoEndorsementReservation},
		{func(r *recognizer.Recognizer) error {
			for i := uint8(0); i < recognizer2.YearlyEndorsementCount; i++ {
				err := r.ReserveEndorsement()
				if err != nil {
					return err
				}
				err = r.CompleteEndorsement()
				if err != nil {
					return err
				}
			}
			return nil
		}, recognizer.ErrNoEndorsementReservation},
		{func(r *recognizer.Recognizer) error {
			err := r.ReserveEndorsement()
			if err != nil {
				return err
			}
			err = r.ReleaseEndorsementReservation()
			if err != nil {
				return err
			}
			return nil
		}, recognizer.ErrNoEndorsementReservation},
	}
	ef, err := NewEndorsedFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	rf, err := recognizer.NewRecognizerFakeFactory()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			e, err := ef.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			r, err := rf.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			artifactId, err := artifact.NewArtifactId(ef.CurrentArtifactId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = c.Prepare(r)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = e.ReceiveEndorsement(*r, artifactId, time.Now())
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}
