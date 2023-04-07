package cpu6502

type Instruction struct {
	index    int
	name     string
	OpCode   func() uint8
	AddrMode func() uint8
	cycles   uint8
}

type Instructions []Instruction

func (l *Instructions) Fill(c *CPU6502) *Instructions {
	return &Instructions{
		Instruction{0, "BRK", c.BRK, c.IMM, 7},
		Instruction{1, "ORA", c.ORA, c.IZX, 6},
		Instruction{2, "???", c.NIM, c.IMP, 2},
		Instruction{3, "SLO", c.SLO, c.IZX, 8},
		Instruction{4, "???", c.NOP, c.ZP0, 3},
		Instruction{5, "ORA", c.ORA, c.ZP0, 3},
		Instruction{6, "ASL", c.ASL, c.ZP0, 5},
		Instruction{7, "SLO", c.SLO, c.ZP0, 5},
		Instruction{8, "PHP", c.PHP, c.IMP, 3},
		Instruction{9, "ORA", c.ORA, c.IMM, 2},
		Instruction{10, "ASL", c.ASL, c.IMP, 2},
		Instruction{11, "???", c.NIM, c.IMP, 2},
		Instruction{12, "???", c.NOP, c.ABS, 4},
		Instruction{13, "ORA", c.ORA, c.ABS, 4},
		Instruction{14, "ASL", c.ASL, c.ABS, 6},
		Instruction{15, "SLO", c.SLO, c.ABS, 6},
		Instruction{16, "BPL", c.BPL, c.REL, 2},
		Instruction{17, "ORA", c.ORA, c.IZY, 5},
		Instruction{18, "???", c.NIM, c.IMP, 2},
		Instruction{19, "SLO", c.SLO, c.IZY, 8},
		Instruction{20, "???", c.NOP, c.ZPX, 4},
		Instruction{21, "ORA", c.ORA, c.ZPX, 4},
		Instruction{22, "ASL", c.ASL, c.ZPX, 6},
		Instruction{23, "SLO", c.SLO, c.ZPX, 6},
		Instruction{24, "CLC", c.CLC, c.IMP, 2},
		Instruction{25, "ORA", c.ORA, c.ABY, 4},
		Instruction{26, "???", c.NOP, c.IMP, 2},
		Instruction{27, "SLO", c.SLO, c.ABY, 7},
		Instruction{28, "???", c.NOP, c.ABX, 4},
		Instruction{29, "ORA", c.ORA, c.ABX, 4},
		Instruction{30, "ASL", c.ASL, c.ABX, 7},
		Instruction{31, "SLO", c.SLO, c.ABX, 7},
		Instruction{32, "JSR", c.JSR, c.ABS, 6},
		Instruction{33, "AND", c.AND, c.IZX, 6},
		Instruction{34, "???", c.NIM, c.IMP, 2},
		Instruction{35, "RLA", c.RLA, c.IZX, 8},
		Instruction{36, "BIT", c.BIT, c.ZP0, 3},
		Instruction{37, "AND", c.AND, c.ZP0, 3},
		Instruction{38, "ROL", c.ROL, c.ZP0, 5},
		Instruction{39, "RLA", c.RLA, c.ZP0, 5},
		Instruction{40, "PLP", c.PLP, c.IMP, 4},
		Instruction{41, "AND", c.AND, c.IMM, 2},
		Instruction{42, "ROL", c.ROL, c.IMP, 2},
		Instruction{43, "???", c.NIM, c.IMP, 2},
		Instruction{44, "BIT", c.BIT, c.ABS, 4},
		Instruction{45, "AND", c.AND, c.ABS, 4},
		Instruction{46, "ROL", c.ROL, c.ABS, 6},
		Instruction{47, "RLA", c.RLA, c.ABS, 6},
		Instruction{48, "BMI", c.BMI, c.REL, 2},
		Instruction{49, "AND", c.AND, c.IZY, 5},
		Instruction{50, "???", c.NIM, c.IMP, 2},
		Instruction{51, "RLA", c.RLA, c.IZY, 8},
		Instruction{52, "???", c.NOP, c.ZPX, 4},
		Instruction{53, "AND", c.AND, c.ZPX, 4},
		Instruction{54, "ROL", c.ROL, c.ZPX, 6},
		Instruction{55, "RLA", c.RLA, c.ZPX, 6},
		Instruction{56, "SEC", c.SEC, c.IMP, 2},
		Instruction{57, "AND", c.AND, c.ABY, 4},
		Instruction{58, "???", c.NOP, c.IMP, 2},
		Instruction{59, "RLA", c.RLA, c.ABY, 7},
		Instruction{60, "???", c.NOP, c.ABX, 4},
		Instruction{61, "AND", c.AND, c.ABX, 4},
		Instruction{62, "ROL", c.ROL, c.ABX, 7},
		Instruction{63, "RLA", c.RLA, c.ABX, 7},
		Instruction{64, "RTI", c.RTI, c.IMP, 6},
		Instruction{65, "EOR", c.EOR, c.IZX, 6},
		Instruction{66, "???", c.NIM, c.IMP, 2},
		Instruction{67, "SRE", c.SRE, c.IZX, 8},
		Instruction{68, "???", c.NOP, c.ZP0, 3},
		Instruction{69, "EOR", c.EOR, c.ZP0, 3},
		Instruction{70, "LSR", c.LSR, c.ZP0, 5},
		Instruction{71, "SRE", c.SRE, c.ZP0, 5},
		Instruction{72, "PHA", c.PHA, c.IMP, 3},
		Instruction{73, "EOR", c.EOR, c.IMM, 2},
		Instruction{74, "LSR", c.LSR, c.IMP, 2},
		Instruction{75, "ALR", c.ALR, c.IMM, 2},
		Instruction{76, "JMP", c.JMP, c.ABS, 3},
		Instruction{77, "EOR", c.EOR, c.ABS, 4},
		Instruction{78, "LSR", c.LSR, c.ABS, 6},
		Instruction{79, "SRE", c.SRE, c.ABS, 6},
		Instruction{80, "BVC", c.BVC, c.REL, 2},
		Instruction{81, "EOR", c.EOR, c.IZY, 5},
		Instruction{82, "???", c.NIM, c.IMP, 2},
		Instruction{83, "SRE", c.SRE, c.IZY, 8},
		Instruction{84, "???", c.NOP, c.ZPX, 4},
		Instruction{85, "EOR", c.EOR, c.ZPX, 4},
		Instruction{86, "LSR", c.LSR, c.ZPX, 6},
		Instruction{87, "SRE", c.SRE, c.ZPX, 6},
		Instruction{88, "CLI", c.CLI, c.IMP, 2},
		Instruction{89, "EOR", c.EOR, c.ABY, 4},
		Instruction{90, "???", c.NOP, c.IMP, 2},
		Instruction{91, "SRE", c.SRE, c.ABY, 7},
		Instruction{92, "???", c.NOP, c.ABX, 4},
		Instruction{93, "EOR", c.EOR, c.ABX, 4},
		Instruction{94, "LSR", c.LSR, c.ABX, 7},
		Instruction{95, "SRE", c.SRE, c.ABX, 7},
		Instruction{96, "RTS", c.RTS, c.IMP, 6},
		Instruction{97, "ADC", c.ADC, c.IZX, 6},
		Instruction{98, "???", c.NIM, c.IMP, 2},
		Instruction{99, "RRA", c.RRA, c.IZX, 8},
		Instruction{100, "???", c.NOP, c.ZP0, 3},
		Instruction{101, "ADC", c.ADC, c.ZP0, 3},
		Instruction{102, "ROR", c.ROR, c.ZP0, 5},
		Instruction{103, "RRA", c.RRA, c.ZP0, 5},
		Instruction{104, "PLA", c.PLA, c.IMP, 4},
		Instruction{105, "ADC", c.ADC, c.IMM, 2},
		Instruction{106, "ROR", c.ROR, c.IMP, 2},
		Instruction{107, "???", c.NIM, c.IMP, 2},
		Instruction{108, "JMP", c.JMP, c.IND, 5},
		Instruction{109, "ADC", c.ADC, c.ABS, 4},
		Instruction{110, "ROR", c.ROR, c.ABS, 6},
		Instruction{111, "RRA", c.RRA, c.ABS, 6},
		Instruction{112, "BVS", c.BVS, c.REL, 2},
		Instruction{113, "ADC", c.ADC, c.IZY, 5},
		Instruction{114, "???", c.NIM, c.IMP, 2},
		Instruction{115, "RRA", c.RRA, c.IZY, 8},
		Instruction{116, "???", c.NOP, c.ZPX, 4},
		Instruction{117, "ADC", c.ADC, c.ZPX, 4},
		Instruction{118, "ROR", c.ROR, c.ZPX, 6},
		Instruction{119, "RRA", c.RRA, c.ZPX, 6},
		Instruction{120, "SEI", c.SEI, c.IMP, 2},
		Instruction{121, "ADC", c.ADC, c.ABY, 4},
		Instruction{122, "???", c.NOP, c.IMP, 2},
		Instruction{123, "RRA", c.RRA, c.ABY, 7},
		Instruction{124, "???", c.NOP, c.ABX, 4},
		Instruction{125, "ADC", c.ADC, c.ABX, 4},
		Instruction{126, "ROR", c.ROR, c.ABX, 7},
		Instruction{127, "RRA", c.RRA, c.ABX, 7},
		Instruction{128, "???", c.NOP, c.IMM, 2},
		Instruction{129, "STA", c.STA, c.IZX, 6},
		Instruction{130, "???", c.NOP, c.IMM, 2},
		Instruction{131, "SAX", c.SAX, c.IZX, 6},
		Instruction{132, "STY", c.STY, c.ZP0, 3},
		Instruction{133, "STA", c.STA, c.ZP0, 3},
		Instruction{134, "STX", c.STX, c.ZP0, 3},
		Instruction{135, "SAX", c.SAX, c.ZP0, 3},
		Instruction{136, "DEY", c.DEY, c.IMP, 2},
		Instruction{137, "???", c.NOP, c.IMP, 2},
		Instruction{138, "TXA", c.TXA, c.IMP, 2},
		Instruction{139, "???", c.NIM, c.IMP, 2},
		Instruction{140, "STY", c.STY, c.ABS, 4},
		Instruction{141, "STA", c.STA, c.ABS, 4},
		Instruction{142, "STX", c.STX, c.ABS, 4},
		Instruction{143, "SAX", c.SAX, c.ABS, 4},
		Instruction{144, "BCC", c.BCC, c.REL, 2},
		Instruction{145, "STA", c.STA, c.IZY, 6},
		Instruction{146, "???", c.NIM, c.IMP, 2},
		Instruction{147, "???", c.NIM, c.IMP, 6},
		Instruction{148, "STY", c.STY, c.ZPX, 4},
		Instruction{149, "STA", c.STA, c.ZPX, 4},
		Instruction{150, "STX", c.STX, c.ZPY, 4},
		Instruction{151, "SAX", c.SAX, c.ZPY, 4},
		Instruction{152, "TYA", c.TYA, c.IMP, 2},
		Instruction{153, "STA", c.STA, c.ABY, 5},
		Instruction{154, "TXS", c.TXS, c.IMP, 2},
		Instruction{155, "???", c.NIM, c.IMP, 5},
		Instruction{156, "???", c.NOP, c.IMP, 5},
		Instruction{157, "STA", c.STA, c.ABX, 5},
		Instruction{158, "???", c.NIM, c.IMP, 5},
		Instruction{159, "???", c.NIM, c.IMP, 5},
		Instruction{160, "LDY", c.LDY, c.IMM, 2},
		Instruction{161, "LDA", c.LDA, c.IZX, 6},
		Instruction{162, "LDX", c.LDX, c.IMM, 2},
		Instruction{163, "LAX", c.LAX, c.IZX, 6},
		Instruction{164, "LDY", c.LDY, c.ZP0, 3},
		Instruction{165, "LDA", c.LDA, c.ZP0, 3},
		Instruction{166, "LDX", c.LDX, c.ZP0, 3},
		Instruction{167, "LAX", c.LAX, c.ZP0, 3},
		Instruction{168, "TAY", c.TAY, c.IMP, 2},
		Instruction{169, "LDA", c.LDA, c.IMM, 2},
		Instruction{170, "TAX", c.TAX, c.IMP, 2},
		Instruction{171, "???", c.NIM, c.IMP, 2},
		Instruction{172, "LDY", c.LDY, c.ABS, 4},
		Instruction{173, "LDA", c.LDA, c.ABS, 4},
		Instruction{174, "LDX", c.LDX, c.ABS, 4},
		Instruction{175, "LAX", c.LAX, c.ABS, 4},
		Instruction{176, "BCS", c.BCS, c.REL, 2},
		Instruction{177, "LDA", c.LDA, c.IZY, 5},
		Instruction{178, "???", c.NIM, c.IMP, 2},
		Instruction{179, "LAX", c.LAX, c.IZY, 5},
		Instruction{180, "LDY", c.LDY, c.ZPX, 4},
		Instruction{181, "LDA", c.LDA, c.ZPX, 4},
		Instruction{182, "LDX", c.LDX, c.ZPY, 4},
		Instruction{183, "LAX", c.LAX, c.ZPY, 4},
		Instruction{184, "CLV", c.CLV, c.IMP, 2},
		Instruction{185, "LDA", c.LDA, c.ABY, 4},
		Instruction{186, "TSX", c.TSX, c.IMP, 2},
		Instruction{187, "???", c.NIM, c.IMP, 4},
		Instruction{188, "LDY", c.LDY, c.ABX, 4},
		Instruction{189, "LDA", c.LDA, c.ABX, 4},
		Instruction{190, "LDX", c.LDX, c.ABY, 4},
		Instruction{191, "LAX", c.LAX, c.ABY, 4},
		Instruction{192, "CPY", c.CPY, c.IMM, 2},
		Instruction{193, "CMP", c.CMP, c.IZX, 6},
		Instruction{194, "???", c.NOP, c.IMM, 2},
		Instruction{195, "DCP", c.DCP, c.IZX, 8},
		Instruction{196, "CPY", c.CPY, c.ZP0, 3},
		Instruction{197, "CMP", c.CMP, c.ZP0, 3},
		Instruction{198, "DEC", c.DEC, c.ZP0, 5},
		Instruction{199, "DCP", c.DCP, c.ZP0, 5},
		Instruction{200, "INY", c.INY, c.IMP, 2},
		Instruction{201, "CMP", c.CMP, c.IMM, 2},
		Instruction{202, "DEX", c.DEX, c.IMP, 2},
		Instruction{203, "???", c.NIM, c.IMP, 2},
		Instruction{204, "CPY", c.CPY, c.ABS, 4},
		Instruction{205, "CMP", c.CMP, c.ABS, 4},
		Instruction{206, "DEC", c.DEC, c.ABS, 6},
		Instruction{207, "DCP", c.DCP, c.ABS, 6},
		Instruction{208, "BNE", c.BNE, c.REL, 2},
		Instruction{209, "CMP", c.CMP, c.IZY, 5},
		Instruction{210, "???", c.NIM, c.IMP, 2},
		Instruction{211, "DCP", c.DCP, c.IZY, 8},
		Instruction{212, "???", c.NOP, c.ZPX, 4},
		Instruction{213, "CMP", c.CMP, c.ZPX, 4},
		Instruction{214, "DEC", c.DEC, c.ZPX, 6},
		Instruction{215, "DCP", c.DCP, c.ZPX, 6},
		Instruction{216, "CLD", c.CLD, c.IMP, 2},
		Instruction{217, "CMP", c.CMP, c.ABY, 4},
		Instruction{218, "NOP", c.NOP, c.IMP, 2},
		Instruction{219, "DCP", c.DCP, c.ABY, 7},
		Instruction{220, "???", c.NOP, c.ABX, 4},
		Instruction{221, "CMP", c.CMP, c.ABX, 4},
		Instruction{222, "DEC", c.DEC, c.ABX, 7},
		Instruction{223, "DCP", c.DCP, c.ABX, 7},
		Instruction{224, "CPX", c.CPX, c.IMM, 2},
		Instruction{225, "SBC", c.SBC, c.IZX, 6},
		Instruction{226, "???", c.NOP, c.IMM, 2},
		Instruction{227, "ISC", c.ISC, c.IZX, 8},
		Instruction{228, "CPX", c.CPX, c.ZP0, 3},
		Instruction{229, "SBC", c.SBC, c.ZP0, 3},
		Instruction{230, "INC", c.INC, c.ZP0, 5},
		Instruction{231, "ISC", c.ISC, c.ZP0, 5},
		Instruction{232, "INX", c.INX, c.IMP, 2},
		Instruction{233, "SBC", c.SBC, c.IMM, 2},
		Instruction{234, "NOP", c.NOP, c.IMP, 2},
		Instruction{235, "SBC", c.SBC, c.IMM, 2},
		Instruction{236, "CPX", c.CPX, c.ABS, 4},
		Instruction{237, "SBC", c.SBC, c.ABS, 4},
		Instruction{238, "INC", c.INC, c.ABS, 6},
		Instruction{239, "ISC", c.ISC, c.ABS, 6},
		Instruction{240, "BEQ", c.BEQ, c.REL, 2},
		Instruction{241, "SBC", c.SBC, c.IZY, 5},
		Instruction{242, "???", c.NIM, c.IMP, 2},
		Instruction{243, "ISC", c.ISC, c.IZY, 8},
		Instruction{244, "???", c.NOP, c.ZPX, 4},
		Instruction{245, "SBC", c.SBC, c.ZPX, 4},
		Instruction{246, "INC", c.INC, c.ZPX, 6},
		Instruction{247, "ISC", c.ISC, c.ZPX, 6},
		Instruction{248, "SED", c.SED, c.IMP, 2},
		Instruction{249, "SBC", c.SBC, c.ABY, 4},
		Instruction{250, "NOP", c.NOP, c.IMP, 2},
		Instruction{251, "ISC", c.ISC, c.ABY, 7},
		Instruction{252, "???", c.NOP, c.ABX, 4},
		Instruction{253, "SBC", c.SBC, c.ABX, 4},
		Instruction{254, "INC", c.INC, c.ABX, 7},
		Instruction{255, "ISC", c.ISC, c.ABX, 7}}
}
