package bitset

import (
	"fmt"
	"strings"
)

const (
	setU64BitsN uint64 = 9223372036854775808 // 0b1000000000000000000000000000000000000000000000000000000000000000
)

type BitSet struct {
	slots  []uint64
	length int
}

func NewBitSet(length int) *BitSet {
	numberOfSlots := length / 64
	if length > numberOfSlots*64 {
		numberOfSlots += 1
	}
	return &BitSet{
		slots:  make([]uint64, numberOfSlots),
		length: length,
	}
}

func (s *BitSet) SetAll(value int) {
	var bits uint64
	if value != 0 {
		bits = ^bits
	}

	for i := range s.slots {
		s.slots[i] = bits
	}
}

func (s *BitSet) Set(n int) {
	if n > s.length {
		return
	}

	slot := n / 64
	bits := n % 64
	s.slots[slot] = s.slots[slot] | (setU64BitsN >> uint64(bits))
}

func (s *BitSet) Get(n int) bool {
	if n > s.length {
		return false
	}

	slot := n / 64
	bits := n % 64
	return s.slots[slot]&(setU64BitsN>>uint64(bits)) != 0
}

func (s *BitSet) String() string {
	var sb strings.Builder

	for _, slot := range s.slots {
		sb.WriteString(fmt.Sprintf("%064b ", slot))
	}

	return sb.String()
}
