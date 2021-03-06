package member

import (
	"fmt"
	"testing"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/stretchr/testify/assert"
)

func TestTenantMemberIdEquals(t *testing.T) {
	cases := []struct {
		TenantId       uint64
		MemberId       uint64
		OtherTenantId  uint64
		OtherMemberId  uint64
		ExpectedResult bool
	}{
		{1, 2, 1, 2, true},
		{1, 1, 1, 2, false},
		{2, 2, 1, 2, false},
		{1, 2, 1, 1, false},
		{1, 2, 2, 2, false},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			id, err := NewTenantMemberId(c.TenantId, c.MemberId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			otherId, err := NewTenantMemberId(c.OtherTenantId, c.OtherMemberId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			r := id.Equals(otherId)
			assert.Equal(t, c.ExpectedResult, r)
		})
	}
}

func TestTenantMemberIdExport(t *testing.T) {
	cid, err := NewTenantMemberId(1, 2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, TenantMemberIdState{
		TenantId: 1,
		MemberId: 2,
	}, cid.Export())
}

func TestRecognizerExportTo(t *testing.T) {
	var actualExporter TenantMemberIdExporter
	cid, err := NewTenantMemberId(1, 2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cid.ExportTo(&actualExporter)
	assert.Equal(t, TenantMemberIdExporter{
		TenantId: seedwork.NewUint64Exporter(1),
		MemberId: seedwork.NewUint64Exporter(2),
	}, actualExporter)
}
