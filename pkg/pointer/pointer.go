package pointer

func Pointer[T any](val T) *T {
	return &val
}

func Value[T any](ptr *T) T {
	if ptr == nil {
		var z T

		return z
	}

	return *ptr
}
