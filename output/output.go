package output

import (
	"io"

	"github.com/alxmsl/qrencode-go/qrencode"
)

type Interface interface {
	Output(*qrencode.Grid, io.Writer)
}
