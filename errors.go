package jexp

import "fmt"

type ErrorCode string

const (
	ErrorCodeOther                      ErrorCode = "other"
	ErrorCodeInvalidOperandsCount       ErrorCode = "invalid_operands_count"
	ErrorCodePropertyNotFound           ErrorCode = "property_not_found"
	ErrorCodeInvalidPropertyType        ErrorCode = "invalid_property_type"
	ErrorCodeUnsupportedDataType        ErrorCode = "unsupported_data_type"
	ErrorCodeUnsupportedExpressionType  ErrorCode = "unsupported_expression_type"
	ErrorCodeUnexpectedExpressionType   ErrorCode = "unexpected_expression_type"
	ErrorCodeCantResolveToExpressonType ErrorCode = "cant_resolve_to_expression_type"
	ErrorCodeIncompatibleEqualOperands  ErrorCode = "incompatible_equal_operands"
)

type PropertyNotFoundError string

func (e PropertyNotFoundError) Error() string {
	return "property not found: " + string(e)
}

type UnexpectedExpressionTypeError string

func (e UnexpectedExpressionTypeError) Error() string {
	return "unexpected expression type, expected: " + string(e)
}

type UnsupportedDataTypeError string

func (e UnsupportedDataTypeError) Error() string {
	return "unsupported data type: " + string(e)
}

type UnsupportedExpressionTypeError string

func (e UnsupportedExpressionTypeError) Error() string {
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
