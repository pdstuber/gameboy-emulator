package types

type CPU interface {
	ReadMemoryAndIncrementProgramCounter() byte
	SetProgramCounter(address Address)
	SetRegisterA(value uint16)
	SetRegisterBC(value uint16)
	SetRegisterDE(value uint16)
	SetRegisterHL(value uint16)
	SetRegisterSP(value uint16)

	GetRegisterA() uint16
	GetRegisterB() uint16
	GetRegisterC() uint16
	GetRegisterD() uint16
	GetRegisterE() uint16
	GetRegisterH() uint16
	GetRegisterHL() uint16
	GetRegisterL() uint16

	SetFlagZ(bool)
	SetFlagN(bool)
	SetFlagH(bool)
	SetFlagC(bool)
}

type Instruction interface {
	Execute(cpu CPU) (int, error)
}
