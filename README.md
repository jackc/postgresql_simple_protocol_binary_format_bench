# PostgreSQL Simple Protocol Binary Format Benchmark

This is a benchmark testing the change proposed in
https://github.com/postgresql-interfaces/enhancement-ideas/discussions/5 as implemented by
https://github.com/davecramer/postgres/tree/format_binary.

It executes a simple select that returns an `int4` column, a `text` column, and a `timestamptz` column.

## Results

```
jack@glados ~/dev/postgresql_simple_protocol_binary_format_bench » DATABASE_URL='database=postgres port=15432' go test -bench=. -benchtime=5s
goos: darwin
goarch: amd64
pkg: github.com/jackc/postgresql_simple_protocol_binary_format_bench
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkTextFormat1Row-16         	  124485	     47964 ns/op
BenchmarkBinaryFormat1Row-16       	  128959	     46008 ns/op
BenchmarkTextFormat100Rows-16      	   41881	    144596 ns/op
BenchmarkBinaryFormat100Rows-16    	   61512	     97302 ns/op
PASS
ok  	github.com/jackc/postgresql_simple_protocol_binary_format_bench	27.545s
```

At 100 rows the text format is 48% slower than the binary format.
