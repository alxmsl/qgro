package qrencode

import "bytes"

type vector struct {
	bits []bool
}

func (v *vector) Length() int {
	return len(v.bits)
}

func (v *vector) Get(i int) bool {
	return v.bits[i]
}

func (v *vector) AppendBit(b bool) {
	v.bits = append(v.bits, b)
}

func (v *vector) Append(b, count int) {
	for i := uint(count); i > 0; i-- {
		v.AppendBit((b>>(i-1))&1 == 1)
	}
}

func (v *vector) AppendBits(b vector) {
	v.bits = append(v.bits, b.bits...)
}

func (v *vector) String() string {
	b := bytes.Buffer{}
	for i, l := 0, v.Length(); i < l; i++ {
		if v.Get(i) {
			b.WriteString("1")
		} else {
			b.WriteString("0")
		}
	}
	return b.String()
}
