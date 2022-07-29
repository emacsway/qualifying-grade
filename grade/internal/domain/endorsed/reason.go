package endorsed

import "github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"

func NewReason(reason string) (Reason, error) {
	return Reason(reason), nil
}

type Reason string

func (r Reason) Export() string {
	return string(r)
}

func (r Reason) ExportTo(ex interfaces.Exporter[string]) {
	ex.SetState(string(r))
}
