# batch-processor.golang

![](https://img.shields.io/badge/language-Go-00ADD8) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

## install

```
go mod tidy
```

## run

```
go run main.go
```

## build

```
# unix like
go build main.go -o batch

# windows
go build main.go -o batch.exe
```

## How to use

```
# local test
go run main.go -p [process name]

# deploy mode
./batch -p [process name]
```
