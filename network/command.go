package network

type Command interface {
	IsMalformed() bool
}
