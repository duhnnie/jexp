package expression

import "math"

type ClampExpression struct {
	Value Expression[float64]
	Min   Expression[float64]
	Max   Expression[float64]
}

func NewClamp(value, min, max Expression[float64]) *ClampExpression {
	return &ClampExpression{Value: value, Min: min, Max: max}
}

func (c *ClampExpression) Resolve(ctx ExpressionContext) (float64, string, error) {
	value, errorPath, err := c.Value.Resolve(ctx)

	if err != nil {
		return 0, getPath(expTypeClamp, "operands[0]", errorPath), err
	}

	min, errorPath, err := c.Min.Resolve(ctx)

	if err != nil {
		return 0, getPath(expTypeClamp, "operands[1]", errorPath), err
	}

	max, errorPath, err := c.Max.Resolve(ctx)

	if err != nil {
		return 0, getPath(expTypeClamp, "operands[2]", errorPath), err
	}

	return math.Min(
		math.Max(value, min),
		max,
	), "", nil
}
