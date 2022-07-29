package endorsed

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewEndorsedFakeFactory() (*EndorsedFakeFactory, error) {
	idFactory, err := member.NewTenantMemberIdFakeFactory()
	if err != nil {
		return nil, err
	}
	idFactory.MemberId = 2
	return &EndorsedFakeFactory{
		Id:                idFactory,
		Grade:             0,
		CreatedAt:         time.Now(),
		CurrentArtifactId: 1000,
	}, nil
}

type EndorsedFakeFactory struct {
	Id                   *member.TenantMemberIdFakeFactory
	Grade                uint8
	ReceivedEndorsements []*EndorsementFakeFactory2
	CreatedAt            time.Time
	CurrentArtifactId    uint64
}

func (f *EndorsedFakeFactory) achieveGrade() error {
	currentGrade := shared.WithoutGrade
	targetGrade, err := shared.NewGrade(f.Grade)
	if err != nil {
		return err
	}
	for currentGrade < targetGrade {
		r, err := recognizer.NewRecognizerFakeFactory()
		if err != nil {
			return err
		}
		rId, err := member.NewTenantMemberIdFakeFactory()
		if err != nil {
			return err
		}
		rId.MemberId = 1000
		r.Id = rId
		recognizerGrade, _ := currentGrade.Next()
		r.Grade = recognizerGrade.Export()
		var endorsementCount uint = 0
		for !currentGrade.NextGradeAchieved(endorsementCount) {
			err = f.receiveEndorsement(r)
			if err != nil {
				return err
			}
			endorsementCount += 2
		}
		currentGrade, err = currentGrade.Next()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *EndorsedFakeFactory) ReceiveEndorsement(r *recognizer.RecognizerFakeFactory) error {
	err := f.achieveGrade()
	if err != nil {
		return err
	}
	err = f.receiveEndorsement(r)
	if err != nil {
		return err
	}
	return nil
}

func (f *EndorsedFakeFactory) receiveEndorsement(r *recognizer.RecognizerFakeFactory) error {
	e, err := NewEndorsementFakeFactory2(r)
	if err != nil {
		return err
	}
	e.ArtifactId = f.CurrentArtifactId
	f.CurrentArtifactId += 1
	e.CreatedAt = time.Now()
	f.ReceivedEndorsements = append(f.ReceivedEndorsements, e)
	return nil
}

func (f EndorsedFakeFactory) Create() (*Endorsed, error) {
	err := f.achieveGrade()
	if err != nil {
		return nil, err
	}
	id, err := member.NewTenantMemberId(f.Id.TenantId, f.Id.MemberId)
	if err != nil {
		return nil, err
	}
	e, err := NewEndorsed(id, f.CreatedAt)
	if err != nil {
		return nil, err
	}
	for _, entf := range f.ReceivedEndorsements {
		r, err := entf.Recognizer.Create()
		if err != nil {
			return nil, err
		}
		artifactId, err := artifact.NewArtifactId(entf.ArtifactId)
		if err != nil {
			return nil, err
		}
		err = r.ReserveEndorsement()
		if err != nil {
			return nil, err
		}
		err = e.ReceiveEndorsement(*r, artifactId, entf.CreatedAt)
		if err != nil {
			return nil, err
		}
		e.IncreaseVersion()
	}
	return e, nil
}

func NewEndorsementFakeFactory2(r *recognizer.RecognizerFakeFactory) (*EndorsementFakeFactory2, error) {
	return &EndorsementFakeFactory2{
		Recognizer: r,
		ArtifactId: 6,
		CreatedAt:  time.Now(),
	}, nil
}

type EndorsementFakeFactory2 struct {
	Recognizer *recognizer.RecognizerFakeFactory
	ArtifactId uint64
	CreatedAt  time.Time
}
