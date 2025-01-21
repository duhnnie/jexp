package expression

import "fmt"

type EqualExpression[T Types] struct {
	Operands []Expression[T]
}

func NewEqual[T Types](operands ...Expression[T]) *EqualExpression[T] {
	return &EqualExpression[T]{Operands: operands}
}

func (e *EqualExpression[T]) Resolve(ctx ExpressionContext) (bool, string, error) {
	if len(e.Operands) < 2 {
		return false, "", &ErrorInvalidOperandsCount{minCount: 2, actualCount: len(e.Operands)}
	}

	var previous T

	for index, operand := range e.Operands {
		current, errorPath, err := operand.Resolve(ctx)

		if err != nil {
			return false, getPath(expTypeEqual, fmt.Sprintf("operands[%d]", index), errorPath), err
		}

		if index > 0 {
			if previous != current {
				return false, "", nil
			}
		}

		previous = current
	}

	return true, "", nil
}
