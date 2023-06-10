package test

import (
	"testing"

	ed "github.com/RSheremeta/gob-proto-bench/encode_decode"
	"github.com/RSheremeta/gob-proto-bench/proto_gen"
	. "github.com/RSheremeta/gob-proto-bench/structs"
	"google.golang.org/protobuf/proto"
)

func BenchmarkEncodeSingle(b *testing.B) {
	// Gob

	gobBnchmkrs := []struct {
		name   string
		target any
		encFn  func(any) []byte
	}{
		{name: "type=GOB struct_size=tiny", target: NewTiny(), encFn: ed.EncodeGob},
		{name: "type=GOB struct_size=medium", target: NewMedium(), encFn: ed.EncodeGob},
		{name: "type=GOB struct_size=big", target: NewBig(), encFn: ed.EncodeGob},
		{name: "type=GOB struct_size=huge", target: NewHuge(), encFn: ed.EncodeGob},
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
		{name: "type=Proto struct_size=tiny", target: NewTinyPb(), encFn: ed.EncodeProto},
		{name: "type=Proto struct_size=medium", target: NewMediumPb(), encFn: ed.EncodeProto},
		{name: "type=Proto struct_size=big", target: NewBigPb(), encFn: ed.EncodeProto},
		{name: "type=Proto struct_size=huge", target: NewHugePb(), encFn: ed.EncodeProto},
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

func BenchmarkDecodeSingle(b *testing.B) {
	// Gob

	b.Run("type=GOB struct_size=tiny", func(b *testing.B) {
		bytes := ed.EncodeGob(NewTiny())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result Tiny
			ed.DecodeGob(bytes, &result)
		}
	})

	b.Run("type=GOB struct_size=medium", func(b *testing.B) {
		bytes := ed.EncodeGob(NewMedium())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result Medium
			ed.DecodeGob(bytes, &result)
		}
	})

	b.Run("type=GOB struct_size=big", func(b *testing.B) {
		bytes := ed.EncodeGob(NewBig())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result Big
			ed.DecodeGob(bytes, &result)
		}
	})

	b.Run("type=GOB struct_size=huge", func(b *testing.B) {
		bytes := ed.EncodeGob(NewHuge())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result Huge
			ed.DecodeGob(bytes, &result)
		}
	})

	// Proto

	b.Run("type=Proto struct_size=tiny", func(b *testing.B) {
		bytes := ed.EncodeProto(NewTinyPb())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.Tiny
			ed.DecodeProto(bytes, &result)
		}
	})

	b.Run("type=Proto struct_size=medium", func(b *testing.B) {
		bytes := ed.EncodeProto(NewMediumPb())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.Medium
			ed.DecodeProto(bytes, &result)
		}
	})

	b.Run("type=Proto struct_size=big", func(b *testing.B) {
		bytes := ed.EncodeProto(NewBigPb())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.Big
			ed.DecodeProto(bytes, &result)
		}
	})

	b.Run("type=Proto struct_size=huge", func(b *testing.B) {
		bytes := ed.EncodeProto(NewHugePb())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.Huge
			ed.DecodeProto(bytes, &result)
		}
	})

}
