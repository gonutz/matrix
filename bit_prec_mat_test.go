package matrix

import "testing"

func TestSize(t *testing.T) {
	m := NewBitPrecMat(3, 5, 10)
	if w, h := m.Size(); w != 3 || h != 5 {
		t.Error("wrong size", w, h)
	}
}

func TestDataLength(t *testing.T) {
	if l := len(NewBitPrecMat(0, 0, 10).data); l != 1 {
		t.Error("wrong data length", l)
	}
	if l := len(NewBitPrecMat(1, 1, 10).data); l != 2 {
		t.Error("wrong data length", l)
	}
	if l := len(NewBitPrecMat(3, 1, 10).data); l != 2 {
		t.Error("wrong data length", l)
	}
	if l := len(NewBitPrecMat(4, 1, 10).data); l != 3 {
		t.Error("wrong data length", l)
	}
	// 15 values à 10 bit makes 150 bit, thus 5 uint32 values plus 1 for padding
	if l := len(NewBitPrecMat(3, 5, 10).data); l != 6 {
		t.Error("wrong data length", l)
	}
}

func TestSet(t *testing.T) {
	m := NewBitPrecMat(3, 2, 12)
	checkUint32s(t, m.data, []uint32{0x00000000, 0x00000000, 0x00000000, 0x00000000})

	m.Set(0, 0, 0xFFF)
	checkUint32s(t, m.data, []uint32{0xFFF00000, 0x00000000, 0x00000000, 0x00000000})

	m.Set(2, 0, 0xABC)
	checkUint32s(t, m.data, []uint32{0xFFF000AB, 0xC0000000, 0x00000000, 0x00000000})

	m.Set(2, 1, 0x123)
	checkUint32s(t, m.data, []uint32{0xFFF000AB, 0xC0000001, 0x23000000, 0x00000000})

	m.Set(1, 0, 0x456)
	checkUint32s(t, m.data, []uint32{0xFFF456AB, 0xC0000001, 0x23000000, 0x00000000})

	m.Set(0, 0, 0x789)
	checkUint32s(t, m.data, []uint32{0x789456AB, 0xC0000001, 0x23000000, 0x00000000})

	m.Set(2, 0, 0xDEF)
	checkUint32s(t, m.data, []uint32{0x789456DE, 0xF0000001, 0x23000000, 0x00000000})
}

func TestGet(t *testing.T) {
	// see test above for the values
	m := NewBitPrecMat(3, 2, 12)
	m.data = []uint32{0x789456DE, 0xF0000001, 0x23000000, 0x00000000}
	if v := m.Get(0, 0); v != 0x789 {
		t.Errorf("0x%X", v)
	}
	if v := m.Get(1, 0); v != 0x456 {
		t.Errorf("0x%X", v)
	}
	if v := m.Get(2, 0); v != 0xDEF {
		t.Errorf("0x%X", v)
	}
	if v := m.Get(0, 1); v != 0x000 {
		t.Errorf("0x%X", v)
	}
	if v := m.Get(1, 1); v != 0x000 {
		t.Errorf("0x%X", v)
	}
	if v := m.Get(2, 1); v != 0x123 {
		t.Errorf("0x%X", v)
	}
}

/*
                 0                  |                  1                  |                  2
01234567 89012345 67890123 45678901 | 01234567 89012345 67890123 45678901 | 01234567 89012345 67890123 45678901
00000000 00001111 11111111 22222222 | 22223333 33333333 44444444 44445555 | 55555555 xxxxxxxx xxxxxxxx xxxxxxxx
*/

func checkUint32s(t *testing.T, have, want []uint32) {
	eq := len(have) == len(want)
	if eq {
		for i := range have {
			if have[i] != want[i] {
				eq = false
				break
			}
		}
	}
	if !eq {
		t.Errorf("want\n%v\nbut have\n%v\n", want, have)
	}
}

func checkUint64s(t *testing.T, have, want []uint64) {
	eq := len(have) == len(want)
	if eq {
		for i := range have {
			if have[i] != want[i] {
				eq = false
				break
			}
		}
	}
	if !eq {
		t.Errorf("want\n%v\nbut have\n%v\n", want, have)
	}
}
