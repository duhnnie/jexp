package expression

import "strconv"

type CompareExpression[T comparable] struct {
	expType  string
	fn       compareFn[T]
	Operands []Expression[T]
}

func NewCompare[T comparable](expType string, operands ...Expression[T]) (*CompareExpression[T], error) {
	fn := getCompareFn[T](expType)

	if fn == nil {
		return nil, ErrorUnknownCompareType(expType)
	}

	return &CompareExpression[T]{expType: expType, fn: fn, Operands: operands}, nil
}

func (e *CompareExpression[T]) GetType(ctx ExpressionContext) string {
	return e.expType
}

func (e *CompareExpression[T]) Resolve(ctx ExpressionContext) (bool, string, error) {
	if len(e.Operands) < 2 {
		return false, "", &ErrorInvalidOperandsCount{minCount: 2, actualCount: len(e.Operands)}
	}

	var prev T

	for index, operand := range e.Operands {
		value, errorPath, err := operand.Resolve(ctx)

		if err != nil {
			return false, getPath(string(e.expType), strconv.Itoa(index), errorPath), err
		}

		if index > 0 && !e.fn(prev, value) {
			return false, "", nil
		}

		prev = value
	}

	return true, "", nil
}
