package suspicion

import (
	"fmt"
	"testing"
)

type T struct {
	t     *testing.T
	buf   []byte
	index int
}

func Run(t *testing.T, f func(*T)) {
	// generate 200 buffers
	// with random bytes in them
	for i := 0; i < 200; i++ {
		buf := generateBuf(64 << 10)
		tcontext := &T{
			t:   t,
			buf: buf,
		}
		f(tcontext)
	}
}

func (t *T) Failln(i ...interface{}) {
	fmt.Println(i...)
	t.t.Fail()
}
