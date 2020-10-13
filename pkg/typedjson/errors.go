package typedjson

import "fmt"

type TypeNotDefinedError struct{}

func (e *TypeNotDefinedError) Error() string {
	return "type tag is not defined"
}

type InvalidJsonTypeError struct {
	Expected, Actual string
}

func (e *InvalidJsonTypeError) Error() string {
	return fmt.Sprintf("invalid JSON type: expected: \"%s\", actual \"%s\"", e.Expected, e.Actual)
}
