package codec

import (
	"bytes"
	"encoding/gob"
)

//object ---> gob  []byte
func GobEncode(obj interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	//编码后的结果输出到 buf里面
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(obj); err != nil {
		return nil, err
	}
	//返回buf中bytes数组
	return buf.Bytes(), nil
}

// []byte ---->gob   ----> object

func GobDecode(data []byte, val interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	return decoder.Decode(val)
}

//func main() {
//
//}
