package test

import (
	"testing"

	ed "github.com/RSheremeta/gob-proto-bench/encode_decode"
	"github.com/RSheremeta/gob-proto-bench/proto_gen"
	. "github.com/RSheremeta/gob-proto-bench/structs"
)

func BenchmarkEncodeSliceComplexMap(b *testing.B) {
	b.Run("type=GOB struct_size=huge_complex_map_slice", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := ed.EncodeGob(NewComplexAndHugeSlc())
			_ = res
		}
	})

	b.Run("type=Proto struct_size=huge_complex_map_slice", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res := ed.EncodeProto(NewComplexAndHugePbSlc())
			_ = res
		}
	})
}

func BenchmarkDecodeSliceComplexMap(b *testing.B) {
	b.Run("type=GOB struct_size=huge_complex_map_slice", func(b *testing.B) {
		bytes := ed.EncodeGob(NewComplexAndHugeSlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result ComplexAndHugeArray
			ed.DecodeGob(bytes, &result)
		}
	})

	b.Run("type=Proto struct_size=huge_complex_map_slice", func(b *testing.B) {
		bytes := ed.EncodeProto(NewComplexAndHugePbSlc())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result proto_gen.ComplexAndHugeArray
			ed.DecodeProto(bytes, &result)
		}
	})
}
