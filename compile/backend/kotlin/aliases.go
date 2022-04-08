package kotlin

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeAliases(w io.Writer, i *ir.IR) {
	for _, a := range i.Aliases {
		fmt.Fprintf(w, "typealias %s = %s\n", NameConv(a.Name), TypeConv(a.OriginalType))
	}
}
