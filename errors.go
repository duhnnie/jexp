package jexp

import "fmt"

type ErrorCode string

const (
	ErrorOther                      ErrorCode = "other"
	ErrorPropertyNotFound           ErrorCode = "property_not_found"
	ErrorInvalidPropertyType        ErrorCode = "invalid_property_type"
	ErrorUnsupportedDataType        ErrorCode = "unsupported_data_type"
	ErrorUnsupportedExpressionType  ErrorCode = "unsupported_expression_type"
	ErrorUnexpectedExpressionType   ErrorCode = "unexpected_expression_type"
	ErrorCantResolveToExpressonType ErrorCode = "cant_resolve_to_expression_type"
	ErrorIncompatibleEqualOperands  ErrorCode = "incompatible_equal_operands"
)

type PropertyNotFoundError string

func (e PropertyNotFoundError) Error() string {
	return "property not found: " + string(e)
}

type UnexpectedExpressionTypeError string

func (e UnexpectedExpressionTypeError) Error() string {
	return "unexpected expression type, expected: " + string(e)
}

type UnsupportedDataType string

func (e UnsupportedDataType) Error() string {
	return "unsupported data type: " + string(e)
}

type UnsupportedExpressionType string

func (e UnsupportedExpressionType) Error() string {
	return "unsupported expression type: " + string(e)
}

type InvalidPropertyTypeError string

func (e InvalidPropertyTypeError) Error() string {
	return fmt.Sprintf("invalid property type, expected: %s", string(e))
}

type CantResolveToExpressionTypeError string

func (e CantResolveToExpressionTypeError) Error() string {
	return "can't resolve to expression type " + string(e)
}

type JExpError struct {
	Code ErrorCode
	Err  error
}

func NewJExpError(code ErrorCode, err error) *JExpError {
	return &JExpError{Code: code, Err: err}
}

func (e *JExpError) Error() string {
	if e.Err == nil {
		return string(e.Code)
	}

	return fmt.Sprintf("%s", e.Err)
}
