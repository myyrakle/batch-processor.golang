# batch-processor.golang

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
