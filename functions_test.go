package alpaca

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func testFunc(t *testing.T, testCases []string, name string, f func(otto.FunctionCall) otto.Value) {
	vm := otto.New()
	vm.Set(name, f)
	for i, v := range testCases {
		val, err := vm.Run(v)
		if err != nil {
			t.Error(err)
		}
		b, err := val.ToBoolean()
		if err != nil {
			t.Error(err)
		}
		if b != true {
			t.Errorf("[%d] mismatch: %#v", i, val)
		}
	}
}

func TestIsPlainHostName(t *testing.T) {
	tcs := []string{
		`isPlainHostName("www.mozilla.org") === false`,
		`isPlainHostName("www") === true`,
	}
	testFunc(t, tcs, "isPlainHostName", isPlainHostName)
}

func TestDnsDomainIs(t *testing.T) {
	tcs := []string{
		`dnsDomainIs("www.mozilla.org", ".mozilla.org") === true`,
		`dnsDomainIs("www", ".mozilla.org") === false`,
	}
	testFunc(t, tcs, "dnsDomainIs", dnsDomainIs)
}

func TestLocalHostOrDomainsIs(t *testing.T) {
	tcs := []string{
		`localHostOrDomainIs("www.mozilla.org" , "www.mozilla.org") === true`,
		`localHostOrDomainIs("www"             , "www.mozilla.org") === true`,
		`localHostOrDomainIs("www.google.com"  , "www.mozilla.org") === false`,
		`localHostOrDomainIs("home.mozilla.org", "www.mozilla.org") === false`,
	}
	testFunc(t, tcs, "localHostOrDomainIs", localHostOrDomainIs)
}

func TestIsResolvable(t *testing.T) {
	tcs := []string{
		`isResolvable("www.mozilla.org") === true`,
		`isResolvable("notfound.example.com") === false`,
	}
	testFunc(t, tcs, "isResolvable", isResolvable)
}
