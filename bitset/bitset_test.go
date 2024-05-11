package bitset

import "testing"

func TestNewBitSet(t *testing.T) {
	testCases := []struct {
		length             int
		expectedSlotNumber int
	}{
		{
			length:             0,
			expectedSlotNumber: 0,
		},
		{
			length:             1,
			expectedSlotNumber: 1,
		},
		{
			length:             64,
			expectedSlotNumber: 1,
		},
		{
			length:             65,
			expectedSlotNumber: 2,
		},
	}

	for _, testCase := range testCases {
		bitSet := NewBitSet(testCase.length)
		if len(bitSet.slots) != testCase.expectedSlotNumber {
			t.Errorf("unexpected slot number, expect `%d` but got `%d` when test with length `%d`", testCase.expectedSlotNumber, len(bitSet.slots), testCase.length)
		}
	}
}

func TestSetGetAnd(t *testing.T) {
	bitSet := NewBitSet(128)
	if bitSet.Get(0) {
		t.Error("unexpcted result, Get(0) should got false")
	}

	bitSet.Set(0)
	if !bitSet.Get(0) {
		t.Error("unexpcted result, Get(0) should got true")
	}

}
