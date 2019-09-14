run:
	go build && ./blockchain-go

docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	docker build -t blockchain-go .

init:
	GO111MODULE=on go mod init github.com/ragul28/blockchain-go
	GO111MODULE=on go get -u

mod:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod vendor
