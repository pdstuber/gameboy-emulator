package types

type Instruction interface {
	Execute() error
}
