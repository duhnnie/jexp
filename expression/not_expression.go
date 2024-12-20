package expression

import "strconv"

type NotExpression struct {
	Operands Expression[bool]
}

func NewNot(operands Expression[bool]) *NotExpression {
	return &NotExpression{Operands: operands}
}

func (n *NotExpression) Resolve(ctx ExpressionContext) (bool, string, error) {
	value, errorPath, err := n.Operands.Resolve(ctx)

	if err != nil {
		return false, getPath(expTypeNot, strconv.Itoa(0), errorPath), err
	}

	return !value, "", nil
}
