package errors

import (
	"beemongo/constants"
)

var BusinessErrorName = "BusinessError"

type BusinessError struct {
	Msg     string
	Status  constants.ResponseStatus
	Content interface{}
}

func (e BusinessError) Error() string {
	return BusinessErrorName
}
