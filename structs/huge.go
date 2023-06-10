package structs

import "github.com/RSheremeta/gob-proto-bench/proto_gen"

type Huge struct {
	BigSlice        []*Big
	BiggerSlice     []*Big
	TheBiggestSlice []*Big
}

func NewHuge() Huge {
	return Huge{
		BigSlice:        NewBigPtrSlc(),
		BiggerSlice:     makeBigPtrSlc(256),
		TheBiggestSlice: makeBigPtrSlc(1024),
	}
}

func NewHugePtr() *Huge {
	return &Huge{
		BigSlice:        NewBigPtrSlc(),
		BiggerSlice:     makeBigPtrSlc(256),
		TheBiggestSlice: makeBigPtrSlc(1024),
	}
}

type HugeArray struct {
	HugeArr []*Huge
}

func NewHugeSlc() *HugeArray {
	size := 50
	res := make([]*Huge, size)

	for i := 0; i < size; i++ {
		res[i] = NewHugePtr()
	}

	return &HugeArray{HugeArr: res}
}

// this func is intended to serve and mock proto structs eligible for encoding/decoding

func NewHugePb() *proto_gen.Huge {
	return &proto_gen.Huge{
		BigSlice:        NewBigPbSlc(),
		BiggerSlice:     makeBigPbSlc(256),
		TheBiggestSlice: makeBigPbSlc(1024),
	}
}

func NewHugePbSlc() *proto_gen.HugeArray {
	return makeHugePbSlc(50)
}

func makeHugePbSlc(size int) *proto_gen.HugeArray {
	res := make([]*proto_gen.Huge, size)

	for i := 0; i < size; i++ {
		res[i] = NewHugePb()
	}

	return &proto_gen.HugeArray{HugeArr: res}
}
