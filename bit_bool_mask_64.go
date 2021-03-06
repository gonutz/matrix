//+build amd64 arm64

package matrix

// NewBitBoolMat returns a matrix of bool values, represented as single bits to
// save memory.
func NewBitBoolMat(w, h int) BitBoolMat {
	return BitBoolMat{
		data: make([]uint64, (w*h+63)/64),
		w:    w,
		h:    h,
	}
}

// BitBoolMat is a matrix of bool values, internally represented as single bits
// to save memory.
type BitBoolMat struct {
	data []uint64
	w, h int
}

// Size returns the width and height of the matrix, i.e. the number of columns
// and rows, in that order.
func (m *BitBoolMat) Size() (w, h int) {
	return m.w, m.h
}

// Get returns the value at the given position, x is the column and y is the
// row, starting top-left at 0,0.
func (m *BitBoolMat) Get(x, y int) bool {
	i := uint(x + y*m.w)
	return m.data[i/64]&(1<<(i%64)) != 0
}

// Set sets the matrix value at the given position, x is the column and y is the
// row, starting top-left at 0,0.
func (m *BitBoolMat) Set(x, y int, to bool) {
	i := uint(x + y*m.w)
	if to {
		m.data[i/64] |= (1 << (i % 64))
	} else {
		m.data[i/64] &= ^(1 << (i % 64))
	}
}
