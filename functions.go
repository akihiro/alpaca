package alpaca

import (
	"strings"

	"github.com/robertkrimen/otto"
)

func isPlainHostName(call otto.FunctionCall) otto.Value {
	host := call.Argument(0).String()
	v, _ := otto.ToValue(!strings.Contains(host, "."))
	return v
}
