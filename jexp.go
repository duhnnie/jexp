package jexp

import (
	"encoding/json"
	"fmt"

	"github.com/duhnnie/jexp/expression"
)

func New[T expression.Types](jsonData []byte) (expression.Expression[T], *JExpError) {
	var dict map[string]interface{}

	err := json.Unmarshal(jsonData, &dict)

	if err != nil {
		return nil, NewJExpError(ErrorOther, err)
	}

	if parsed, err := parseDict(dict); err != nil {
		return nil, err
	} else if jexp, ok := parsed.(expression.Expression[T]); !ok {
		var x T
		return nil, NewJExpError(ErrorCantResolveToExpressonType, CantResolveToExpressionTypeError(fmt.Sprintf("%T", x)))
	} else {
		return jexp, nil
	}
}

func parseToOperandsArray(dict map[string]interface{}) ([]interface{}, *JExpError) {
	if intfc, exists := dict["operands"]; !exists {
		return nil, NewJExpError(ErrorPropertyNotFound, PropertyNotFoundError("operands"))
	} else if interfaceArr, ok := intfc.([]interface{}); !ok {
		return nil, NewJExpError(ErrorInvalidPropertyType, &InvalidPropertyTypeError{"operands", "array"})
	} else {
		return interfaceArr, nil
	}
}

func parseToExpressionArray[T expression.Types](arr []interface{}) ([]expression.Expression[T], *JExpError) {
	var operands []expression.Expression[T]

	for index, intfc := range arr {
		if operandDict, ok := intfc.(map[string]interface{}); !ok {
			// TODO: consider removing the property name in the error, since that could be retrieved from the errorPath
			return nil, NewJExpError(ErrorInvalidPropertyType, &InvalidPropertyTypeError{fmt.Sprintf("operands[%d]", index), "object"})
		} else {
			operand, err := parseDict(operandDict)

			if err != nil {
				return nil, err
			}

			if operandExpression, ok := operand.(expression.Expression[T]); !ok {
				var x T
				return nil, NewJExpError(
					ErrorUnexpectedExpressionType,
					UnexpectedExpressionTypeError(
						fmt.Sprintf("%T", x),
					))
			} else {
				operands = append(operands, operandExpression)
			}
		}
	}

	return operands, nil
}

func parseOperands[T expression.Types](dict map[string]interface{}) ([]expression.Expression[T], *JExpError) {
	interfaceArr, err := parseToOperandsArray(dict)

	if err != nil {
		return nil, err
	}

	return parseToExpressionArray[T](interfaceArr)
}

func parseSubstract(dict map[string]interface{}) (interface{}, *JExpError) {
	arr, err := parseToOperandsArray(dict)

	if err != nil {
		return nil, err
	}

	if expArr, err := parseToExpressionArray[float64](arr); err == nil {
		return expression.NewSubstract(expArr...), nil
	} else {
		return nil, err
	}
}

func parseVariable(dict map[string]interface{}) (interface{}, *JExpError) {
	if value, exists := dict["value"]; !exists {
		return nil, NewJExpError(ErrorPropertyNotFound, PropertyNotFoundError("value"))
	} else if dataType, exists := dict["dataType"]; !exists {
		return nil, NewJExpError(ErrorPropertyNotFound, PropertyNotFoundError("dataType"))
	} else {
		switch dataType {
		case "int":
			return expression.NewVariable[int64](value.(string)), nil
		case "float":
			return expression.NewVariable[float64](value.(string)), nil
		case "string":
			return expression.NewVariable[string](value.(string)), nil
		case "bool":
			return expression.NewVariable[bool](value.(string)), nil
		default:
			return nil, NewJExpError(ErrorUnsupportedDataType, UnsupportedDataType(dataType.(string)))
		}
	}
}

func parseNot(dict map[string]interface{}) (interface{}, *JExpError) {
	if intfc, exists := dict["expression"]; !exists {
		return nil, NewJExpError(ErrorPropertyNotFound, PropertyNotFoundError("expression"))
	} else if operandDict, ok := intfc.(map[string]interface{}); !ok {
		return nil, NewJExpError(ErrorInvalidPropertyType, &InvalidPropertyTypeError{"expression", "object"})
	} else if operandExpression, err := parseDict(operandDict); err != nil {
		return nil, err
	} else if booleanOperand, ok := operandExpression.(expression.Expression[bool]); !ok {
		return nil, NewJExpError(ErrorUnexpectedExpressionType, UnexpectedExpressionTypeError(fmt.Sprintf("%T", booleanOperand)))
	} else {
		return expression.NewNot(booleanOperand), nil
	}
}

func parseOr(dict map[string]interface{}) (*expression.OrExpression, *JExpError) {
	operandExpressions, err := parseOperands[bool](dict)

	if err != nil {
		return nil, err
	}

	return expression.NewOr(operandExpressions...), nil
}

func parseAnd(dict map[string]interface{}) (*expression.AndExpression, *JExpError) {
	operandExpressions, err := parseOperands[bool](dict)

	if err != nil {
		return nil, err
	}

	return expression.NewAnd(operandExpressions...), nil
}

func parseEqual(dict map[string]interface{}) (interface{}, *JExpError) {
	arr, err := parseToOperandsArray(dict)

	if err != nil {
		return nil, err
	}

	// TODO: Evauluate error type to determine if it makes sense to continue parsing to other types.
	if expArr, err := parseToExpressionArray[float64](arr); err == nil {
		return expression.NewEqual(expArr...), nil
	} else if expArr, err := parseToExpressionArray[bool](arr); err == nil {
		return expression.NewEqual(expArr...), nil
	} else if expArr, err := parseToExpressionArray[string](arr); err == nil {
		return expression.NewEqual(expArr...), nil
	} else if expArr, err := parseToExpressionArray[int64](arr); err == nil {
		return expression.NewEqual(expArr...), nil
	} else {
		return nil, NewJExpError(ErrorIncompatibleEqualOperands, nil)
	}
}

func parseConstant(dict map[string]interface{}) (interface{}, *JExpError) {
	if value, exists := dict["value"]; !exists {
		return nil, NewJExpError(ErrorPropertyNotFound, PropertyNotFoundError("value"))
	} else if dataType, exists := dict["dataType"]; !exists {
		switch v := value.(type) {
		case float64:
			return expression.NewConstant[float64](v), nil
		case string:
			return expression.NewConstant[string](v), nil
		case bool:
			return expression.NewConstant[bool](v), nil
		default:
			return nil, NewJExpError(ErrorPropertyNotFound, PropertyNotFoundError("dataType"))
		}
	} else if dataType, ok := dataType.(string); !ok {
		return nil, NewJExpError(ErrorInvalidPropertyType, &InvalidPropertyTypeError{"dataType", "string"})
	} else {
		switch dataType {
		case "int", "float":
			if floatValue, ok := value.(float64); !ok {
				return nil, NewJExpError(ErrorCantResolveToExpressonType, CantResolveToExpressionTypeError(dataType))
			} else if dataType == "int" {
				return expression.NewConstant[int64](int64(floatValue)), nil
			} else {
				return expression.NewConstant[float64](floatValue), nil
			}
		case "string":
			return expression.NewConstant[string](value.(string)), nil
		case "bool":
			return expression.NewConstant[bool](value.(bool)), nil
		default:
			return nil, NewJExpError(ErrorUnsupportedDataType, UnsupportedDataType(dataType))
		}
	}
}

func parseDict(dict map[string]interface{}) (interface{}, *JExpError) {
	if t, exists := dict["type"]; !exists {
		return nil, NewJExpError(ErrorPropertyNotFound, PropertyNotFoundError("type"))
	} else {
		switch t {
		case "and":
			return parseAnd(dict)
		case "or":
			return parseOr(dict)
		case "not":
			return parseNot(dict)
		case "eq":
			return parseEqual(dict)
		case "subs":
			return parseSubstract(dict)
		case "var":
			return parseVariable(dict)
		case "const":
			return parseConstant(dict)
		default:
			return nil, NewJExpError(ErrorUnsupportedExpressionType, UnsupportedExpressionType(t.(string)))
		}
	}
}
