# gobenchmarks

My collection of Go benchmarks

## Run All
```bash
go test -benchmem -bench .
```

## Run ConcatStr Benchmark

```bash
go test -benchmem -bench ConcatStr
```
## Run PrintStr Benchmark

```bash
go test -benchmem -bench PrintStr > result.txt
```

Remove Logs
```vim
:g/hello world/d
:g/hello   world/d
```
