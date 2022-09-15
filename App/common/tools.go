package common

type Result[T any] struct {
	Value        T
	ErrorMessage string
}

func (result Result[T]) IsSuccess() bool {
	return result.ErrorMessage == ""
	// can also try to use reflect.ValueOf(&result.value).Elem().IsZero()
}

func GetSuccessResult[T any](value T) Result[T] {
	return Result[T]{value, ""}
}

func GetFailedResult[T any](errorMessage string) Result[T] {
	return Result[T]{*new(T), errorMessage}
}
