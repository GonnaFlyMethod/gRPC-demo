# gRPC-demo

## Generating go files from .proto files
```bash
protoc --go_out=plugins=grpc:chat chat.proto
```

## Run the demo
prepare 2 opening tabs in your terminal and type:
```bash
go run cmd/server/main.go
```
and then
```bash
go run cmd/client/main.go
```
