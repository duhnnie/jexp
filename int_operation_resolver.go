package jexp

import (
	"math"

	"github.com/duhnnie/jexp/expression"
)

const (
	CLAMP        = "clamp"
	SUBSTRACTION = "substraction"
)

type intOperationResolver struct {
	ExpressionResolver *JExp
}

func (r *intOperationResolver) clamp(opValue, opMin, opMax expression.Expression) (int64, error) {
	var value, min, max int64

	if err := r.ExpressionResolver.ResolveToInt(opValue, &value); err != nil {
		return 0, err
	} else if err := r.ExpressionResolver.ResolveToInt(opMin, &min); err != nil {
		return 0, err
	} else if err := r.ExpressionResolver.ResolveToInt(opMax, &max); err != nil {
		return 0, err
	}

	clamped := math.Min(
		math.Max(
			float64(value),
			float64(min),
		),
		float64(max),
	)

	return int64(clamped), nil
}

func (r *intOperationResolver) substract(op1, op2 expression.Expression) (int64, error) {
	var v1, v2 int64

	if err := r.ExpressionResolver.ResolveToInt(op1, &v1); err != nil {
		return 0, err
	} else if err := r.ExpressionResolver.ResolveToInt(op2, &v2); err != nil {
		return 0, err
	}

	return v1 - v2, nil
}

func (r *intOperationResolver) Resolve(name string, operands []expression.Expression) (int64, error) {
	if name == CLAMP && len(operands) != 3 {
		return 0, &ErrorInvalidOperandsCount{name, 3, 3}
	} else if name == SUBSTRACTION && len(operands) != 2 {
		return 0, &ErrorInvalidOperandsCount{name, 2, 2}
	}

	switch name {
	case CLAMP:
		return r.clamp(operands[0], operands[1], operands[2])
	case SUBSTRACTION:
		return r.substract(operands[0], operands[1])
	default:
		return 0, ErrorUnknownOperationName(name)
	}
}
