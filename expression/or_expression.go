package expression

import (
	"strconv"

	"github.com/duhnnie/godash"
)

type OrExpression struct {
	Operands []Expression[bool]
}

func NewOr(operands ...Expression[bool]) *OrExpression {
	return &OrExpression{Operands: operands}
}

func (o *OrExpression) Resolve(ctx ExpressionContext) (res bool, errorPath string, err error) {
	var lastProcessedPath string = ""

	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
				errorPath = lastProcessedPath
			} else {
				panic(r)
			}
		}
	}()

	if len(o.Operands) < 2 {
		return false, "", &ErrorInvalidOperandsCount{minCount: 2, actualCount: len(o.Operands)}
	}

	var r bool = godash.Reduce(o.Operands, func(acc bool, operand Expression[bool], index int, _ []Expression[bool]) bool {
		r, errorPath, err := operand.Resolve(ctx)

		if err != nil {
			lastProcessedPath = getPath(expTypeEqual, strconv.Itoa(index), errorPath)
			panic(err)
		}

		return acc || r
	}, false)

	return r, "", nil
}
