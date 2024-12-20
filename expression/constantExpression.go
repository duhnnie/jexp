package expression

type constantExpressionValueType interface {
	int | float64 | bool
}

type ConstantExpression[U constantExpressionValueType] struct {
	Type  string `json:"type" required:"true"`
	Value U      `json:"value" required:"true"`
}

func (exp *ConstantExpression[U]) UnmarshalJSON(data []byte) error {
	return unmarshalJSON(exp, data)
}
