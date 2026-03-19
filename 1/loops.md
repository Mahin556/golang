```go
package main
import ("fmt")

func main() {
  
  for {
    fmt.Printf("Infinite loop")
  }
}
```
```go
package main
import ("fmt")

func main() {
  i := 1
  for i <= 10 {
    fmt.Printf("Number: %d\n", i)
    i++
  }
}
```
```go
package main
import ("fmt")

func main() {
  
  for i:=1; i<=10;i++ {
    fmt.Printf("Number: %d\n", i)
  }
}
```
```go
package main
import ("fmt")

func main() {
  for num := range 13 {
    fmt.Println(num)
  }
}
```