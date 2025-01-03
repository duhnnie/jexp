package expression

type ExpressionContext interface {
	Get(string) (interface{}, error)
}
