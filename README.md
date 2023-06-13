# gob-proto-bench

#### Benchmarking Gob vs Protobuf.
This is Part 2 of the experiment of measuring Gob's (de)serialization performance. See Part 1 [here](https://github.com/RSheremeta/gob-serialization)
This repo is made just out of curiosity and has an experimental reasoning.

### Description
There are 4 structs: `Tiny`, `Medium`, `Big`, `Huge` and a struct called `ComplexAndHuge` which holds `type complexMap map[string]*Huge`.

Each of them has a respective size and complexity.
And each of them builds the following targets to run encoding & decoding benchmarks against: 
- `Single` means a single struct - eg - `Tiny{}`
- `Slice` means a slice of given struct - eg - `[]*Tiny`

### Usage
You can just look at and analyze the results I got while running on my machine by observing csv samples in **results_run/sample/** dir.

Also, feel free to play on your own by doing the following:

Run a corresponding **Make** command in terminal and observe the results in **results_run/** dir.
Otherwise run `make run-all` in order to run everything. Warning: beware as it's a very CPU/RAM consuming operation.

### Results
**Software and Hardware:**
The experiment is run on:
- MacBook Pro 2021 16" with macOS Ventura 13.2.1 
- 32gb RAM
- Apple M1 Pro CPU
- Go 1.20.3

Here is the full analysis of mine about all the results run: [Analysis.md](https://github.com/RSheremeta/gob-proto-bench/blob/master/Analysis.md)

[**map_ptr_slice_decode.csv:**](https://github.com/RSheremeta/gob-proto-bench/blob/master/results_run/sample/map_slice_decode.csv)

```bash
goos: darwin
goarch: arm64
pkg: github.com/RSheremeta/gob-proto-bench/test
BenchmarkDecodeSliceComplexMap/type=GOB_struct_size=huge_complex_map_slice-10         	      10	 573950896 ns/op	633277254 B/op	10645891 allocs/op
BenchmarkDecodeSliceComplexMap/type=Proto_struct_size=huge_complex_map_slice-10       	      10	2153558850 ns/op	2337695762 B/op	46637425 allocs/op
PASS
ok  	github.com/RSheremeta/gob-proto-bench/test	35.524s
```

[**all.csv:**](https://github.com/RSheremeta/gob-proto-bench/blob/master/results_run/sample/all.csv)

```bash
goos: darwin
goarch: arm64
pkg: github.com/RSheremeta/gob-proto-bench/test
BenchmarkEncodeSingleComplexMap/type=GOB_struct_size=huge_complex_map-10         	      10	  51486367 ns/op	84000476 B/op	  731909 allocs/op
BenchmarkEncodeSingleComplexMap/type=Proto_struct_size=huge_complex_map-10       	      10	  42282621 ns/op	43574969 B/op	  466007 allocs/op
BenchmarkDecodeSingleComplexMap/type=GOB_struct_size=huge_complex_map-10         	      10	  55438804 ns/op	36013430 B/op	 1064916 allocs/op
BenchmarkDecodeSingleComplexMap/type=Proto_struct_size=huge_complex_map-10       	      10	  43338646 ns/op	46754095 B/op	  932749 allocs/op
BenchmarkEncodeSliceComplexMap/type=GOB_struct_size=huge_complex_map_slice-10    	      10	 478913167 ns/op	809327492 B/op	 7318093 allocs/op
BenchmarkEncodeSliceComplexMap/type=Proto_struct_size=huge_complex_map_slice-10  	      10	2021125029 ns/op	2181639455 B/op	23299823 allocs/op
BenchmarkDecodeSliceComplexMap/type=GOB_struct_size=huge_complex_map_slice-10    	      10	 565821717 ns/op	633277498 B/op	10645893 allocs/op
BenchmarkDecodeSliceComplexMap/type=Proto_struct_size=huge_complex_map_slice-10  	      10	2124560267 ns/op	2337694348 B/op	46637417 allocs/op
BenchmarkEncodeSingle/type=GOB_struct_size=tiny-10                               	      10	      2708 ns/op	    1136 B/op	      20 allocs/op
BenchmarkEncodeSingle/type=GOB_struct_size=medium-10                             	      10	      4771 ns/op	    1720 B/op	      38 allocs/op
BenchmarkEncodeSingle/type=GOB_struct_size=big-10                                	      10	      6317 ns/op	    2792 B/op	      61 allocs/op
BenchmarkEncodeSingle/type=GOB_struct_size=huge-10                               	      10	    740108 ns/op	 1184470 B/op	    5407 allocs/op
BenchmarkEncodeSingle/type=Proto_struct_size=tiny-10                             	      10	       950.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkEncodeSingle/type=Proto_struct_size=medium-10                           	      10	      1017 ns/op	      48 B/op	       1 allocs/op
BenchmarkEncodeSingle/type=Proto_struct_size=big-10                              	      10	      1258 ns/op	     128 B/op	       1 allocs/op
BenchmarkEncodeSingle/type=Proto_struct_size=huge-10                             	      10	    464279 ns/op	  172032 B/op	       1 allocs/op
BenchmarkDecodeSingle/type=GOB_struct_size=tiny-10                               	      10	     11879 ns/op	    6776 B/op	     180 allocs/op
BenchmarkDecodeSingle/type=GOB_struct_size=medium-10                             	      10	     15504 ns/op	    8728 B/op	     237 allocs/op
BenchmarkDecodeSingle/type=GOB_struct_size=big-10                                	      10	     18975 ns/op	   11379 B/op	     317 allocs/op
BenchmarkDecodeSingle/type=GOB_struct_size=huge-10                               	      10	   1167321 ns/op	  732968 B/op	   21628 allocs/op
BenchmarkDecodeSingle/type=Proto_struct_size=tiny-10                             	      10	       204.2 ns/op	      72 B/op	       2 allocs/op
BenchmarkDecodeSingle/type=Proto_struct_size=medium-10                           	      10	       675.0 ns/op	     256 B/op	       6 allocs/op
BenchmarkDecodeSingle/type=Proto_struct_size=big-10                              	      10	       925.0 ns/op	     680 B/op	      14 allocs/op
BenchmarkDecodeSingle/type=Proto_struct_size=huge-10                             	      10	    906150 ns/op	  934972 B/op	   18652 allocs/op
BenchmarkEncodeSlice/type=GOB_struct_size=tiny-10                                	      10	      4858 ns/op	    1664 B/op	      25 allocs/op
BenchmarkEncodeSlice/type=GOB_struct_size=medium-10                              	      10	      7342 ns/op	    4504 B/op	      54 allocs/op
BenchmarkEncodeSlice/type=GOB_struct_size=big-10                                 	      10	     34996 ns/op	   48297 B/op	     272 allocs/op
BenchmarkEncodeSlice/type=GOB_struct_size=huge-10                                	      10	  34902979 ns/op	65902422 B/op	  266108 allocs/op
BenchmarkEncodeSlice/type=Proto_struct_size=tiny-10                              	      10	      1329 ns/op	      80 B/op	       1 allocs/op
BenchmarkEncodeSlice/type=Proto_struct_size=medium-10                            	      10	      2862 ns/op	     512 B/op	       1 allocs/op
BenchmarkEncodeSlice/type=Proto_struct_size=big-10                               	      10	     19642 ns/op	    6528 B/op	       1 allocs/op
BenchmarkEncodeSlice/type=Proto_struct_size=huge-10                              	      10	  22906321 ns/op	 8314889 B/op	       1 allocs/op
BenchmarkDecodeSlice/type=GOB_struct_size=tiny-10                                	      10	     12954 ns/op	    8035 B/op	     219 allocs/op
BenchmarkDecodeSlice/type=GOB_struct_size=medium-10                              	      10	     19488 ns/op	   11513 B/op	     322 allocs/op
BenchmarkDecodeSlice/type=GOB_struct_size=big-10                                 	      10	     66958 ns/op	   40086 B/op	    1133 allocs/op
BenchmarkDecodeSlice/type=GOB_struct_size=huge-10                                	      10	  54951358 ns/op	36011614 B/op	 1064871 allocs/op
BenchmarkDecodeSlice/type=Proto_struct_size=tiny-10                              	      10	      1104 ns/op	     544 B/op	      15 allocs/op
BenchmarkDecodeSlice/type=Proto_struct_size=medium-10                            	      10	      3179 ns/op	    2872 B/op	      66 allocs/op
BenchmarkDecodeSlice/type=Proto_struct_size=big-10                               	      10	     34592 ns/op	   35080 B/op	     708 allocs/op
BenchmarkDecodeSlice/type=Proto_struct_size=huge-10                              	      10	  43635229 ns/op	46749566 B/op	  932608 allocs/op
PASS
ok  	github.com/RSheremeta/gob-proto-bench/test	67.650s
```