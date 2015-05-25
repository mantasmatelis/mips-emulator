package main

func (m *Machine) add(inst uint32) {
	s, t, d := parseStd(inst)
	result := m.GetReg(t) + m.GetReg(s)
	m.logStd("add", s, t, d, result)
	m.SetReg(d, result)
}

func (m *Machine) sub(inst uint32) {
	s, t, d := parseStd(inst)
	result := m.GetReg(s) - m.GetReg(t)
	m.logStd("sub", s, t, d, result)
	m.SetReg(d, result)
}

func (m *Machine) mult(inst uint32) {
	s, t := parseSt(inst)
	result := uint64(int64(m.GetReg(s)) * int64(m.GetReg(t)))
	lo, hi := uint32(result), uint32(result>>32)
	m.logSt("mult", s, t, lo, hi)
	m.SetLoHi(lo, hi)
}

func (m *Machine) multu(inst uint32) {
	s, t := parseSt(inst)
	result := uint64(uint64(m.GetReg(s)) * uint64(m.GetReg(t)))
	lo, hi := uint32(result), uint32(result>>32)
	m.logSt("multu", s, t, lo, hi)
	m.SetLoHi(lo, hi)
}

func (m *Machine) div(inst uint32) {
	s, t := parseSt(inst)
	sVal, tVal := int32(m.GetReg(s)), int32(m.GetReg(t))
	lo, hi := uint32(sVal/tVal), uint32(sVal%tVal)
	m.logSt("div", s, t, lo, hi)
	m.SetLoHi(lo, hi)
}

func (m *Machine) divu(inst uint32) {
	s, t := parseSt(inst)
	sVal, tVal := m.GetReg(s), m.GetReg(t)
	lo, hi := sVal/tVal, sVal%tVal
	m.logSt("divu", s, t, lo, hi)
	m.SetLoHi(lo, hi)
}

func (m *Machine) mfhi(inst uint32) {
	d := parseD(inst)
	hi := m.GetHi()
	m.logD("mfhi", d, hi)
	m.SetReg(d, hi)
}

func (m *Machine) mflo(inst uint32) {
	d := parseD(inst)
	lo := m.GetLo()
	m.logD("mflo", d, lo)
	m.SetReg(d, lo)
}

func (m *Machine) lis(inst uint32) {
	d := parseD(inst)
	result := m.GetMem(m.pc)
	m.logD("lis", d, result)
	m.pc += 4
	m.SetReg(d, result)
}

func (m *Machine) lw(inst uint32) {
	s, t, i := parseSti(inst)
	mem := m.GetMem(m.GetReg(s) + uint32(i))
	m.logStiWord("lw", s, t, i, mem)
	m.SetReg(t, mem)
}

func (m *Machine) sw(inst uint32) {
	s, t, i := parseSti(inst)
	mem := m.GetReg(t)
	m.logStiWord("sw", s, t, i, mem)
	m.SetMem(m.GetReg(s)+uint32(i), mem)
}

func (m *Machine) slt(inst uint32) {
	s, t, d := parseStd(inst)
	result := uint32(0)
	if int32(m.GetReg(s)) < int32(m.GetReg(t)) {
		result = 1
	}
	m.logStd("slt", s, t, d, result)
	m.SetReg(d, result)
}

func (m *Machine) sltu(inst uint32) {
	s, t, d := parseStd(inst)
	result := uint32(0)
	if m.GetReg(s) < m.GetReg(t) {
		result = 1
	}
	m.logStd("sltu", s, t, d, result)
	m.SetReg(d, result)
}

func (m *Machine) beq(inst uint32) {
	s, t, i := parseSti(inst)
	m.logStiBranch("beq", s, t, i)
	if m.GetReg(s) == m.GetReg(t) {
		m.pc += uint32(i) * 4
	}
}

func (m *Machine) bne(inst uint32) {
	s, t, i := parseSti(inst)
	m.logStiBranch("bne", s, t, i)
	if m.GetReg(s) != m.GetReg(t) {
		m.pc += uint32(i) * 4
	}
}

func (m *Machine) jr(inst uint32) {
	s := parseS(inst)
	addr := m.GetReg(s)
	m.logS("jr", s, addr)
	m.pc = addr
}

func (m *Machine) jalr(inst uint32) {
	s := parseS(inst)
	addr := m.GetReg(s)
	m.logS("jalr", s, addr)
	m.SetReg(31, m.pc)
	m.pc = addr
}
