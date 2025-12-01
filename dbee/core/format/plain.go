package format

import (
	"fmt"

	"github.com/kndndrj/nvim-dbee/dbee/core"
)

type Plain struct {
	col int
}

func NewPlain(col int) *Plain {
	return &Plain{col}
}

func (pf *Plain) Format(header core.Header, rows []core.Row, _ *core.FormatterOptions) ([]byte, error) {
	cell := rows[0][pf.col]

	return fmt.Appendf(nil, "%v", cell), nil
}
