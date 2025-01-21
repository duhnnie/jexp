package expression

type comparable interface {
	string | float64
}

type compareFn[T comparable] func(a, b T) bool

func gt[T comparable](a, b T) bool {
	return a > b
}

func lt[T comparable](a, b T) bool {
	return a < b
}

func gte[T comparable](a, b T) bool {
	return a >= b
}

func lte[T comparable](a, b T) bool {
	return a <= b
}

func getCompareFn[T comparable](expType string) compareFn[T] {
	switch expType {
	case "gt":
		return gt
	case "lt":
		return lt
	case "gte":
		return gte
	case "lte":
		return lte
	default:
		return nil
	}
}
