package TryError

import (
	"github.com/yyyear/YY"
)

type Result[T any] struct {
	Value T
	Error error
}

func Try[T any](value T, err error) Result[T] {
	return Result[T]{Value: value, Error: err}
}
func TryError(err error) Result[bool] {
	return Result[bool]{Value: err == nil, Error: err}
}

func (r Result[T]) Do() T {
	if r.Error != nil {
		YY.ErrorLeve(1, "错误处理:", r.Error)
	}
	return r.Value
}

func (r Result[T]) Catch(f func(error)) {
	if r.Error != nil {
		f(r.Error)
	}
	return
}
