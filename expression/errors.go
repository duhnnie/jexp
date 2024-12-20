package expression

import (
	"fmt"
	"reflect"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type ErrorUnknownExpressionType string

func (e ErrorUnknownExpressionType) Error() string {
	return fmt.Sprintf("unknown expression type \"%s\"", string(e))
}

type ErrorMissingRequiredProperty string

func (m ErrorMissingRequiredProperty) Error() string {
	return fmt.Sprintf("missing required property: %s", string(m))
}

type ErrorNotSupportedKind reflect.Kind

func (e ErrorNotSupportedKind) Error() string {
	return fmt.Sprintf("expression unmarshalling doesnt support type %d", e)
}

const (
	ErrorNoExpressionTypeFound  = Error("no \"type\" property found for operation expression")
	ErrorInvalidDataTypeForType = Error("invalid data type for \"type\" property")
)
