package expression

type OperationExpression struct {
	Type     string       `json:"type" required:"true"`
	Name     string       `json:"name" required:"true"`
	Operands []Expression `json:"operands" required:"true"`
}

func (exp *OperationExpression) UnmarshalJSON(data []byte) error {
	return unmarshalJSON(exp, data)
}
