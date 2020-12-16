package alpaca

import (
	"encoding/binary"
	"net"
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

func isResolvable(call otto.FunctionCall) otto.Value {
	host := call.Argument(0).String()
	addrs, err := net.LookupHost(host)
	if err != nil {
		return otto.FalseValue()
	}
	if len(addrs) == 0 {
		return otto.FalseValue()
	}
	return otto.TrueValue()
}

func isInNet(call otto.FunctionCall) otto.Value {
	host := call.Argument(0).String()
	addrs, err := net.LookupIP(host)
	if err != nil {
		return otto.UndefinedValue()
	}
	var ip net.IP = nil
	for _, addr := range addrs {
		ip4 := addr.To4()
		if ip4 != nil {
			ip = ip4
			break
		}
	}
	if ip == nil {
		return otto.UndefinedValue()
	}
	prefix := net.ParseIP(call.Argument(1).String())
	mask := net.ParseIP(call.Argument(2).String())
	if ip == nil || prefix == nil || mask == nil {
		return otto.UndefinedValue()
	}
	ipnet := net.IPNet{IP: prefix, Mask: net.IPMask(mask)}
	if ipnet.Contains(ip) {
		return otto.TrueValue()
	} else {
		return otto.FalseValue()
	}
}

func dnsResolve(call otto.FunctionCall) otto.Value {
	host := call.Argument(0).String()
	addrs, err := net.LookupIP(host)
	if err != nil {
		return otto.UndefinedValue()
	}
	var ip net.IP = nil
	for _, addr := range addrs {
		ip4 := addr.To4()
		if ip4 != nil {
			ip = ip4
			break
		}
	}
	if ip == nil {
		return otto.UndefinedValue()
	}
	v, _ := otto.ToValue(ip.String())
	return v
}

func convertAddr(call otto.FunctionCall) otto.Value {
	ip := net.ParseIP(call.Argument(0).String())
	if ip == nil {
		return otto.UndefinedValue()
	}
	ip4 := ip.To4()
	if ip4 == nil {
		return otto.UndefinedValue()
	}
	num := binary.BigEndian.Uint32(ip4)
	v, _ := otto.ToValue(num)
	return v
}
