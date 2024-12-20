package expression

import (
	"encoding/json"
	"reflect"
)

var typeOfOperationExpression = reflect.TypeOf(OperationExpression{})
var typeOfVariableExpresion = reflect.TypeOf(VariableExpression{})
var typeOfIntConstantExpression = reflect.TypeOf(ConstantExpression[int]{})

var typeMap = map[string]reflect.Type{
	ExpTypeBooleanOperation: typeOfOperationExpression,
	ExpTypeIntOperation:     typeOfOperationExpression,
	ExpTypeIntVariable:      typeOfVariableExpresion,
	ExpTypeIntConstant:      typeOfIntConstantExpression,
}

func UnmarshalExpression(data []byte) (Expression, error) {
	m := map[string]interface{}{}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	var T reflect.Type
	var exp Expression

	if expressionTypeI, exists := m["type"]; !exists {
		return nil, ErrorNoExpressionTypeFound
	} else if expressionType, ok := expressionTypeI.(string); !ok {
		return nil, ErrorInvalidDataTypeForType
	} else if T, exists = typeMap[expressionType]; !exists {
		return nil, ErrorUnknownExpressionType(expressionType)
	}

	exp = reflect.New(T).Interface().(Expression)
	err := json.Unmarshal(data, &exp)

	return exp, err
}
