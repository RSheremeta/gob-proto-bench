package structs

import "github.com/RSheremeta/gob-proto-bench/proto_gen"

type Tiny struct {
	TinyID  int64
	SomeStr string
}

func NewTiny() Tiny {
	return Tiny{
		TinyID:  1,
		SomeStr: "foo bar",
	}
}

func NewTinyPtr() *Tiny {
	return &Tiny{
		TinyID:  1,
		SomeStr: "foo bar",
	}
}

type TinyArray struct {
	TinyArr []*Tiny
}

func NewTinySlc() *TinyArray {
	size := 5
	res := make([]*Tiny, size)

	for i := 0; i < size; i++ {
		res[i] = NewTinyPtr()
	}

	return &TinyArray{TinyArr: res}
}

// these funcs are intended to serve and mock proto structs eligible for encoding/decoding

func NewTinyPb() *proto_gen.Tiny {
	return &proto_gen.Tiny{
		TinyId:  1,
		SomeStr: "foo bar",
	}
}

func NewTinyPbSlc() *proto_gen.TinyArray {
	size := 5
	res := make([]*proto_gen.Tiny, size)

	for i := 0; i < size; i++ {
		res[i] = NewTinyPb()
	}

	return &proto_gen.TinyArray{TinyArr: res}
}
