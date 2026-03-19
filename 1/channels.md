```go
package main

import (
	"fmt")

func main() {
	messsageChan := make(chan string)

	messsageChan <- "Hello World"

	msg := <-messsageChan	

	fmt.Println(msg)
}

// $ go run route1.go 
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
//         C:/Users/Lenovo/Documents/git-repos/golang/go-routines/route1.go:9 +0x36
// exit status 2
```
```go
package main

import (
	"fmt"
	"time"
)

func receiveMessage(messageChan chan string) {
	fmt.Println(<-messageChan)
}

func main() {
	messsageChan := make(chan string)

	go receiveMessage(messsageChan)

	messsageChan <- "Hello World"

	time.Sleep(1 * time.Second)
}
```
```go
package main

import (
	"fmt"
	"math/rand"
)

func receiveMessage(intChan chan int) {
	for val := range intChan {
		fmt.Println("Received:", val)
	}
}

func main() {
	intChan := make(chan int)

	go receiveMessage(intChan)

	for {
		intChan <- rand.Intn(100)
	}
}
```
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func receiveMessage(intChan chan int) {
	for val := range intChan {
		time.Sleep(5 * time.Second)
		fmt.Println("Received:", val)
	}
}

func main() {
	intChan := make(chan int)

	go receiveMessage(intChan)

	for {
		intChan <- rand.Intn(100)
	}
}
```
```go
package main

import (
	"fmt"
)

func receiveMessage(intChan chan int) {
	for val := range intChan {
		fmt.Println("Received:", val)
	}
}

func main() {
	var input int
	intChan := make(chan int)

	go receiveMessage(intChan)

	for {
		fmt.Println("Enter a number:")
		fmt.Scanf("%d", &input)
		intChan <- input
	}
}
```
```go
package main

import (
	"fmt"
)

func sum(result chan int, num1 int, num2 int) {
	result <- num1 + num2
}

func main() {
	result := make(chan int)

	go sum(result, 5, 10)
	fmt.Println("Calculating the sum of 5 and 10...")
	fmt.Println("Result:", <-result) //reading a value from the channel, it is blocking that's why we does not need any wait or sleep function here
}
```
```go
package main

import (
	"fmt"
	"time"
)

func sum(done chan bool) {
	defer func() { done <- true }()
	fmt.Printf("Processing...")
	time.Sleep(2 * time.Second)
}

func main() {
	done := make(chan bool)
	go sum(done)
  <-done //reading a value from the channel, it is blocking that's why we does not need any wait or sleep function here
}
```
```go

```
```go
package main

import "fmt"

func receiver(ch chan int) {
	val := <-ch
	fmt.Println("Received: ",val)
}

func main() {
	ch := make(chan int)

	go receiver(ch)
	ch <- 42
	fmt.Println("Done")
}
```

---

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```

---

```go
package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, wg *sync.WaitGroup ) {
	defer wg.Done()
	for i:=1;i<=5;i++ {
		fmt.Println("Sending: ", i)
		ch <- i // Send data into channel
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
Sending:  2
Received:  1
Received:  2
Sending:  3
Sending:  4
Received:  3
Received:  4
Sending:  5
Received:  5
Done!
```

### Buffer channel
```go
package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, wg *sync.WaitGroup ) {
	defer wg.Done()
	for i:=1;i<=5;i++ {
		fmt.Println("Sending: ", i)
		ch <- i // Send data into channel
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
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(2)
	go sender(ch, &wg)
	go receiver(ch, &wg)

	wg.Wait()
	fmt.Println("Done!")
}
```
#### Expected Output
```bash
Sending:  1
Sending:  2
Sending:  3
Sending:  4
Sending:  5
Received:  1
Received:  2
Received:  3
Received:  4
Received:  5
Done!
```
```bash
## Why Sender Finishes First Now?
Unbuffered ch := make(chan int)
┌──────────┐                ┌──────────┐
│  sender  │ ── blocks ──►  │ receiver │  must handoff one-by-one
└──────────┘                └──────────┘

Buffered ch := make(chan int, 5)
┌──────────┐    ┌───────────────┐    ┌──────────┐
│  sender  │ ──►│ [ 1 2 3 4 5 ] │───►│ receiver │  sender dumps all, no blocking
└──────────┘    └───────────────┘    └──────────┘
```

---

### Force Strict Alternating (Send → Receive → Send → Receive): Use two channels to synchronize
```go
package main

import (
	"fmt"
	"sync"
	// "time"
)

func sender(dataCh chan int, ackCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i:=1;i<=5;i++ {
		fmt.Println("Sending: ", i)
		dataCh <- i // Send data into channel
		<-ackCh // Wait for consumer to acknowledge before sending next
	}
	close(dataCh) // Close channel when done sending
}

func receiver(dataCh chan int, ackCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range dataCh { // Receive until channel is closed
		fmt.Println("Received: ", val)
		ackCh <- true // Acknowledge receipt
	}
}

func main() {
	dataCh := make(chan int)
	ackCh := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)
	go sender(dataCh, ackCh, &wg)
	go receiver(dataCh, ackCh, &wg)

	wg.Wait()
	fmt.Println("Done!")
}
```

---

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2) //it is like a counter count + 2
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done() //count - 1
		for i:=1;i<=5;i++ {
			fmt.Println("Sending: ", i)
			ch <- i
		}
		close(ch)
		} (ch, &wg)
		
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done() //count - 1
		for val := range ch {
			fmt.Println("Receing: ", val)
		}
	} (ch, &wg)

	wg.Wait() //is count = 0
	fmt.Println("Done")
}
```

---

```bash
wg.Wait() is like a person standing at a door
waiting for everyone to come home before locking up.

They don't re-check by running around the house again —
they just STAND THERE and wait to be notified.

wg.Wait() called
      │
      ▼
  counter == 0? ──── YES ──► continue execution
      │
      NO
      │
      ▼
  PUT goroutine to SLEEP  ← just pauses, no re-execution
      │
      │   (other goroutines run...)
      │
      ▼
  wg.Done() called somewhere
  counter -= 1
      │
  counter == 0? ──── NO ──► stay asleep
      │
      YES
      │
      ▼
  WAKE UP main goroutine ✅
  continue after wg.Wait()
```

---

```go
var wg sync.WaitGroup

wg.Add(2)          // 👈 "I'm expecting 2 goroutines to finish"
//   counter = 2

go func() {
    defer wg.Done() // 👈 "I'm done!" → counter = 1
    // ... work ...
}()

go func() {
    defer wg.Done() // 👈 "I'm done!" → counter = 0
    // ... work ...
}()

wg.Wait()          // 👈 "Wait here until counter = 0"
fmt.Println("Done")
```
```go
package main

import (
	"fmt"
	"time"
)

func sendEmail(emails chan string, done chan bool) {
	defer func() {
		done <- true
	}()
	for email := range emails {
		fmt.Printf("Sending email to %s\n", email)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	emails := make(chan string, 100)
	done := make(chan bool)

	go sendEmail(emails, done)
	
	for i := 1; i <=10; i++ {
		emails <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(emails)
	<- done
}
```
```go
package main

import "fmt"

// send-only channel
func sendData(ch chan<- int) {
    ch <- 100
}

// receive-only channel
func receiveData(ch <-chan int) {
    fmt.Println("Received:", <-ch)
}

func main() {
    ch := make(chan int)

    go sendData(ch)
    receiveData(ch)
}
```
```go
// Go program to illustrate how to
// use for loop in the channel

package main

import "fmt"

// Main function
func main() {

    // Creating a channel
    // Using make() function
    mychnl := make(chan string)

    // Anonymous goroutine
    go func() {
        mychnl <- "GFG"
        mychnl <- "gfg"
        mychnl <- "Geeks"
        mychnl <- "GeeksforGeeks"
        close(mychnl)
    }()

    // Using for loop
    for res := range mychnl {
        fmt.Println(res)
    }
}
```
```go
package main

import "fmt"

// send-only channel
func sendData(ch chan<- int) {
    ch <- 100
		// <-ch // This will cause a deadlock since the channel is send-only
}

// receive-only channel
func receiveData(ch <-chan int) {
    fmt.Println("Received:", <-ch)
}

func main() {
    ch := make(chan int)

    go sendData(ch)
    receiveData(ch)
}
```
```go
package main

import "fmt"

func main() {
    ch := make(chan int, 5)

    // send data
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch) // ✅ close after sending

    // receive all data
    for val := range ch { // range stops when channel is closed
        fmt.Println(val)
    }
}
//Only the sender should close a channel, never the receiver.
```
```go
// Check if Channel is Closed
ok := <-ch
// val, ok := <-ch

if !ok {
    fmt.Println("Channel is closed")
} else {
    fmt.Println("Received:", val)
}
```
```go
//Like a switch but for channels — picks whichever channel is ready first.
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "one"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "two"
    }()

    // picks whichever arrives first
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received from ch1:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received from ch2:", msg2)
        }
    }
}
```
```go
select {
case val := <-ch:
    fmt.Println("Received:", val)
default:
    fmt.Println("No data, moving on...") // runs if channel not ready
}
```
```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
    }
}

func main() {
    jobs := make(chan int, 10)
    var wg sync.WaitGroup

    // start 3 workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }

    // send 9 jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs) // no more jobs

    wg.Wait() // wait for all workers to finish
    fmt.Println("All jobs done!")
}
```
```go
// Length of the Channel: In channel, you can find the length of the channel using len() function. Here, the length indicates the number of value queued in the channel buffer. Example:

// Go program to illustrate how to
// find the length of the channel

package main

import "fmt"
a
// Main function
func main() {

    // Creating a channel
    // Using make() function
    mychnl := make(chan string, 4)
    mychnl <- "GFG"
    mychnl <- "gfg"
    mychnl <- "Geeks"
    mychnl <- "GeeksforGeeks"

    // Finding the length of the channel
    // Using len() function
    fmt.Println("Length of the channel is: ", len(mychnl))
}
```
```go
// Capacity of the Channel: In channel, you can find the capacity of the channel using cap() function. Here, the capacity indicates the size of the buffer. Example:

// Go program to illustrate how to
// find the capacity of the channel

package main

import "fmt"

// Main function
func main() {

    // Creating a channel
    // Using make() function
    mychnl := make(chan string, 5)
    mychnl <- "GFG"
    mychnl <- "gfg"
    mychnl <- "Geeks"
    mychnl <- "GeeksforGeeks"

    // Finding the capacity of the channel
    // Using cap() function
    fmt.Println("Capacity of the channel is: ", cap(mychnl))
}
```