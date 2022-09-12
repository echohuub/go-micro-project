
GOPATH:=$(shell go env GOPATH)
#.PHONY: init
#init:
#	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#	go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest
#	go install github.com/micro/micro/v3/cmd/protoc-gen-openapi@latest
#
#.PHONY: api
#api:
#	protoc --openapi_out=. --proto_path=. proto/user.proto

.PHONY: proto
proto:
	#protoc --proto_path=. --micro_out=. --go_out=:. proto/user.proto
	docker run --rm -v $(PWD):$(PWD) -w $(PWD) protoc-gen-micro-v2:1.0 protoc -I ./ --go_out=./ --micro_out=./ ./proto/product/*.proto
	
.PHONY: build
build:
	#go build -o user *.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o product-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t product-service:latest .
