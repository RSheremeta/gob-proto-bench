package test

import (
	"testing"

	ed "github.com/RSheremeta/gob-proto-bench/encode_decode"
	"github.com/RSheremeta/gob-proto-bench/proto_gen"
	. "github.com/RSheremeta/gob-proto-bench/structs"
	"google.golang.org/protobuf/proto"
)

func BenchmarkEncodeSlice(b *testing.B) {
	
	// Gob

	gobBnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=tiny", target: NewTinySlc(), encFn: ed.EncodeGob},
		{name: "type=GOB struct_size=medium", target: NewMediumSlc(), encFn: ed.EncodeGob},
		{name: "type=GOB struct_size=big", target: NewBigSlc(), encFn: ed.EncodeGob},
		{name: "type=GOB struct_size=huge", target: NewHugeSlc(), encFn: ed.EncodeGob},
	}


	for _, bm := range gobBnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				res := bm.encFn(strct)
				_ = res
			}
		})
	}

	// Proto

	protoBnchmkrs := []struct {
		name   string
		target proto.Message
		encFn  func(proto.Message) []byte
	}{
		{name: "type=Proto struct_size=tiny", target: NewTinyPbSlc(), encFn: ed.EncodeProto},
		{name: "type=Proto struct_size=medium", target: NewMediumPbSlc(), encFn: ed.EncodeProto},
		{name: "type=Proto struct_size=big", target: NewBigPbSlc(), encFn: ed.EncodeProto},
		{name: "type=Proto struct_size=huge", target: NewHugePbSlc(), encFn: ed.EncodeProto},
	}

	for _, bm := range protoBnchmkrs {
		b.Run(bm.name, func(b *testing.B) {
			strct := bm.target
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				res := bm.encFn(strct)
				_ = res
			}
		})
	}
}

func BenchmarkDecodeSlice(b *testing.B) {
	// Gob

	b.Run("type=GOB struct_size=tiny", func(b *testing.B) {
		bytes := ed.EncodeGob(NewTinySlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result TinyArray
			ed.DecodeGob(bytes, &result)
		}
	})

	b.Run("type=GOB struct_size=medium", func(b *testing.B) {
		bytes := ed.EncodeGob(NewMediumSlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result MediumArray
			ed.DecodeGob(bytes, &result)
		}
	})

	b.Run("type=GOB struct_size=big", func(b *testing.B) {
		bytes := ed.EncodeGob(NewBigSlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result BigArray
			ed.DecodeGob(bytes, &result)
		}
	})

	b.Run("type=GOB struct_size=huge", func(b *testing.B) {
		bytes := ed.EncodeGob(NewHugeSlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result HugeArray
			ed.DecodeGob(bytes, &result)
		}
	})

	// Proto

	b.Run("type=Proto struct_size=tiny", func(b *testing.B) {
		bytes := ed.EncodeProto(NewTinyPbSlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.TinyArray
			ed.DecodeProto(bytes, &result)
		}
	})

	b.Run("type=Proto struct_size=medium", func(b *testing.B) {
		bytes := ed.EncodeProto(NewMediumPbSlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.MediumArray
			ed.DecodeProto(bytes, &result)
		}
	})

	b.Run("type=Proto struct_size=big", func(b *testing.B) {
		bytes := ed.EncodeProto(NewBigPbSlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.BigArray
			ed.DecodeProto(bytes, &result)
		}
	})

	b.Run("type=Proto struct_size=huge", func(b *testing.B) {
		bytes := ed.EncodeProto(NewHugePbSlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.HugeArray
			ed.DecodeProto(bytes, &result)
		}
	})

}
