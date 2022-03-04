# cng-hello-backend

Showcase of gRPC in Go.

#

## 1. How to run:

### Local
```
go run .\cmd\server.go

go run .\cmd\client.go

```

### docker-compose
```
docker build -f build/package/docker/client/Dockerfile -t cng-hello-grpc-client .
docker build -f build/package/docker/server/Dockerfile -t cng-hello-grpc-server .

cd test/docker/cng-hello-backend-standalone

docker-compose up
```

### Required
- Setup protoc (https://grpc.io/docs/languages/go/quickstart/)
```
go install github.com/bufbuild/buf/cmd/buf@v1.0.0
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
``` 
### Generate proto
```
cd /api
buf generate
```

#

## 2. To be done
### gRPC
 - Add client side streaming 
 - Add server side streaming 
 - Add bidirectional streaming 