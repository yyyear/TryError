package TryError

import (
	"github.com/yyyear/YY"
	"google.golang.org/protobuf/proto"
)

// Encode proto Model 转 []byte
func Encode[T proto.Message](m T) []byte {
	b, err := proto.Marshal(m)
	if err != nil {
		YY.ErrorLeve(1, "转换数据错误:", err)
	}
	return b
}

// Decode 转换 proto Model
func Decode[T proto.Message](bytes []byte, m T) Result[T] {
	err := proto.Unmarshal(bytes, m)
	if err != nil {
		return Try(m, err)
	}
	return Try(m, nil)
}
