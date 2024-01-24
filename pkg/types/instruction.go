package types

type CPU interface {
	ReadMemoryAndIncrementProgramCounter() byte
	ReadMemory(address Address) byte
	WriteMemory(address Address, data byte)
	SetProgramCounter(value uint16)
	GetProgramCounter() uint16
	SetRegisterA(value uint8)
	SetRegisterB(value uint8)
	SetRegisterC(value uint8)
	SetRegisterD(value uint8)
	SetRegisterH(value uint8)
	SetRegisterE(value uint8)
	SetRegisterL(value uint8)
	SetRegisterBC(value uint16)
	SetRegisterDE(value uint16)
	SetRegisterHL(value uint16)
	SetRegisterSP(value uint16)

	GetRegisterA() uint8
	GetRegisterB() uint8
	GetRegisterC() uint8
	GetRegisterD() uint8
	GetRegisterE() uint8
	GetRegisterH() uint8
	GetRegisterHL() uint16
	GetRegisterL() uint8
	GetRegisterSP() uint16
	GetRegisterDE() uint16
	GetRegisterBC() uint16
	GetRegisterAF() uint16
	UnsetFlagZero()

	UnsetFlagHalfCarry()

	UnsetFlagCarry()
	UnsetFlagSubtraction()
	SetFlagZero()

	SetFlagCarry()

	SetFlagHalfCarry()

	SetFlagSubtraction()
	GetFlagZero() bool

	GetFlagCarry() bool

	GetFlagHalfCarry() bool
	GetFlagSubtraction() bool
}

type Instruction interface {
	Execute(cpu CPU) (int, error)
}
