package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type CPU6502 struct {
	Bus   *Bus      // Bus connection pointer
	flags FLAGS6502 // 6502 Flags

	// Registers
	status uint8  // Status
	a      uint8  // Accumulator
	x      uint8  // X
	y      uint8  // Y
	stkp   uint8  // Stack pointer
	pc     uint16 // Program counter

	// Data
	fetched uint8

	// Addressing
	addr_abs uint16
	addr_rel uint16
	opcode   uint8
	cycles   uint8

	tmp uint16

	// Instruction lookup table
	lookup Instructions
}

// Setup the CPU
func (c *CPU6502) Initialize() {
	c.flags = FLAGS6502{}
	c.flags.Initialize()
	c.InitInternals()
}

func (c *CPU6502) InitInternals() {
	// Init Registers
	c.status = 0x00
	c.a = 0x00
	c.x = 0x00
	c.y = 0x00
	c.stkp = 0x00
	c.pc = 0x0000

	// Init Data
	c.fetched = 0x00
	c.addr_abs = 0x0000
	c.addr_rel = 0x0000
	c.opcode = 0x00
	c.cycles = 0

	c.lookup = Instructions{}
	c.lookup = *c.lookup.Fill(c)
}

// Read from the bus
func (c *CPU6502) Read(addr uint16) uint8 {
	return c.Bus.Read(addr, false)
}

// Write to the bus
func (c *CPU6502) Write(addr uint16, data uint8) {
	c.Bus.Write(addr, data)
}

func (c *CPU6502) Clock() {
	if c.cycles == 0 {
		c.opcode = c.Read(c.pc)
		c.pc++
		addrname := runtime.FuncForPC(reflect.ValueOf(c.lookup[c.opcode].AddrMode).Pointer()).Name()
		addrname = addrname[len(addrname)-6 : len(addrname)-3]
		fmt.Println(c.lookup[c.opcode].name, addrname, "PC:", c.pc)
		c.cycles = c.lookup[c.opcode].cycles
		addrCycle := c.lookup[c.opcode].AddrMode()
		opCycle := c.lookup[c.opcode].OpCode()

		c.cycles += (addrCycle & opCycle)

		c.SetFlag(c.flags.U, true)
	}
	c.cycles--
}

func (c *CPU6502) Fetch() uint8 {
	if runtime.FuncForPC(reflect.ValueOf(c.lookup[c.opcode].AddrMode).Pointer()).Name() != runtime.FuncForPC(reflect.ValueOf(c.IMP).Pointer()).Name() {
		c.fetched = c.Read(c.addr_abs)
	}
	return c.fetched
}

func (c *CPU6502) GetFlag(f uint8) uint8 {
	if (c.status & f) > 0 {
		return 1
	} else {
		return 0
	}
}

func (c *CPU6502) SetFlag(f uint8, v bool) {
	if v {
		c.status |= f
	} else {
		c.status &= ^f
	}
}

func (c *CPU6502) PrintFlags() {
	fmt.Println("B:", c.GetFlag(c.flags.B))
	fmt.Println("C:", c.GetFlag(c.flags.C))
	fmt.Println("D:", c.GetFlag(c.flags.D))
	fmt.Println("I:", c.GetFlag(c.flags.I))
	fmt.Println("N:", c.GetFlag(c.flags.N))
	fmt.Println("U:", c.GetFlag(c.flags.U))
	fmt.Println("V:", c.GetFlag(c.flags.V))
	fmt.Println("Z:", c.GetFlag(c.flags.Z))
}

func (c *CPU6502) PrintRegisters() {
	fmt.Println("ACC:", c.a)
	//fmt.Printf(" - %b\n", c.a)
	fmt.Println("X:", c.x)
	fmt.Println("Y:", c.y)
	fmt.Println("FETCH:", c.fetched)
}

func (c *CPU6502) Reset() {
	c.a = 0
	c.x = 0
	c.y = 0
	c.stkp = 0xFD
	c.status = 0x00 | c.GetFlag(c.flags.U)

	c.addr_abs = 0xFFFC
	var lo uint16 = uint16(c.Read(c.addr_abs + 0))
	var hi uint16 = uint16(c.Read(c.addr_abs + 1))

	c.pc = (hi << 8) | lo
	c.addr_rel = 0x0000
	c.addr_abs = 0x0000
	c.fetched = 0x00

	c.cycles = 8
}

func (c *CPU6502) IRQ() {
	if c.GetFlag(c.flags.I) == 0 {
		c.Write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
		c.stkp--
		c.Write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
		c.stkp--

		c.SetFlag(c.flags.B, false)
		c.SetFlag(c.flags.U, true)
		c.SetFlag(c.flags.I, true)
		c.Write(0x0100+uint16(c.stkp), c.status)
		c.stkp--

		c.addr_abs = 0xFFFE
		var lo uint16 = uint16(c.Read(c.addr_abs + 0))
		var hi uint16 = uint16(c.Read(c.addr_abs + 1))
		c.pc = (hi >> 8) | lo

		c.cycles = 7
	}
}

func (c *CPU6502) NMI() {
	c.Write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
	c.stkp--
	c.Write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
	c.stkp--

	c.SetFlag(c.flags.B, false)
	c.SetFlag(c.flags.U, true)
	c.SetFlag(c.flags.I, true)
	c.Write(0x0100+uint16(c.stkp), c.status)
	c.stkp--

	c.addr_abs = 0xFFFE
	var lo uint16 = uint16(c.Read(c.addr_abs + 0))
	var hi uint16 = uint16(c.Read(c.addr_abs + 1))
	c.pc = (hi >> 8) | lo

	c.cycles = 8
}