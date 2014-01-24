package apperror

import (
	"strconv"
)

type Apperror struct {
	id          int
	message     string
	sys_message string
	field       string
}

func NewDBError(message string, err error) Apperror {
	if len(message) == 0 {
		message = "Some error occured while querying database. Please try later."
	}
	return Apperror{id: 1, message: message, sys_message: err.Error()}
}

func NewRequiredError(message string, field string) Apperror {
	if len(message) == 0 {
		message = "Field required"
	}
	return Apperror{id: 2, message: message, field: field}
}

func NewInvalidInputError(message string, err error, field string) Apperror {
	if len(message) == 0 {
		message = "Invalid field"
	}
	return Apperror{id: 3, message: message, sys_message: err.Error(), field: field}
}

func NewResourceMissingError(message string) Apperror {
	if len(message) == 0 {
		message = "Resource missing"
	}
	return Apperror{id: 4, message: message}
}

func (err *Apperror) GetId() int {
	return err.id
}

func (err *Apperror) GetIdString() string {
	return strconv.Itoa(err.id)
}

func (err *Apperror) GetMessage() string {
	return err.message
}

func (err *Apperror) GetSysMesage() string {
	return err.sys_message
}

func (err *Apperror) GetField() string {
	return err.field
}
