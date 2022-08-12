package codec

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	F1 string
	F2 int
}


func TestGob(t *testing.T) {
	should := assert.New(t)
	gobBytes, err := GobEncode(&TestStruct{F1: "testf1", F2: 11})
	if should.NoError(err) {
		fmt.Println(gobBytes)
	}
	obj := TestStruct{}
	err = GobDecode(gobBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj)
	}
}
