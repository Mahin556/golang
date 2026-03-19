```go
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var input1 int
	var input2 int
	var choice string

	for {
		fmt.Print("Enter first number: ")
		fmt.Scanln(&input1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&input2)

		go processing(input1, input2)

		fmt.Print("Do you want to continue? (y/n): ")
		fmt.Scanln(&choice)
		if strings.ToLower(choice) == "n" || strings.ToLower(choice) == "no" {
			time.Sleep(10 * time.Second)
			break
		}
	}
}

func processing(input1 int, input2 int) {
	time.Sleep(10 * time.Second)
	fmt.Printf("Processed %d and %d\n", input1, input2)
}
```
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go Display("Go Routine")
	fmt.Printf("Printf is executing in main function\n")
	time.Sleep(time.Microsecond * 3600)
}

func Display(str string) {
	for i := 0; i<5; i++ {
		fmt.Println(str)	
	}
}
```
```go
//Anonymous Goroutines
package main

import (
	"fmt"
	"time"
)

func main() {
	go func(s string) {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
			time.Sleep(500 * time.Millisecond)
		}
	}("Go Routine")

	time.Sleep(time.Second * 2) // Allow Goroutine to finish
	fmt.Printf("Printf is executing in main function\n")
}
```

---

### Goroutines cannot return values. They run in the background — nobody is there to receive the return value.
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go func(s string) {
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			return s  // ❌ WRONG — goroutines cannot return values
		}
	}("Go Routine")

	time.Sleep(time.Second * 2) // Allow Goroutine to finish
	fmt.Printf("Printf is executing in main function\n")
}
```

---

### Use channel to get value back
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string) // ✅ channel to receive value

    go func(s string) {
        for i := 0; i < 5; i++ {
            time.Sleep(500 * time.Millisecond)
        }
        ch <- s // ✅ send value through channel
    }("Go Routine")

    result := <-ch // ✅ receive value from channel
    fmt.Println("Received:", result)
    fmt.Println("Printf is executing in main function")
}
```

---

### Use sync.WaitGroup (Best Practice) - To wait for goroutine(s) to finish
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(1) // ✅ tell WaitGroup to wait for 1 goroutine

    go func(s string) {
        defer wg.Done() // ✅ mark done when finished
        for i := 0; i < 5; i++ {
            time.Sleep(500 * time.Millisecond)
            fmt.Println(s, "iteration:", i)
        }
    }("Go Routine")

    wg.Wait() // ✅ waits until goroutine finishes (better than time.Sleep)
    fmt.Println("Printf is executing in main function")
}
```
```go
package main

import (  
    "fmt"
    "time"
		"sync"
)

func main() {  
	var wg sync.WaitGroup
	wg.Add(1)

	go func (name string) {
		defer wg.Done()
		fmt.Println(name)
	} ("mahin-" + time.Now().String())
		
	wg.Wait()

}
```
```go
package main

import (  
    "fmt"
    "time"
		"sync"
)

func main() {  
	var wg sync.WaitGroup
	for i:=1; i<=5; i++ {
		wg.Add(1)
		go func (name string) {
			defer wg.Done()
			fmt.Println(name)
		} ("task-" + fmt.Sprint(i) + "-" + time.Now().String())
	}

	wg.Wait()

}
```
```go
package main

import (  
    "fmt"
    "time"
		"sync"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(name)
	time.Sleep(time.Millisecond*500)
}

func main() {  
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go task("mahin-" + time.Now().String(), &wg)
	}

	wg.Wait()

}
```

---


```go
package main

import (
	"fmt"
	"time"
)

func Aname() {
	var names = []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	for _,name := range names {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(name)
	}
}

func Bscore() {
	var scores = []int{85, 90, 78, 92, 88}
	for score := range scores {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(score)
	}
}

func main() {
	go Aname()
	go Bscore()
	time.Sleep(time.Second * 3)
	fmt.Println("Done")
}
```

```go
package main

import (  
    "fmt"
    "time"
)

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}

// 250 400 500 750 800 1000 1200 1250 1600 2000 3000  
//  1   a   2   3   b    4   c     5   d    e    main terminated
```
```go
package main

import (  
    "fmt"
    "time"
)

func task(name string) {
	fmt.Println(name)
	time.Sleep(time.Millisecond*500)
}

func main() {  
	time.Sleep(time.Second * 1)
	for i := 1; i <= 5; i++ {
		go task("mahin-" + time.Now().String())
	}
	time.Sleep(time.Second * 1)
}
```
```go
package main

import (  
    "fmt"
    "time"
)

func task(name string) {
	fmt.Println(name)
	time.Sleep(time.Millisecond*500)
}

func main() {  
	time.Sleep(time.Second * 1)
	for i := 1; i <= 5; i++ {
		task("mahin-" + time.Now().String())
	}
	time.Sleep(time.Second * 1)
}
```
```go
package main

import (  
    "fmt"
    "time"
)


func main() {  
	time.Sleep(time.Second * 1)
	for i := 0; i < 10; i++ {
		go func (i int) {
			fmt.Println(i)
		} (i)
	}
 	time.Sleep(time.Second * 1)
}
// 0
// 9
// 5
// 6
// 7
// 8
// 2
// 1
// 3
// 4
```

### Add time.Sleep to slow down producer
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(ch chan int, wg *sync.WaitGroup ) {
	defer wg.Done()
	for i:=1;i<=5;i++ {
		fmt.Println("Sending: ", i)
		ch <- i // Send data into channel
		time.Sleep(100 * time.Millisecond)  // Give consumer time to run
	}
	close(ch) // Close channel when done sending
}

func receiver(ch chan int, wg *sync.WaitGroup ) {
	defer wg.Done()
	for val := range ch { // Receive until channel is closed
		fmt.Println("Received: ", val)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go sender(ch, &wg)
	go receiver(ch, &wg)

	wg.Wait()
	fmt.Println("Done!")
}
```
```bash
Sending:  1
Received:  1
Sending:  2
Received:  2
Sending:  3
Received:  3
Sending:  4
Received:  4
Sending:  5
Received:  5
Done!
```

---

### Waiting Groups
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // 👈 "I'm expecting 2 goroutines to finish"
	// wg.Add(1) // 👈 Only First gorouter runs completely
	//   counter = 2

	go func() {
		defer wg.Done() // 👈 "I'm done!" → counter = 1
		fmt.Println("First gorouter executed")
		// ... work ...
	}()

	go func() {
		defer wg.Done() // 👈 "I'm done!" → counter = 0
		fmt.Println("Second gorouter executed")
		// ... work ...
	}()

	wg.Wait() // 👈 "Wait here until counter = 0"
	fmt.Println("Done")
}
```

---
### References:
- https://www.geeksforgeeks.org/go-language/goroutines-concurrency-in-golang/
- https://golangbot.com/concurrency/
- https://golangbot.com/goroutines/
- https://www.geeksforgeeks.org/go-language/golang-goroutine-vs-thread/