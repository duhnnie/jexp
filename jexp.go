package jexp

import (
	"reflect"

	"github.com/duhnnie/jexp/expression"
)

type JExp struct {
	variables         VariableContainer
	booleanOpResolver *booleanOperationResolver
	intOpResolver     *intOperationResolver
	intConstResolver  *IntConstantResolver
}

func New(variables VariableContainer) *JExp {
	booleanOpResolver := booleanOperationResolver{}
	intOpResolver := intOperationResolver{}
	intConstResolver := IntConstantResolver{}

	this := JExp{
		variables:         variables,
		booleanOpResolver: &booleanOpResolver,
		intOpResolver:     &intOpResolver,
		intConstResolver:  &intConstResolver,
	}

	booleanOpResolver.ExpressionResolver = &this
	intOpResolver.ExpressionResolver = &this

	return &this
}

func resolve(j *JExp, v any) (interface{}, error) {
	switch exp := v.(type) {
	case uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64, string, bool:
		return v, nil
	case *expression.OperationExpression:
		return j.resolveOperationExpression(v.(*expression.OperationExpression))
	case *expression.VariableExpression:
		return j.resolveVariableExpression(v.(*expression.VariableExpression))
	case *expression.ConstantExpression[int]:
		return j.intConstResolver.Resolve(exp)
	default:
		return nil, ErrorCantResolveToExpression
	}
}

func resolveToType[T any](j *JExp, e expression.Expression, out *T) error {
	if res, err := resolve(j, e); err != nil {
		return err
	} else if v, ok := res.(T); ok {
		*out = v
		return nil
	} else {
		return ErrorCantResolveToType(reflect.ValueOf(out).Elem().Type().Name())
	}
}

func (j *JExp) resolveOperationExpression(exp *expression.OperationExpression) (interface{}, error) {
	switch exp.Type {
	case expression.ExpTypeBooleanOperation:
		return j.booleanOpResolver.Resolve(exp.Name, exp.Operands)
	case expression.ExpTypeIntOperation:
		return j.intOpResolver.Resolve(exp.Name, exp.Operands)
	default:
		return nil, expression.ErrorUnknownExpressionType(exp.Type)
	}
}

func (j *JExp) resolveVariableExpression(exp *expression.VariableExpression) (interface{}, error) {
	switch exp.Type {
	case expression.ExpTypeIntVariable:
		if v, err := j.variables.GetFloat64(exp.Name); err != nil {
			return 0, err
		} else {
			return int64(v), nil
		}
	default:
		return nil, expression.ErrorUnknownExpressionType(exp.Type)
	}
}

func (j *JExp) Resolve(v any) (interface{}, error) {
	return resolve(j, v)
}

func (j *JExp) ResolveToBoolean(e expression.Expression, out *bool) error {
	return resolveToType(j, e, out)
}

func (j *JExp) ResolveToInt(e expression.Expression, out *int64) error {
	return resolveToType(j, e, out)
}

func (j *JExp) ResolveToUInt(e expression.Expression, out *uint64) error {
	return resolveToType(j, e, out)
}
