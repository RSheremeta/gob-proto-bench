package structs

import (
	"fmt"

	"github.com/RSheremeta/gob-proto-bench/proto_gen"
)

type ComplexMap map[string]*Huge

type ComplexAndHuge struct {
	ComplexMap ComplexMap
}

func NewComplexAndHuge() ComplexAndHuge {
	return ComplexAndHuge{
		ComplexMap: makeComplexMap(50),
	}
}

func NewComplexAndHugePtr() *ComplexAndHuge {
	return &ComplexAndHuge{
		ComplexMap: makeComplexMap(50),
	}
}

type ComplexAndHugeArray struct {
	ComplexAndHugeArr []*ComplexAndHuge
}

func NewComplexAndHugeSlc() *ComplexAndHugeArray {
	size := 10
	var res []*ComplexAndHuge

	for i := 0; i < size; i++ {
		res = append(res, NewComplexAndHugePtr())
	}

	return &ComplexAndHugeArray{ComplexAndHugeArr: res}
}

func makeComplexMap(size int) ComplexMap {
	res := make(ComplexMap, size)

	for i := 0; i < size; i++ {
		k := fmt.Sprintf("%v", i)
		res[k] = NewHugePtr()
	}

	return res
}

// these funcs are intended to serve and mock proto structs eligible for encoding/decoding
type ComplexMapPb map[string]*proto_gen.Huge

func NewComplexAndHugePb() *proto_gen.ComplexAndHuge {
	return &proto_gen.ComplexAndHuge{
		ComplexMap: makeComplexMapPb(50),
	}
}

func NewComplexAndHugePbSlc() *proto_gen.ComplexAndHugeArray {
	return makeComplexAndHugePbSlc(50)
}

func makeComplexAndHugePbSlc(size int) *proto_gen.ComplexAndHugeArray {
	res := make([]*proto_gen.ComplexAndHuge, size)

	for i := 0; i < size; i++ {
		res[i] = NewComplexAndHugePb()
	}

	return &proto_gen.ComplexAndHugeArray{ComplexArr: res}
}

func makeComplexMapPb(size int) ComplexMapPb {
	res := make(ComplexMapPb, size)

	for i := 0; i < size; i++ {
		k := fmt.Sprintf("%v", i)
		res[k] = NewHugePb()
	}

	return res
}
