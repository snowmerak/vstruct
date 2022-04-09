package kotlin

import (
	"fmt"
	"io"

	"github.com/lemon-mint/vstruct/ir"
)

func writeEnums(w io.Writer, i *ir.IR) {
	for _, e := range i.Enums {
		fmt.Fprintf(w, "public enum class %s(val index: %s) {\n", NameConv(e.Name), NumberConv(true, e.Size*8))
		for i, o := range e.Options {
			fmt.Fprintf(w, "\t%s(index = %du),\n", NameConv(o), i)
		}
		fmt.Fprintf(w, "}\n\n")
	}
}
