package utils

import "github.com/LeonLow97/inventory-management-system-golang-react-postgresql/api/handlers"

type ApiError struct {
	Err    string
	Status int
}

type ApiSuccess struct {
	Success string
	Status  int
}

type ApiGetSuccess struct {
	Success  string
	Status   int
	Result []handlers.User
}

func (e ApiError) Error() string {
	return e.Err
}

func (s ApiSuccess) Error() string {
	return s.Success
}

func (s ApiGetSuccess) Error() string {
	return s.Success
}
