package jexp

import (
	"encoding/json"
	"fmt"

	"github.com/duhnnie/jexp/expression"
)

func New[T expression.Types](jsonData []byte) (expression.Expression[T], string, *JExpError) {
	var dict map[string]interface{}

	err := json.Unmarshal(jsonData, &dict)

	if err != nil {
		return nil, "", NewJExpError(ErrorCodeOther, err)
	}

	if parsed, pathErr, err := parseDict(dict); err != nil {
		return nil, pathErr, err
	} else if jexp, ok := parsed.(expression.Expression[T]); !ok {
		var x T
		return nil, "[root]", NewJExpError(ErrorCodeCantResolveToExpressonType, CantResolveToExpressionTypeError(fmt.Sprintf("%T", x)))
	} else {
		return jexp, "", nil
	}
}

func parseToOperandsArray(dict map[string]interface{}) ([]interface{}, string, *JExpError) {
	if intfc, exists := dict["operands"]; !exists {
		return nil, "", NewJExpError(ErrorCodePropertyNotFound, PropertyNotFoundError("operands"))
	} else if interfaceArr, ok := intfc.([]interface{}); !ok {
		return nil, ".operands", NewJExpError(ErrorCodeInvalidPropertyType, InvalidPropertyTypeError("array"))
	} else {
		return interfaceArr, "", nil
	}
}

func parseToExpressionArray(arr []interface{}) ([]interface{}, string, *JExpError) {
	var iExpArray []interface{}

	for index, intfc := range arr {
		if operandDict, ok := intfc.(map[string]interface{}); !ok {
			// TODO: consider removing the property name in the error, since that could be retrieved from the errorPath
			return nil, fmt.Sprintf("[%d]", index), NewJExpError(ErrorCodeInvalidPropertyType, InvalidPropertyTypeError("object"))
		} else {
			operand, errPath, err := parseDict(operandDict)

			if err != nil {
				return nil, fmt.Sprintf("[%d]%s", index, errPath), err
			}

			iExpArray = append(iExpArray, operand)
		}
	}

	return iExpArray, "", nil
}

func parseToExpressionGenericArray[T expression.Types](arr []interface{}) ([]expression.Expression[T], string, *JExpError) {
	var expressions []expression.Expression[T]

	for index, intfc := range arr {
		if expressionItem, ok := intfc.(expression.Expression[T]); !ok {
			var x T

			return nil, fmt.Sprintf("[%d]", index), NewJExpError(
				ErrorCodeUnexpectedExpressionType,
				UnexpectedExpressionTypeError(
					fmt.Sprintf("%T", x),
				))
		} else {
			expressions = append(expressions, expressionItem)
		}
	}

	return expressions, "", nil
}

func parseOperands[T expression.Types](dict map[string]interface{}) ([]expression.Expression[T], string, *JExpError) {
	interfaceArr, errPath, err := parseToOperandsArray(dict)

	if err != nil {
		return nil, errPath, err
	}

	intfcArray, errPath, err := parseToExpressionArray(interfaceArr)

	if err != nil {
		return nil, ".operands" + errPath, err
	}

	expressions, errPath, err := parseToExpressionGenericArray[T](intfcArray)

	if err != nil {
		return nil, ".operands" + errPath, err
	}

	return expressions, "", nil
}

func parseSubstract(dict map[string]interface{}) (interface{}, string, *JExpError) {
	arr, errPath, err := parseToOperandsArray(dict)

	if err != nil {
		return nil, errPath, err
	}

	intfcArray, errPath, err := parseToExpressionArray(arr)

	if err != nil {
		return nil, ".operands" + errPath, err
	}

	if expArr, _, err := parseToExpressionGenericArray[float64](intfcArray); err == nil {
		return expression.NewSubstract(expArr...), "", nil
	} else {
		return nil, ".operands", NewJExpError(ErrorCodeIncompatibleEqualOperands, nil)
	}
}

func parseVariable(dict map[string]interface{}) (interface{}, string, *JExpError) {
	if value, exists := dict["value"]; !exists {
		return nil, "", NewJExpError(ErrorCodePropertyNotFound, PropertyNotFoundError("value"))
	} else if dataType, exists := dict["dataType"]; !exists {
		return nil, "", NewJExpError(ErrorCodePropertyNotFound, PropertyNotFoundError("dataType"))
	} else {
		switch dataType {
		case "number":
			return expression.NewVariable[float64](value.(string)), "", nil
		case "string":
			return expression.NewVariable[string](value.(string)), "", nil
		case "boolean":
			return expression.NewVariable[bool](value.(string)), "", nil
		default:
			return nil, ".dataType", NewJExpError(ErrorCodeUnsupportedDataType, UnsupportedDataTypeError(dataType.(string)))
		}
	}
}

func parseNot(dict map[string]interface{}) (interface{}, string, *JExpError) {
	if intfc, exists := dict["expression"]; !exists {
		return nil, "", NewJExpError(ErrorCodePropertyNotFound, PropertyNotFoundError("expression"))
	} else if expressionMap, ok := intfc.(map[string]interface{}); !ok {
		return nil, ".expression", NewJExpError(ErrorCodeInvalidPropertyType, InvalidPropertyTypeError("object"))
	} else if operandExpression, errPath, err := parseDict(expressionMap); err != nil {
		return nil, ".expression" + errPath, err
	} else if booleanOperand, ok := operandExpression.(expression.Expression[bool]); !ok {
		return nil, ".expression", NewJExpError(ErrorCodeUnexpectedExpressionType, UnexpectedExpressionTypeError(fmt.Sprintf("%T", booleanOperand)))
	} else {
		return expression.NewNot(booleanOperand), "", nil
	}
}

func parseOr(dict map[string]interface{}) (*expression.OrExpression, string, *JExpError) {
	operandExpressions, errPath, err := parseOperands[bool](dict)

	if err != nil {
		return nil, errPath, err
	}

	return expression.NewOr(operandExpressions...), "", nil
}

func parseAnd(dict map[string]interface{}) (*expression.AndExpression, string, *JExpError) {
	operandExpressions, errPath, err := parseOperands[bool](dict)

	if err != nil {
		return nil, errPath, err
	}

	return expression.NewAnd(operandExpressions...), "", nil
}

func parseEqual(dict map[string]interface{}) (interface{}, string, *JExpError) {
	arr, errPath, err := parseToOperandsArray(dict)

	if err != nil {
		return nil, errPath, err
	}

	iExpArray, errPath, err := parseToExpressionArray(arr)

	if err != nil {
		return nil, ".operands" + errPath, err
	}

	// At this point, every expression in array is a valid one.
	// But now we need to try to type assert them into the same type.
	// If this fails, we return an error notifying about incompatible types.
	if expArr, _, err := parseToExpressionGenericArray[float64](iExpArray); err == nil {
		return expression.NewEqual(expArr...), "", nil
	} else if expArr, _, err := parseToExpressionGenericArray[bool](iExpArray); err == nil {
		return expression.NewEqual(expArr...), "", nil
	} else if expArr, _, err := parseToExpressionGenericArray[string](iExpArray); err == nil {
		return expression.NewEqual(expArr...), "", nil
	} else {
		return nil, ".operands", NewJExpError(ErrorCodeIncompatibleEqualOperands, nil)
	}
}

func parseConstant(dict map[string]interface{}) (interface{}, string, *JExpError) {
	if value, exists := dict["value"]; !exists {
		return nil, ".value", NewJExpError(ErrorCodePropertyNotFound, PropertyNotFoundError("value"))
	} else if dataType, exists := dict["dataType"]; !exists {
		switch v := value.(type) {
		case float64:
			return expression.NewConstant(v), "", nil
		case string:
			return expression.NewConstant(v), "", nil
		case bool:
			return expression.NewConstant(v), "", nil
		default:
			return nil, "", NewJExpError(ErrorCodePropertyNotFound, PropertyNotFoundError("dataType"))
		}
	} else if dataType, ok := dataType.(string); !ok {
		return nil, ".dataType", NewJExpError(ErrorCodeInvalidPropertyType, InvalidPropertyTypeError("string"))
	} else {
		switch dataType {
		case "number":
			if floatValue, ok := value.(float64); !ok {
				return nil, ".value", NewJExpError(ErrorCodeCantResolveToExpressonType, CantResolveToExpressionTypeError(dataType))
			} else {
				return expression.NewConstant(floatValue), "", nil
			}
		case "string":
			return expression.NewConstant(value.(string)), "", nil
		case "boolean":
			return expression.NewConstant(value.(bool)), "", nil
		default:
			return nil, ".dataType", NewJExpError(ErrorCodeUnsupportedDataType, UnsupportedDataTypeError(dataType))
		}
	}
}

func parseClamp(dict map[string]interface{}) (interface{}, string, *JExpError) {
	arr, errPath, err := parseToOperandsArray(dict)

	if err != nil {
		return nil, errPath, err
	}

	intfcArray, errPath, err := parseToExpressionArray(arr)

	if err != nil {
		return nil, ".operands" + errPath, err
	}

	if len(intfcArray) != 3 {
		return nil, ".operands", NewJExpError(ErrorCodeInvalidOperandsCount, expression.NewInvalidOperandsCountError(3, 3, len(intfcArray)))
	}

	if expressions, errPath, err := parseToExpressionGenericArray[float64](intfcArray); err != nil {
		return nil, ".operands" + errPath, err
	} else {
		return expression.NewClamp(expressions[0], expressions[1], expressions[2]), "", nil
	}
}

func parseDict(dict map[string]interface{}) (interface{}, string, *JExpError) {
	if t, exists := dict["type"]; !exists {
		return nil, "", NewJExpError(ErrorCodePropertyNotFound, PropertyNotFoundError("type"))
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
		case "clamp":
			return parseClamp(dict)
		default:
			return nil, "", NewJExpError(ErrorCodeUnsupportedExpressionType, UnsupportedExpressionTypeError(t.(string)))
		}
	}
}
