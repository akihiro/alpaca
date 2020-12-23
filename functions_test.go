package alpaca

import (
	"testing"
)

func testFunc(t *testing.T, testCases []string) {
	e, err := NewEngine(nil)
	if err != nil {
		t.Error(err)
	}
	vm := e.(*engine).engine
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
	testFunc(t, tcs)
}

func TestDnsDomainIs(t *testing.T) {
	tcs := []string{
		`dnsDomainIs("www.mozilla.org", ".mozilla.org") === true`,
		`dnsDomainIs("www", ".mozilla.org") === false`,
	}
	testFunc(t, tcs)
}

func TestLocalHostOrDomainsIs(t *testing.T) {
	tcs := []string{
		`localHostOrDomainIs("www.mozilla.org" , "www.mozilla.org") === true`,
		`localHostOrDomainIs("www"             , "www.mozilla.org") === true`,
		`localHostOrDomainIs("www.google.com"  , "www.mozilla.org") === false`,
		`localHostOrDomainIs("home.mozilla.org", "www.mozilla.org") === false`,
	}
	testFunc(t, tcs)
}

func TestIsResolvable(t *testing.T) {
	tcs := []string{
		`isResolvable("www.mozilla.org") === true`,
		`isResolvable("notfound.example.com") === false`,
	}
	testFunc(t, tcs)
}

func TestIsInNet(t *testing.T) {
	tcs := []string{
		`isInNet("localhost", "127.0.0.0", "255.0.0.0") === true`,
		`isInNet("192.168.0.1", "192.168.0.0", "255.255.255.0") === true`,
		`isInNet("172.16.0.1", "192.168.0.0", "255.255.255.0") === false`,
	}
	testFunc(t, tcs)
}

func TestDnsResolve(t *testing.T) {
	tcs := []string{
		`dnsResolve("localhost") === "127.0.0.1"`,
		`dnsResolve("example.com") !== "1.1.1.1"`,
	}
	testFunc(t, tcs)
}

func TestConvertAddr(t *testing.T) {
	tcs := []string{
		`convert_addr("104.16.41.2") === 1745889538`,
		`convert_addr("127.0.0.1") === 2130706433`,
	}
	testFunc(t, tcs)
}

func TestMyIpAddress(t *testing.T) {
	tcs := []string{
		`myIpAddress() === "127.0.0.1"`,
	}
	testFunc(t, tcs)
}

func TestDnsDomainLevels(t *testing.T) {
	tcs := []string{
		`dnsDomainLevels("www") === 0`,
		`dnsDomainLevels("mozilla.org") === 1`,
		`dnsDomainLevels("www.mozilla.org") === 2`,
	}
	testFunc(t, tcs)
}

func TestShExpMatch(t *testing.T) {
	tcs := []string{
		`shExpMatch("http://home.netscape.com/people/ari/index.html"     , "*/ari/*") === true`,
		`shExpMatch("http://home.netscape.com/people/montulli/index.html", "*/ari/*") === false`,
	}
	testFunc(t, tcs)
}
