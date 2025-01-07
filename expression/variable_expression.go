package expression

type VariableExpression[T Types] struct {
	Name string
}

func NewVariable[T Types](name string) *VariableExpression[T] {
	return &VariableExpression[T]{Name: name}
}

func (v *VariableExpression[T]) Resolve(ctx ExpressionContext) (T, string, error) {
	var value T
	var ok bool

	rawValue, err := ctx.Get(v.Name)

	if err != nil {
		return value, "", err
	} else if value, ok = rawValue.(T); ok {
		return value, "", nil
	} else {
		return value, getPath(expTypeVar, v.Name, ""), &ErrorCantResolveToType{expectedTypeValue: value, actualTypeValue: rawValue}
	}
}
