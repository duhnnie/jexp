package expression

import (
	"fmt"
)

// type Error string

// func (e Error) Error() string {
// 	return string(e)
// }

// type ErrorUnknownExpressionType string

// func (e ErrorUnknownExpressionType) Error() string {
// 	return fmt.Sprintf("unknown expression type \"%s\"", string(e))
// }

// type ErrorMissingRequiredProperty string

// func (m ErrorMissingRequiredProperty) Error() string {
// 	return fmt.Sprintf("missing required property: %s", string(m))
// }

// type ErrorNotSupportedKind reflect.Kind

// func (e ErrorNotSupportedKind) Error() string {
// 	return fmt.Sprintf("expression unmarshalling doesnt support type %d", e)
// }

// const (
// 	ErrorNoExpressionTypeFound  = Error("no \"type\" property found for operation expression")
// 	ErrorInvalidDataTypeForType = Error("invalid data type for \"type\" property")
// )

// ----

type ErrorInvalidOperandsCount struct {
	minCount    uint
	maxCount    uint
	actualCount int
}

func (e *ErrorInvalidOperandsCount) Error() string {
	switch {
	case e.minCount == e.maxCount:
		return fmt.Sprintf("operands count should be equal to %d, got %d", e.minCount, e.actualCount)
	case e.minCount == 0:
		return fmt.Sprintf("operands count should at most %d, got %d", e.maxCount, e.actualCount)
	case e.maxCount == 0:
		return fmt.Sprintf("operands count should at least %d, got %d", e.minCount, e.actualCount)
	case e.minCount > e.maxCount:
		return fmt.Sprintf("operands count should be between %d and %d, got %d", e.maxCount, e.minCount, e.actualCount)
	default:
		return fmt.Sprintf("operands count should be between %d and %d, got %d", e.minCount, e.maxCount, e.actualCount)
	}
}

type ErrorCantResolveToType struct {
	expectedTypeValue interface{}
	actualTypeValue   interface{}
}

func (e *ErrorCantResolveToType) Error() string {
	return fmt.Sprintf("can't be resolved to type \"%T\", got type \"%T\"", e.expectedTypeValue, e.actualTypeValue)
}
