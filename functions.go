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

func dnsDomainIs(call otto.FunctionCall) otto.Value {
	host := call.Argument(0).String()
	domain := call.Argument(1).String()
	tokens := strings.SplitN(host, ".", 2)
	if len(tokens) == 1 {
		return otto.FalseValue()
	}
	if "."+tokens[1] == domain {
		return otto.TrueValue()
	} else {
		return otto.FalseValue()
	}
}

func localHostOrDomainIs(call otto.FunctionCall) otto.Value {
	host := call.Argument(0).String()
	hostdom := call.Argument(1).String()
	if host == hostdom {
		return otto.TrueValue()
	}
	tokens := strings.SplitN(host, ".", 2)
	if len(tokens) == 1 {
		if host == strings.SplitN(hostdom, ".", 2)[0] {
			return otto.TrueValue()
		} else {
			return otto.FalseValue()
		}
	}
	return otto.FalseValue()
}
