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

	// we use "value" instead of "name" because that's the name used in the JSON schema
	if err != nil {
		return value, "value", err
	} else if value, ok = rawValue.(T); ok {
		return value, "", nil
	} else {
		return value, getPath(expTypeVar, "value", ""), &ErrorCantResolveToType{expectedTypeValue: value, actualTypeValue: rawValue}
	}
}
