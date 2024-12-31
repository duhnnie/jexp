package expression

type Types interface {
	float64 | string | bool | int64
}

const (
	expTypeAnd       = "and"
	expTypeConst     = "const"
	expTypeEqual     = "eq"
	expTypeNot       = "not"
	expTypeOr        = "or"
	exptypeSubstract = "subst"
	expTypeVar       = "var"
	expTypeClamp     = "clamp"
)
