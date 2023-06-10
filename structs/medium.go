package structs

import (
	"time"

	"github.com/RSheremeta/gob-proto-bench/proto_gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Medium struct {
	MediumID           int64
	FirstName          string
	LastName           string
	MediumMoneyBalance float64
	CreatedAt          *time.Time
	TinyField          *Tiny
}

func NewMedium() Medium {
	t := time.Now()

	return Medium{
		MediumID:           1,
		FirstName:          "foo",
		LastName:           "bar",
		MediumMoneyBalance: 10.24,
		CreatedAt:          &t,
		TinyField:          NewTinyPtr(),
	}
}

func NewMediumPtr() *Medium {
	t := time.Now()

	return &Medium{
		MediumID:           1,
		FirstName:          "foo",
		LastName:           "bar",
		MediumMoneyBalance: 10.24,
		CreatedAt:          &t,
		TinyField:          NewTinyPtr(),
	}
}

type MediumArray struct {
	MediumArr []*Medium
}

func NewMediumSlc() *MediumArray {
	size := 10
	res := make([]*Medium, size)

	for i := 0; i < size; i++ {
		res[i] = NewMediumPtr()
	}

	return &MediumArray{MediumArr: res}
}

// these funcs are intended to serve and mock proto structs eligible for encoding/decoding

func NewMediumPb() *proto_gen.Medium {
	t := timestamppb.New(time.Now())

	return &proto_gen.Medium{
		MediumId:           1,
		FirstName:          "foo",
		LastName:           "bar",
		MediumMoneyBalance: 10.24,
		CreatedAt:          t,
		TinyField:          NewTinyPb(),
	}
}

func NewMediumPbSlc() *proto_gen.MediumArray {
	size := 10
	res := make([]*proto_gen.Medium, size)

	for i := 0; i < size; i++ {
		res[i] = NewMediumPb()
	}

	return &proto_gen.MediumArray{MediumArr: res}
}
