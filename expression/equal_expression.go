package expression

import "strconv"

type EqualExpression[T supportedTypes] struct {
	Operands []Expression[T]
}

func NewEqual[T supportedTypes](operands ...Expression[T]) *EqualExpression[T] {
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
			return false, getPath(expTypeEqual, strconv.Itoa(index), errorPath), err
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
