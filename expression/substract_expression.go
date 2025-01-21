package expression

type SubsctractExpression struct {
	Operands []Expression[float64]
}

func NewSubstract(operands ...Expression[float64]) *SubsctractExpression {
	return &SubsctractExpression{Operands: operands}
}

func (s *SubsctractExpression) Resolve(ctx ExpressionContext) (float64, string, error) {
	if len(s.Operands) != 2 {
		return 0, "", &ErrorInvalidOperandsCount{minCount: 2, maxCount: 2, actualCount: len(s.Operands)}
	}

	op1, errorPath, err := s.Operands[0].Resolve(ctx)

	if err != nil {
		return 0, getPath(exptypeSubstract, "operands[0]", errorPath), err
	}

	op2, errorPath, err := s.Operands[1].Resolve(ctx)

	if err != nil {
		return 0, getPath(expTypeEqual, "operands[1]", errorPath), err
	}

	return op1 - op2, "", nil
}
