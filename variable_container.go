package jexp

type VariableContainer interface {
	Set(name string, data []byte) error
	GetFloat64(variableName string) (float64, error)
	GetBool(variableName string) (bool, error)
	GetString(variableName string) (string, error)
}
