package suspicion_test

import (
	"testing"

	"github.com/DanielMorsing/suspicion"
)

func TestContext(t *testing.T) {
	suspicion.Run(t, func(t *suspicion.T) {
		f1 := t.DrawFloat64()
		f2 := t.DrawFloat64()
		res := f1 + f2
		if res-f2 != f1 {
			t.Failln("not commutative", f1, f2, res)
		}
	})
}
