from golang:alpine
maintainer jesse.miller@adops.com

run apk add --no-cache git

run mkdir -p /go/src/perseus
workdir /go/src/perseus

run go get -u github.com/alecthomas/gometalinter
run gometalinter --install

copy ./main.go /go/src/perseus
run go-wrapper download

run apk del git

copy . /go/src/perseus/
run go-wrapper install

run go test ./...
run gometalinter --disable-all --enable=golint --enable=gocyclo --enable=gofmt ./...

expose 2120

cmd ["go-wrapper", "run"]