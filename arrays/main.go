package main

import (
	"fmt"
)

func main() {
	var array1 = [50]string{"demo1", "demo2", "demo3"}
	var array2 [50]string

	array2[0] = "demo4"

	

	fmt.Println("Array 1: ", array1)
	fmt.Println("Array 2: ", array2)
}