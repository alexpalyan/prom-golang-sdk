package utils

func Ptr[T any](v T) *T {
	return &v
}

func Value[T any](v *T) T {
	if v == nil {
		var r T
		return r
	}

	return *v
}
