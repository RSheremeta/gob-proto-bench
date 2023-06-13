## Analysis of results run 

This is my extended analysis of all runs.
Feel free to reach out if you spotted an inconsistency or a mistake.

#### Spoiler
I've omitted highlighting ``Medium`` and ``Big`` structs' results, as the most interesting is in the smallest and the biggest stuff here, at least IMHO.

### Single struct

- Encoding:
- **Tiny**:
        Gob loses Protobuf here, as it takes ~5x more time (2762 ns/op vs 587 ns/op in Protobuf), ~71x more memory (1137 B/op vs 16 B/op in Protobuf), ~20x more memory allocations (20 allocs/op vs 1 allocs/op in Protobuf)
- **Huge**:
        Protobuf is *better* than Gob here:
        it's 1.5x faster (456542 ns/op vs Gob's 687012 ns/op),
        ~7x less memory (172032 B/op vs 1184441 B/op in Gob),
        the memory allocation difference is crucial: 1 allocs/op for Proto vs 5407 allocs/op for Gob

- Decoding:
- **Tiny**:
        Gob *loses* while Protobuf is the best here:
        Gob decodes ~23x slower than Protobuf (11621 ns/op vs 504 ns/op),
        ~94x more memory needed (6776 B/op vs 72 B/op in Protobuf),
        ~90x more allocations (180 allocs/op vs Proto's 2 allocs/op)
- **Huge**:
        Protobuf has shown mostly better performance here:
        Protobuf's processing time is 893196 ns/op vs 1144283 ns/op for Gob, which is ~1.3x faster,
        Protobuf's amount of allocs per operation is granularly faster (~1.15x - 18652 allocs/op vs Gob's 21628 allocs/op)
        however, Gob requires a bit less memory: 732947 B/op vs 934971 B/op for Proto, which is ~1.3x faster
