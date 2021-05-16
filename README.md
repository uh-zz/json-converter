# json-converter

json-converter, is conversion package for json, and is 20 times slower than encoding/json.

## Feature

- [x] decode from json string to go struct
- [ ] encode from go struct to json string

## Usage

```bash
git clone https://github.com/uh-zz/json-converter
cd json-converter
go build
./json-converter

```

## Benchmark
- goos: darwin
- goarch: amd64
- cpu: Intel(R) Core(TM) i5-8210Y CPU @ 1.60GHz


| 日付 | パッケージ | 実行した回数 | 1 回あたりの実行にかかった時間(ns/op) | 1 回あたりのアロケーションで確保した容量(B/op) | 1 回あたりのアロケーション回数(allocs/op) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| 2021/05/16 | encoding/json | 983222 | 1104 | 264 | 6 |
| 2021/05/16 | json-converter | 56949 | 20886 | 13258 | 185 |

## Licence

Apache2
