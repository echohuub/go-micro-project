```bash
docker run --rm -v $(PWD):$(PWD) -w $(PWD) micro/micro new user
```

```bash
docker run --rm -v $(PWD):$(PWD) -w $(PWD) protoc-gen-micro-v2:1.0 protoc -I ./ --go_out=./ --micro_out=./ ./proto/category/*.proto
```


```bash
docker run -d --name jaeger -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one
```

