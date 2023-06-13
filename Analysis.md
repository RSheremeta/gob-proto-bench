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


### Slice of structs


- Encoding:
- **Tiny**:
        - Gob holds 2nd place, while Protobuf is on the top by all aspects here:
            - Gob's results: 4692 ns/op, 1665 B/op, 25 allocs/op,
            - Protobuf's results: 1717 ns/op,  80 B/op, 1 allocs/op.
- **Huge**:
        - Proto has shown the best results for encoding huge struct:
            - Protobuf: 24296700 ns/op, 8585216 B/op, 1 allocs/op.
            - Gob: 34797729 ns/op, 65902660 B/op, 266109 allocs/op,
            - ++summary++: Protobuf is ~1.5x faster, ~7.5x less memory consuming and, which is the most impressive: requires only 1 allocation per operation, while Gob needs ~266k for the same.

- Decoding:
- **Tiny**:  
        - Protobuf won here with the following results:
        - 783 ns/op, 544 B/op, 15 allocs/op
- **Huge**:
        - Protobuf also leads a decoding of huge struct with nice results:
            - 42068254 ns/op, 46749560 B/op, 932608 allocs/op
            - Gob wins in memory only - 36011551 B/op which is ~1.3x faster than Proto
