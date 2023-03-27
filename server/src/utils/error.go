package utils

type ApiError struct {
	Err    string
	Status int
}

type ApiSuccess struct {
	Success string
	Status  int
}

func (e ApiError) Error() string {
	return e.Err
}

func (s ApiSuccess) Error() string {
	return s.Success
}
