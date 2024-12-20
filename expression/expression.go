package expression

type supportedTypes interface {
	float64 | string | bool
}

type Expression[T supportedTypes] interface {
	Resolve(ctx ExpressionContext) (T, string, error)
}
