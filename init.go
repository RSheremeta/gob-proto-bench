package main

import (
	"encoding/gob"

	. "github.com/RSheremeta/gob-proto-bench/structs"
)

func init() {
	// single
	gob.Register(Tiny{})
	gob.Register(Medium{})
	gob.Register(Big{})
	gob.Register(Huge{})
	// slices
	gob.Register([]Tiny{})
	gob.Register([]Medium{})
	gob.Register([]Big{})
	gob.Register([]Huge{})
	// slices of pointers
	gob.Register([]*Tiny{})
	gob.Register([]*Medium{})
	gob.Register([]*Big{})
	gob.Register([]*Huge{})
	// all map stuff
	gob.Register(ComplexAndHuge{})
	gob.Register([]ComplexAndHuge{})
	gob.Register([]*ComplexAndHuge{})
	gob.Register(ComplexAndHugeArray{})
	gob.Register([]ComplexAndHugeArray{})
	gob.Register([]*ComplexAndHugeArray{})
}
