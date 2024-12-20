package expression

import (
	"encoding/json"
	"reflect"
)

type Expression interface {
	// GetType() string
}

// TODO: Find a way to incorporate this directly on the struct
func unmarshalJSON(exp Expression, data []byte) error {
	var m map[string]json.RawMessage

	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}

	valuePtr := reflect.ValueOf(exp)
	valueElem := valuePtr.Elem()
	typePtr := reflect.TypeOf(exp)
	typeElem := typePtr.Elem()

	for i := 0; i < typeElem.NumField(); i += 1 {
		structField := typeElem.Field(i)
		jsonField := structField.Tag.Get("json")
		isRequired := structField.Tag.Get("required") == "true"
		fieldValue := valueElem.Field(i)

		if jsonData, exists := m[jsonField]; !exists && isRequired {
			return ErrorMissingRequiredProperty(jsonField)
		} else if !exists {
			continue
		} else {
			switch structField.Type.Kind() {
			case reflect.Float64, reflect.Int, reflect.String, reflect.Bool:
				var jsonValue interface{}

				if err := json.Unmarshal(jsonData, &jsonValue); err != nil {
					return err
				}

				if structField.Type.Kind() == reflect.Int && reflect.TypeOf(jsonValue).Kind() == reflect.Float64 {
					fieldValue.SetInt(int64(jsonValue.(float64)))
				} else {
					fieldValue.Set(reflect.ValueOf(jsonValue))
				}
			case reflect.Slice:
				var rawMessageString []json.RawMessage

				if err := json.Unmarshal(jsonData, &rawMessageString); err != nil {
					return err
				}

				var operands []Expression

				for _, expressionData := range rawMessageString {
					exp, err := UnmarshalExpression(expressionData)

					if err != nil {
						return err
					}

					operands = append(operands, exp)
				}

				fieldValue.Set(reflect.ValueOf(operands))
			default:
				return ErrorNotSupportedKind(structField.Type.Kind())
			}
		}
	}

	return nil
}
