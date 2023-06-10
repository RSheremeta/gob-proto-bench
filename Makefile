run-all: 
		cd test \
		&& go test -bench=. -benchmem -benchtime=10x > ../results_run/all.csv \
		&& cd ..

run-single:
		cd test \
		&& go test -bench=BenchmarkEncodeSingle -benchmem -benchtime=10x > ../results_run/single_encode.csv \
		&& go test -bench=BenchmarkDecodeSingle -benchmem -benchtime=10x > ../results_run/single_decode.csv \
		&& cd ..

run-slice:
		cd test \
		&& go test -bench=BenchmarkEncodeSlice -benchmem -benchtime=10x > ../results_run/slice_encode.csv \
		&& go test -bench=BenchmarkDecodeSlice -benchmem -benchtime=10x > ../results_run/slice_decode.csv \
		&& cd ..

run-map:
		cd test \
		&& go test -bench=BenchmarkEncodeSingleComplexMap -benchmem -benchtime=10x > ../results_run/map_single_encode.csv \
		&& go test -bench=BenchmarkDecodeSingleComplexMap -benchmem -benchtime=10x > ../results_run/map_single_decode.csv \
		&& go test -bench=BenchmarkEncodeSliceComplexMap -benchmem -benchtime=10x > ../results_run/map_slice_encode.csv \
		&& go test -bench=BenchmarkDecodeSliceComplexMap -benchmem -benchtime=10x > ../results_run/map_slice_decode.csv \
		&& cd ..

generate:
	protoc -Iproto --go_opt=module=github.com/RSheremeta/gob_proto_bench --go_out=. proto/*.proto

clean:
	rm proto_gen/*.pb.go

.PHONY: generate clean run-all run-single run-slice run-map