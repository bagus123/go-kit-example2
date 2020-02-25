
# go-kit-example2


## Installation

```shell
# 1. clone repository
git clone https://github.com/bagus123/go-kit-example2.git

# 2. downloads all dependencies and build binary
go build -o bin/app cmd/main.go

# run from binary
./bin/app 

# or run from source
go run cmd/main.go 

# note
# clean cache go
go clean --modcache

# remove unused module
go mod tidy
```


## manual

```shell

# add library go-kit
go get github.com/go-kit/kit

# add library mux (http)
go get github.com/gorilla/mux

```