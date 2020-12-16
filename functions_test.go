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

func TestIsInNet(t *testing.T) {
	tcs := []string{
		`isInNet("localhost", "127.0.0.0", "255.0.0.0") === true`,
		`isInNet("192.168.0.1", "192.168.0.0", "255.255.255.0") === true`,
		`isInNet("172.16.0.1", "192.168.0.0", "255.255.255.0") === false`,
	}
	testFunc(t, tcs, "isInNet", isInNet)
}

func TestDnsResolve(t *testing.T) {
	tcs := []string{
		`dnsResolve("localhost") === "127.0.0.1"`,
		`dnsResolve("example.com") !== "1.1.1.1"`,
	}
	testFunc(t, tcs, "dnsResolve", dnsResolve)
}

func TestConvertAddr(t *testing.T) {
	tcs := []string{
		`convert_addr("104.16.41.2") === 1745889538`,
		`convert_addr("127.0.0.1") === 2130706433`,
	}
	testFunc(t, tcs, "convert_addr", convertAddr)
}
