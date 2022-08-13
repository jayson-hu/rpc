package pb_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"testing"

	"github.com/jayson-hu/rpc/protobuf/pb"
)

func TestMarshal(t *testing.T) {
	should := assert.New(t)
	str := &pb.String{Value: "Hello"}
	//obj ---> 序列化[]byte数组
	pbBytes, err := proto.Marshal(str)
	if should.NoError(err) {
		fmt.Println(pbBytes)
	}
	obj := pb.String{}
	err = proto.Unmarshal(pbBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj.Value)
	}
}
