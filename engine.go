package alpaca

import (
	"github.com/robertkrimen/otto"
)

type Engine interface {
}

type engine struct {
	engine *otto.Otto
}

func NewEngine(pacfile []byte) (Engine, error) {
	vm := otto.New()

	vm.Set("isPlainHostName", isPlainHostName)
	vm.Set("dnsDomainIs", dnsDomainIs)
	vm.Set("localHostOrDomainIs", localHostOrDomainIs)
	vm.Set("isResolvable", isResolvable)
	vm.Set("isInNet", isInNet)
	vm.Set("dnsResolve", dnsResolve)
	vm.Set("convert_addr", convertAddr)
	vm.Set("myIpAddress", myIpAddress)
	vm.Set("dnsDomainLevels", dnsDomainLevels)
	vm.Set("shExpMatch", shExpMatch)

	source, err := vm.Compile("", pacfile)
	if err != nil {
		return nil, err
	}
	_, err = vm.Run(source)
	return &engine{engine: vm}, err
}
