# json-converter

## feature

- [x] decode from json string to go struct
- [ ] encode from go struct to json string

## how to use

```bash
git clone https://github.com/uh-zz/json-converter
cd json-converter
go build
./json-converter


# benchmark
goos: darwin
goarch: amd64
pkg: github.com/uh-zz/json-converter
cpu: Intel(R) Core(TM) i5-8210Y CPU @ 1.60GHz

# encoding/json
BenchmarkEncodingJson-4           983222              1104 ns/op             264 B/op          6 allocs/op

# json-converter
BenchmarkJsonConverter-4           56949             20886 ns/op           13258 B/op        185 allocs/op

実行した回数
1 回あたりの実行に掛かった時間(ns/op)
1 回あたりのアロケーションで確保した容量(B/op)
1 回あたりのアロケーション回数(allocs/op)

```

## Licence

Apache2
