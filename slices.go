package main

import (
	"fmt"
)

func main()  {
	 matriz :=[]int{1,2}
	if matriz == nil {
		fmt.Println("Esta vacio")
	} else {
		fmt.Println(len(matriz))
	}

	arreglo := [3]int{1,2,3}
	slicing := arreglo[:2] // arreglo[1:2] arreglo[0:3]
	fmt.Println(slicing)


	slice := make([]int,3,5)
	slice = append(slice,2)
}