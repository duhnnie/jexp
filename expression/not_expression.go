package expression

import "strconv"

type NotExpression struct {
	Operand Expression[bool]
}

func NewNot(operand Expression[bool]) *NotExpression {
	return &NotExpression{Operand: operand}
}

func (n *NotExpression) Resolve(ctx ExpressionContext) (bool, string, error) {
	value, errorPath, err := n.Operand.Resolve(ctx)

	if err != nil {
		return false, getPath(expTypeNot, strconv.Itoa(0), errorPath), err
	}

	return !value, "", nil
}
