# gRPC and Protocol buffers Golang example

- Local modules
```bash
go get google.golang.org/protobuf/cmd/protoc-gen-go
```

```bash
go get google.golang.org/grpc
```

- Global modules
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

- gRPC Client
```bash
go install github.com/ktr0731/evans@latest
```

- Compile *.proto files
```bash
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```