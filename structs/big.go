package structs

import (
	"time"

	"github.com/RSheremeta/gob-proto-bench/proto_gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Big struct {
	BigID        int64
	Age          uint64
	Country      string
	CountryBytes []byte
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
	TinyField    *Tiny
	MediumField  *Medium
}

func NewBig() Big {
	t := time.Now()

	return Big{
		BigID:        1,
		Age:          uint64(26),
		Country:      "Ukraine",
		CountryBytes: []byte("Ukraine"),
		CreatedAt:    &t,
		UpdatedAt:    &t,
		DeletedAt:    &t,
		TinyField:    NewTinyPtr(),
		MediumField:  NewMediumPtr(),
	}
}

func NewBigPtr() *Big {
	t := time.Now()

	return &Big{
		BigID:        1,
		Age:          uint64(26),
		Country:      "Ukraine",
		CountryBytes: []byte("Ukraine"),
		CreatedAt:    &t,
		UpdatedAt:    &t,
		DeletedAt:    &t,
		TinyField:    NewTinyPtr(),
		MediumField:  NewMediumPtr(),
	}
}

func NewBigPtrSlc() []*Big {
	return makeBigPtrSlc(50)
}

func makeBigPtrSlc(size int) []*Big {
	res := make([]*Big, size)

	for i := 0; i < size; i++ {
		res[i] = NewBigPtr()
	}

	return res
}

type BigArray struct {
	BigArr []*Big
}

func NewBigSlc() *BigArray {
	return makeBigSlc(50)
}

func makeBigSlc(size int) *BigArray {
	res := make([]*Big, size)

	for i := 0; i < size; i++ {
		res[i] = NewBigPtr()
	}

	return &BigArray{BigArr: res}
}

// these funcs are intended to serve and mock proto structs eligible for encoding/decoding

func NewBigPb() *proto_gen.Big {
	t := timestamppb.New(time.Now())

	return &proto_gen.Big{
		BigId:        1,
		Age:          uint64(26),
		Country:      "Ukraine",
		CountryBytes: []byte("Ukraine"),
		CreatedAt:    t,
		UpdatedAt:    t,
		DeletedAt:    t,
		TinyField:    NewTinyPb(),
		MediumField:  NewMediumPb(),
	}
}

func NewBigPbSlc() *proto_gen.BigArray {
	return makeBigPbSlc(50)
}

func makeBigPbSlc(size int) *proto_gen.BigArray {
	res := make([]*proto_gen.Big, size)

	for i := 0; i < size; i++ {
		res[i] = NewBigPb()
	}

	return &proto_gen.BigArray{BigArr: res}
}
