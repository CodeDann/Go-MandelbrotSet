package main

import (
	"fmt"
	"math/cmplx"
	"os"
)

const (
	max_iterations = 200
)

func main() {
	var A [2]complex128
	var B [2]float64
	var Cmplx_CoOrds []complex128 = A[:]
	var Float_CoOrds []float64 = B[:]
	//loops through changing 'c'
	for i := -200; i < 200; i++ {
		for j := -200; j < 200; j++ {
			n := float64(i) / float64(100)
			m := float64(j) / float64(100)
			//sets z0
			zn := complex(0, 0i)
			c := complex(n, m)
			if isInSet(zn, c) {
				//fmt.Printf("%v,%v\n", n, m)
				Cmplx_CoOrds = append(Cmplx_CoOrds, c)
				Float_CoOrds = append(Float_CoOrds, n)
				Float_CoOrds = append(Float_CoOrds, m)
			}
		}
	}
	f, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	//explicitly ignore the error
	_, _ = fmt.Fprintf(f, "x,y\n")
	for k := 0; k < len(Float_CoOrds); k++ {
		_, _ = fmt.Fprintf(f, "%v,%v\n", Float_CoOrds[k], Float_CoOrds[k+1])
		k++
	}
	err = f.Close()
	if err != nil {
		return
	}

}

func mandelB(z complex128, c complex128) complex128 {
	//z(n+1) = z(n)^2 + c
	return z*z + c
}

func isInSet(zn complex128, c complex128) bool {

	//loops through checking mandelB at a depth of 200
	for j := 0; j < max_iterations; j++ {
		//checks next iteration
		zn = mandelB(zn, c)
		//checks if zn is too large
		if cmplx.Abs(zn) > 2 {
			return false
		}
	}
	//otherwise return true
	return true
}
