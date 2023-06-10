package test

import (
	"testing"

	ed "github.com/RSheremeta/gob-proto-bench/encode_decode"
	"github.com/RSheremeta/gob-proto-bench/proto_gen"
	. "github.com/RSheremeta/gob-proto-bench/structs"
)

func BenchmarkEncodeSingleComplexMap(b *testing.B) {
	b.Run("type=GOB struct_size=huge_complex_map", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := ed.EncodeGob(NewComplexAndHuge())
			_ = res
		}
	})

	b.Run("type=Proto struct_size=huge_complex_map", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := ed.EncodeProto(NewComplexAndHugePb())
			_ = res
		}
	})
}

func BenchmarkDecodeSingleComplexMap(b *testing.B) {
	b.Run("type=GOB struct_size=huge_complex_map", func(b *testing.B) {
		bytes := ed.EncodeGob(NewComplexAndHuge())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result ComplexAndHuge
			ed.DecodeGob(bytes, &result)
		}
	})

	b.Run("type=Proto struct_size=huge_complex_map", func(b *testing.B) {
		bytes := ed.EncodeProto(NewComplexAndHugePb())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.ComplexAndHuge
			ed.DecodeProto(bytes, &result)
		}
	})
}
