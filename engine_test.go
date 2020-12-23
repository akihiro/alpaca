package alpaca

import (
	"reflect"
	"testing"

	"github.com/robertkrimen/otto"
)

func TestParaseScriptResponse(t *testing.T) {
	tcs := []struct {
		Input  string
		Expect []Endpoint
	}{
		{
			Input:  `DIRECT`,
			Expect: []Endpoint{Endpoint{Type: Direct}},
		},
		{
			Input: `PROXY proxy.example.com`,
			Expect: []Endpoint{
				Endpoint{Type: Proxy, Host: "proxy.example.com"},
			},
		},
		{
			Input: `HTTP proxy.example.com`,
			Expect: []Endpoint{
				Endpoint{Type: Http, Host: "proxy.example.com"},
			},
		},
		{
			Input: `HTTPS proxy.example.com`,
			Expect: []Endpoint{
				Endpoint{Type: Https, Host: "proxy.example.com"},
			},
		},
		{
			Input: `SOCKS proxy.example.com`,
			Expect: []Endpoint{
				Endpoint{Type: Socks, Host: "proxy.example.com"},
			},
		},
		{
			Input: `SOCKS4 proxy.example.com`,
			Expect: []Endpoint{
				Endpoint{Type: Socks4, Host: "proxy.example.com"},
			},
		},
		{
			Input: `SOCKS6 proxy.example.com`,
			Expect: []Endpoint{
				Endpoint{Type: Socks6, Host: "proxy.example.com"},
			},
		},
		{
			Input: `SOCKS proxy.example.com; DIRECT`,
			Expect: []Endpoint{
				Endpoint{Type: Socks, Host: "proxy.example.com"},
				Endpoint{Type: Direct},
			},
		},
	}
	for i, v := range tcs {
		input, _ := otto.ToValue(v.Input)
		result, err := parseScriptResponse(input)
		if err != nil {
			t.Errorf("[%d] %w", i, err)
		}
		if !reflect.DeepEqual(result, v.Expect) {
			t.Errorf("[%d] mismatch %#v", i, result)
		}
	}
}
