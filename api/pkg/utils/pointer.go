package utils

func ToPointer[T interface{}](v T) *T {
	return &v
}
