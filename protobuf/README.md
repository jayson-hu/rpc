
```
protoc  --go_out=./pb --go_opt=module="github.com/jayson-hu/rpc/protobuf/pb" --go_grpc_out=./pb pb/hello.proto
protoc --go_out=./pb --go_opt=module="github.com/jayson-hu/rpc/protobuf/pb"   ./hello.proto

protoc -I.  --go_opt=module="github.com/jayson-hu/rpc/protobuf/pb"  --go_out=plugins=grpc:. ./hello.proto
protoc -I.  --go_opt=module="github.com/jayson-hu/rpc/protobuf/pb"  --go_out=plugins=grpc:. ./hello.proto
```