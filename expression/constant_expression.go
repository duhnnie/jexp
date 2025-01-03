package expression

type ConstantExpression[T Types] struct {
	value T
}

func NewConstant[T Types](value T) *ConstantExpression[T] {
	return &ConstantExpression[T]{value: value}
}

func (c *ConstantExpression[T]) Resolve(ctx ExpressionContext) (T, string, error) {
	return c.value, "", nil
}
