package slices

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrItIsNotSlice = Error("It is not slice")
)