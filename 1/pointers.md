```go
package main

import (
	"fmt"
)

func main() {
	i , j := 1, 2
	fmt.Println(&i,&j)

	p := &i //pointer to i
	fmt.Println(p) //print memory address of i

	var q *int = &j //pointer to j
	fmt.Println(q)

	fmt.Println(*p) //dereference pointer p to get value of i
	fmt.Println(*q) //dereference pointer q to get value of j

	*p = 10 //change value of i through pointer p
	fmt.Println(i) //i is now 10
}
```
```go
package main

import (
	"fmt"
)

func main() {
	a := 5
	fmt.Println("Value of a: ", a)
	fmt.Println("Square of a: ", squareValCopy(a))
	a = squareValCopy(a)
	fmt.Println("Value of a after squaring: ", a)

	squareValPointer(&a)
	fmt.Println("Value of a after pointer squaring: ", a)
}

func squareValCopy(v int) int { //immutablity
	 return v * v
}

func squareValPointer(u *int) { //efficiency
	*u = *u * *u
	fmt.Println(u, *u)
}
```