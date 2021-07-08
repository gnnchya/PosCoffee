package util

import "fmt"

var (
	UnknownErr = defineErr("UNKNOWN", "UNKNOWN")

	ConvertInputToDomainErr = defineErr("CONVERSION", "INPUT_TO_DOMAIN")

	ValidationCreateErr = defineErr("VALIDATION", "CREATE")
	ValidationUpdateErr = defineErr("VALIDATION", "UPDATE")

	RepoListErr   = defineErr("REPOSITORY", "LIST")
	RepoCreateErr = defineErr("REPOSITORY", "CREATE")
	RepoReadErr   = defineErr("REPOSITORY", "READ")
	RepoUpdateErr = defineErr("REPOSITORY", "UPDATE")
	RepoDeleteErr = defineErr("REPOSITORY", "DELETE")
)

type Error struct {
	Cause   error  `json:"cause"`
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
}

type ErrorFunc func(cause error) (err *Error)

func defineErr(code string, subCode string) (errFn ErrorFunc) {
	return func(cause error) (err *Error) {
		return &Error{
			Cause:   cause,
			Code:    code,
			SubCode: subCode,
		}
	}
}

func TypeOfErr(err error) (typeOf *Error) {
	if e0, ok := err.(*Error); ok {
		return e0
	}
	return &Error{Cause: err}
}

func (e *Error) Error() (msg string) {
	causeStr := ""
	if e.Cause != nil {
		causeStr = e.Cause.Error()
	}
	return fmt.Sprintf("%s (%s:%s)", causeStr, e.Code, e.SubCode)
}

func (e *Error) IsType(errFn ErrorFunc) (is bool) {
	test := errFn(nil)
	return e.Code == test.Code && e.SubCode == test.SubCode
}
