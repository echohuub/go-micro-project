# User Service

```bash
# 生成 micro go 文件
docker run --rm -it -v $(PWD):$(PWD) -w $(PWD) protoc-gen-micro-v2:1.0
protoc -I ./ --go_out=./ --micro_out=./ ./*.proto
```

```bash
# 编译 Go 可执行产物
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user *.go
# 编译 Docker 镜像
docker build -t user:latest .
```