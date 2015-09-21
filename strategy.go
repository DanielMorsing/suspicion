package suspicion

import (
	"math"
)

func (t *T) DrawBytes(n int) []byte {
	nbytes := t.index + n
	buf := t.buf[t.index:nbytes]
	t.index = nbytes
	return buf
}

func (t *T) DrawByte() byte {
	return t.DrawBytes(1)[0]
}

var evilfloats = [...]float64{
	math.Inf(1),
	math.Inf(-1),
	math.NaN(),
	math.MaxFloat64,
	-math.MaxFloat64,
	math.SmallestNonzeroFloat64,
	-math.SmallestNonzeroFloat64,
}

func (t *T) DrawFloat64() float64 {
	// [0, 255-32] is int floats
	// [255-32, 255] is EVIL! floats
	b := t.DrawByte()
	if b < 255-32 {
		intbytes := t.DrawBytes(7)
		var retint int64
		for _, b := range intbytes {
			retint = (retint << 8) | int64(b)
		}
		mask := ^int64(-1 << 52)
		retint = retint & mask
		return float64(retint)
	}
	return evilfloats[int(b)%len(evilfloats)]
}
