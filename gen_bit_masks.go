//+build ignore

/*
This program generates the masks to index into a matrix of arbitrary precision,
see bit_prec_mat.go for its usage.
*/

package main

import "fmt"

func main() {
	fmt.Println("var bitMasks = [][]uint64{")
	fmt.Println("\tnil,")
	fmt.Println("\tnil,")
	for bits := uint(2); bits <= 31; bits++ {
		var mask uint64
		for j := uint(0); j < bits; j++ {
			mask |= 1 << (63 - j)
		}
		fmt.Println("\t[]uint64{")
		for i := uint(0); i < 2*32-1-bits; i++ {
			m := mask >> i
			fmt.Printf("\t\t0x%.16X,\n", m)
		}
		fmt.Println("\t},")
	}
	fmt.Println("}")
}
