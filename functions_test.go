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

