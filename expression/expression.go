package expression

type Expression[T Types] interface {
	Resolve(ctx ExpressionContext) (T, string, error)
}
