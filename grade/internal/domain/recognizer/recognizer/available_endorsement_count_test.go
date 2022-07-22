package recognizer

import (
	"fmt"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvailableEndorsementCountConstructor(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedError error
	}{
		{uint8(0), nil},
		{yearlyEndorsementCount / 2, nil},
		{yearlyEndorsementCount, nil},
		{yearlyEndorsementCount + 1, ErrInvalidAvailableEndorsementCount},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, err := NewAvailableEndorsementCount(c.Arg)
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.Arg, uint8(g))
			}
		})
	}
}

func TestAvailableEndorsementCountHasAvailable(t *testing.T) {
	cases := []struct {
		Arg            uint8
		ExpectedResult bool
	}{
		{uint8(0), false},
		{yearlyEndorsementCount / 2, true},
		{yearlyEndorsementCount, true},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, _ := NewAvailableEndorsementCount(c.Arg)
			r := g.HasAvailable()
			assert.Equal(t, c.ExpectedResult, r)
		})
	}
}

func TestAvailableEndorsementCountNext(t *testing.T) {
	cases := []struct {
		Arg           uint8
		ExpectedValue uint8
		ExpectedError error
	}{
		{uint8(0), uint8(0), ErrInvalidAvailableEndorsementCount},
		{yearlyEndorsementCount / 2, yearlyEndorsementCount/2 - 1, nil},
		{yearlyEndorsementCount, yearlyEndorsementCount - 1, nil},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			g, _ := NewAvailableEndorsementCount(c.Arg)
			n, err := g.Decrease()
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, c.ExpectedValue, uint8(n))
			}
		})
	}
}

func TestAvailableEndorsementCountExportTo(t *testing.T) {
	var ex seedwork.Uint8Exporter
	c, _ := NewAvailableEndorsementCount(1)
	c.ExportTo(&ex)
	assert.Equal(t, uint8(ex), c.Export())
}
