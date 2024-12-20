package jexp

import "fmt"

type Error string

func (e Error) Error() string {
	return string(e)
}

const ErrorCantResolveToExpression = Error("can't resolve to a known expression")

type ErrorUnknownOperationName string

func (e ErrorUnknownOperationName) Error() string {
	return fmt.Sprintf("unknown operation name: %s", string(e))
}

type ErrorResolvedToUnexpectedType struct {
	expectedType string
	actualType   interface{}
}

func (e *ErrorResolvedToUnexpectedType) Error() string {
	return fmt.Sprintf("expression resolved to unexpected type (%T), expecting: %s", e.actualType, e.expectedType)
}

type ErrorInvalidOperandsCount struct {
	operationName string
	minCount      uint
	maxCount      uint
}

func (e *ErrorInvalidOperandsCount) Error() string {
	switch {
	case e.minCount == e.maxCount:
		return fmt.Sprintf("operands count for operation \"%s\" should be equal to %d\n", e.operationName, e.minCount)
	case e.minCount == 0:
		return fmt.Sprintf("operands count for operation \"%s\" should at most %d\n", e.operationName, e.maxCount)
	case e.maxCount == 0:
		return fmt.Sprintf("operands count for operation \"%s\" should at least %d\n", e.operationName, e.minCount)
	case e.minCount > e.maxCount:
		return fmt.Sprintf("operands count for operation \"%s\" should be between %d and %d\n", e.operationName, e.maxCount, e.minCount)
	default:
		return fmt.Sprintf("operands count for operation \"%s\" should be between %d and %d\n", e.operationName, e.minCount, e.maxCount)
	}
}

type ErrorUnsupportedTypeForComparison struct {
	unsupportedType any
}

func (e *ErrorUnsupportedTypeForComparison) Error() string {
	return fmt.Sprintf("unsupported type for comparison: %T", e.unsupportedType)
}

type ErrorCantResolveToType string

func (e ErrorCantResolveToType) Error() string {
	return fmt.Sprintf("can't resolve to type: %s", string(e))
}
