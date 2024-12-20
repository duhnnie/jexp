package jexp

import (
	"reflect"

	"github.com/duhnnie/jexp/expression"
)

const (
	AND = "and"
	OR  = "or"
	EQ  = "eq"
	NOT = "not"
)

type booleanOperationResolver struct {
	ExpressionResolver *JExp
}

func (r *booleanOperationResolver) and(op1, op2 expression.Expression) (bool, error) {
	var b1, b2 bool

	if err := r.ExpressionResolver.ResolveToBoolean(op1, &b1); err != nil || !b1 {
		return b1, err
	} else if err := r.ExpressionResolver.ResolveToBoolean(op2, &b2); err != nil {
		return false, err
	}

	return b1 && b2, nil
}

func (r *booleanOperationResolver) or(op1, op2 expression.Expression) (bool, error) {
	var b1, b2 bool

	if err := r.ExpressionResolver.ResolveToBoolean(op1, &b1); err != nil || b1 {
		return b1, err
	} else if err := r.ExpressionResolver.ResolveToBoolean(op2, &b2); err != nil {
		return false, err
	}

	return b1 || b2, nil
}

func (r *booleanOperationResolver) eq(op1, op2 expression.Expression) (bool, error) {
	if v1, err := r.ExpressionResolver.Resolve(op1); err != nil {
		return false, err
	} else if v2, err := r.ExpressionResolver.Resolve(op2); err != nil {
		return false, err
	} else {
		v1t := reflect.TypeOf(v1)
		v2t := reflect.TypeOf(v2)

		if v1t.Name() != v2t.Name() {
			return false, nil
		}

		switch v1.(type) {
		case uint:
			return v1.(uint) == v2.(uint), nil
		case int:
			return v1.(int) == v2.(int), nil
		case uint8:
			return v1.(uint8) == v2.(uint8), nil
		case int8:
			return v1.(int8) == v2.(int8), nil
		case uint16:
			return v1.(uint16) == v2.(uint16), nil
		case int16:
			return v1.(int16) == v2.(int16), nil
		case uint32:
			return v1.(uint32) == v2.(uint32), nil
		case int32:
			return v1.(int32) == v2.(int32), nil
		case uint64:
			return v1.(uint64) == v2.(uint64), nil
		case int64:
			return v1.(int64) == v2.(int64), nil
		case string:
			return v1.(string) == v2.(string), nil
		case bool:
			return v1.(bool) == v2.(bool), nil
		default:
			return false, &ErrorUnsupportedTypeForComparison{v1}
		}
	}
}

func (r *booleanOperationResolver) not(op1 expression.Expression) (bool, error) {
	var v bool

	if err := r.ExpressionResolver.ResolveToBoolean(op1, &v); err != nil {
		return false, err
	}

	return !v, nil
}

func (r *booleanOperationResolver) Resolve(name string, operands []expression.Expression) (bool, error) {
	if name == NOT && len(operands) != 1 {
		return false, &ErrorInvalidOperandsCount{
			operationName: name,
			minCount:      1,
			maxCount:      1,
		}
	} else if name != NOT && len(operands) != 2 {
		return false, &ErrorInvalidOperandsCount{
			operationName: name,
			minCount:      2,
			maxCount:      2,
		}
	}

	switch name {
	case AND:
		return r.and(operands[0], operands[1])
	case OR:
		return r.or(operands[0], operands[1])
	case EQ:
		return r.eq(operands[0], operands[1])
	case NOT:
		return r.not(operands[0])
	default:
		return false, ErrorUnknownOperationName(name)
	}
}
