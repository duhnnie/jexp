package expression

type VariableExpression struct {
	Type string `json:"type" required:"true"`
	Name string `json:"name" required:"true"`
}

func (exp *VariableExpression) UnmarshalJSON(data []byte) error {
	return unmarshalJSON(exp, data)
}
