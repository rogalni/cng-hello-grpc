# cng-hello-backend

Showcase of gRPC in Go.

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

### Compile grpc
- In order to use protoc you need to set it up (https://grpc.io/docs/languages/go/quickstart/)
```
protoc --go_out=plugins=grpc:api\gen .\api\proto\chat.proto
``` 