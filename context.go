package suspicion

import (
	"testing"
	"fmt"
)

type T struct {
	t *testing.T
}

func Run(t *testing.T, f func(*T),) {
	f(&T{t})
	fmt.Println("foo")
}