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
		message = "A required field is missing"
	}
	return Apperror{id: 2, message: message, field: field}
}

func NewInvalidInputError(message string, err error, field string) Apperror {
	if len(message) == 0 {
		message = "A field is invalid"
	}
	return Apperror{id: 2, message: message, sys_message: err.Error(), field: field}
}

func (err *Apperror) getId() int {
	return err.id
}

func (err *Apperror) getIdString() string {
	return strconv.Itoa(err.id)
}

func (err *Apperror) getMessage() string {
	return err.message
}

func (err *Apperror) getSysMesage() string {
	return err.sys_message
}

func (err *Apperror) getfield() string {
	return err.field
}
