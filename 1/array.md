### Store values of same type

```go
var array_name = [length]datatype{values} // here length is defined

// or

var array_name = [...]datatype{values} // here length is inferred

// or

array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred
```
```go
package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
```
```go
package main
import ("fmt")

func main() {
  var arr1 = [...]int{1,2,3}
  arr2 := [...]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
```
```go
package main
import ("fmt")

func main() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)
}
```
```go
package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])
}
```
```go
package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(prices)
}
```
```go
// If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

// Tip: The default value for int is 0, and the default value for string is "".

// Example
package main
import ("fmt")

func main() {
  arr1 := [5]int{} //not initialized
  arr2 := [5]int{1,2} //partially initialized
  arr3 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
}
```
```go
//Initialize Only Specific Elements
package main
import ("fmt")

func main() {
  arr1 := [5]int{1:10,2:40}

  fmt.Println(arr1)
}
```
```go
package main
import ("fmt")

func main() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1))
  fmt.Println(len(arr2))
}
```
```go
package main
import "fmt"

func main() {
    var arr [3]int   // declaration

    arr[0] = 10
    arr[1] = 20
    arr[2] = 30

    fmt.Println(arr)
}
```
```go
var arr [5]int

arr = [5]int{
    0: 100,
    3: 400,
}
```