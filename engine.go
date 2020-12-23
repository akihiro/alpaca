package alpaca

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/robertkrimen/otto"
)

var (
	ErrorNotString = errors.New("return value not string")
)

type ProxyType int

const (
	Direct ProxyType = iota
	Proxy
	Socks
	Http
	Https
	Socks4
	Socks6
)

func ParseProxyType(s string) (ProxyType, error) {
	switch s {
	case "DIRECT":
		return Direct, nil
	case "PROXY":
		return Proxy, nil
	case "SOCKS":
		return Socks, nil
	case "HTTP":
		return Http, nil
	case "HTTPS":
		return Https, nil
	case "SOCKS4":
		return Socks4, nil
	case "SOCKS6":
		return Socks6, nil
	default:
		return ProxyType(0), fmt.Errorf(`invalid format: "%s"`, s)
	}
}

func (p ProxyType) String() string {
	switch p {
	case Direct:
		return "DIRECT"
	case Socks:
		return "SOCKS"
	case Http:
		return "HTTP"
	case Https:
		return "HTTPS"
	case Socks4:
		return "SOCKS4"
	case Socks6:
		return "SOCKS6"
	default:
		return ""
	}
}

type Endpoint struct {
	Type ProxyType
	Host string
}

type Engine interface {
	FindProxyForURL(url.URL) ([]Endpoint, error)
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

func (e *engine) FindProxyForURL(u url.URL) ([]Endpoint, error) {
	host := strings.SplitN(u.Host, ":", 2)[0]
	val, err := e.engine.Call("FindProxyForURL", nil, u.String(), host)
	if err != nil {
		return nil, err
	}
	return parseScriptResponse(val)
}

func parseScriptResponse(val otto.Value) ([]Endpoint, error) {
	if val.IsNull() {
		return []Endpoint{{Type: Direct}}, nil
	}
	if !val.IsString() {
		return nil, ErrorNotString
	}
	ret := val.String()
	var endpoints []Endpoint
	for _, v := range strings.Split(ret, ";") {
		tokens := strings.SplitN(strings.TrimSpace(v), " ", 2)
		var ep Endpoint
		switch len(tokens) {
		case 1:
			if strings.TrimSpace(tokens[0]) != "DIRECT" {
				return nil, fmt.Errorf(`invalid token: "%s" : "%s"`, tokens[0], ret)
			}
			ep.Type = Direct
		case 2:
			typ, err := ParseProxyType(strings.TrimSpace(tokens[0]))
			if err != nil {
				return nil, fmt.Errorf(`invalid format: %w : "%s"`, err, ret)
			}
			ep.Type = typ
			ep.Host = strings.TrimSpace(tokens[1])
		default:
			return nil, fmt.Errorf(`invalid token: "%s" : "%s"`, v, ret)
		}
		endpoints = append(endpoints, ep)
	}
	return endpoints, nil
}
