package expression

type Types interface {
	string | bool | float64
}

const (
	expTypeAnd       = "and"
	expTypeConst     = "const"
	expTypeEqual     = "eq"
	expTypeNot       = "not"
	expTypeOr        = "or"
	exptypeSubstract = "subs"
	expTypeVar       = "var"
	expTypeClamp     = "clamp"
)
