# grpc hello-world

```
protoc -I.  --go_opt=module="github.com/jayson-hu/rpc/grpc/simple/server/pb"  --go_out=plugins=grpc:. hello.proto
```