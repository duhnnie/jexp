package expression

type ConstantExpression[T supportedTypes] struct {
	value T
}

func NewConstant[T supportedTypes](value T) *ConstantExpression[T] {
	return &ConstantExpression[T]{value: value}
}

func (c *ConstantExpression[T]) Resolve(ctx ExpressionContext) (T, string, error) {
	return c.value, "", nil
}
