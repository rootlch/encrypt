package encrypt

import (
	"bytes"
	"errors"
	"io"
)

var (
	ErrKeyEmpty = errors.New("Key cannot be empty")
)

type Xor struct {
	index int
	key   []rune
}

func NewXor(key string) (x Xor, err error) {
	err = ErrKeyEmpty
	if key == "" {
		return
	}

	return Xor{index: 0, key: runes(key)}, nil
}

func (x *Xor) NewReader(rd io.Reader) io.Reader {
	var temp bytes.Buffer
	io.Copy(&temp, rd)
	result := bytes.NewBuffer(x.Encode(temp.Bytes()))
	return result
}

func (x *Xor) Encode(src []byte) (b []byte) {
	for _, v := range src {
		b = append(b, x.xor(v))
	}

	return
}

//Decodes and encodes.
func (x *Xor) xor(src byte) byte {
	src = src ^ byte(x.key[x.index])
	x.index = (x.index + 1) % len(x.key)
	return src
}

//Resets index back to 0.
func (x *Xor) Reset() {
	x.index = 0
}

func runes(s string) (r []rune) {
	for _, v := range s {
		r = append(r, v)
	}
	return
}
