package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewEndorsementFakeFactory() (*EndorsementFakeFactory, error) {
	recognizerIdFactory, err := member.NewTenantMemberIdFakeFactory()
	if err != nil {
		return nil, err
	}
	recognizerIdFactory.MemberId = 1
	endorsedIdFactory, err := member.NewTenantMemberIdFakeFactory()
	if err != nil {
		return nil, err
	}
	endorsedIdFactory.MemberId = 2
	return &EndorsementFakeFactory{
		RecognizerId:      recognizerIdFactory,
		RecognizerGrade:   2,
		RecognizerVersion: 3,
		EndorsedId:        endorsedIdFactory,
		EndorsedGrade:     1,
		EndorsedVersion:   5,
		ArtifactId:        6,
		CreatedAt:         time.Now(),
	}, nil
}

type EndorsementFakeFactory struct {
	RecognizerId      *member.TenantMemberIdFakeFactory
	RecognizerGrade   uint8
	RecognizerVersion uint
	EndorsedId        *member.TenantMemberIdFakeFactory
	EndorsedGrade     uint8
	EndorsedVersion   uint
	ArtifactId        uint64
	CreatedAt         time.Time
}

func (f EndorsementFakeFactory) Create() (Endorsement, error) {
	recognizerId, err := member.NewTenantMemberId(f.RecognizerId.TenantId, f.RecognizerId.MemberId)
	if err != nil {
		return Endorsement{}, err
	}
	recognizerGrade, err := shared.NewGrade(f.RecognizerGrade)
	if err != nil {
		return Endorsement{}, err
	}
	endorsedId, err := member.NewTenantMemberId(f.EndorsedId.TenantId, f.EndorsedId.MemberId)
	if err != nil {
		return Endorsement{}, err
	}
	endorsedGrade, err := shared.NewGrade(f.EndorsedGrade)
	if err != nil {
		return Endorsement{}, err
	}
	artifactId, err := artifact.NewArtifactId(f.ArtifactId)
	if err != nil {
		return Endorsement{}, err
	}
	return NewEndorsement(
		recognizerId, recognizerGrade, f.RecognizerVersion,
		endorsedId, endorsedGrade, f.EndorsedVersion,
		artifactId, f.CreatedAt,
	)
}
