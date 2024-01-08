package types

type CPU interface {
	ReadMemoryAndIncrementProgramCounter() byte
	SetProgramCounter(address Address)
}

type Instruction interface {
	Execute(cpu CPU) error
}
