package expression

import (
	"strconv"
)

type AndExpression struct {
	Operands []Expression[bool]
}

func NewAnd(operands ...Expression[bool]) *AndExpression {
	return &AndExpression{Operands: operands}
}

func (a *AndExpression) Resolve(ctx ExpressionContext) (bool, string, error) {
	if len(a.Operands) < 2 {
		return false, "", &ErrorInvalidOperandsCount{minCount: 2, actualCount: len(a.Operands)}
	}

	var r bool = true

	for index, operand := range a.Operands {
		value, errorPath, err := operand.Resolve(ctx)

		if err != nil {
			return false, getPath(expTypeAnd, strconv.Itoa(index), errorPath), err
		}

		r = r && value

		if !r {
			return false, "", nil
		}
	}

	return r, "", nil
}
