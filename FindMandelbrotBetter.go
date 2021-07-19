package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	//sets z0
	zn := complex(0, 0i)
	
	//loops through changing 'c'
	for k := 0; k < 5; k++{
		fmt.Println(cmplx.Abs(zn), k, "->", isInSet(k, zn))			
	}
}

func mandelB( z complex128, c complex128 ) complex128{
	//z(n+1) = z(n)^2 + c
	return cmplx.Pow(2, z) + c
}

func isInSet( m int, zn complex128 )bool{

	//loops through checking mandelB at a depth of 5
	for j := 0; j < 5; j++{
		//makes a complex c from int k
		c := complex(float64(m), 0i)
		//checks next iteration
		zn = mandelB(zn, c)
	}
	//checks if zn is too large
	if cmplx.Abs(zn) > 2{
		return false
	} else{
	//otherwise print it
		fmt.Println(zn)
		return true
	}
}
