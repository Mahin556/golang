* In Go, variable scope means **where a variable is visible and accessible**
* Scope depends on **where the variable is declared**

---

* Package-level scope (Global scope)

  * Declared outside all functions
  * Accessible to all functions in the same package
  * If capitalized → accessible from other packages

  ```go
  package main

  import "fmt"

  var message = "Hello"     // package-level

  func main() {
      fmt.Println(message)
  }
  ```

  * Lifetime → entire program execution
  * Best practice → avoid too many globals (especially in large projects)

---

* Function-level scope (Local scope)

  * Declared inside a function
  * Only accessible inside that function

  ```go
  func main() {
      var x = 10
      fmt.Println(x)
  }
  ```

  * Cannot access `x` outside `main()`

---

* Block scope

  * Variables declared inside `{}` blocks
  * Only accessible inside that block

  ```go
  func main() {
      if true {
          y := 20
          fmt.Println(y)
      }

      fmt.Println(y)  // ❌ error
  }
  ```

  * `y` exists only inside the `if` block

---

* Short variable declaration scope (`:=`)

  * Works only inside functions
  * Scope limited to the block

  ```go
  func main() {
      z := 30
      fmt.Println(z)
  }
  ```

---

* Shadowing (Very Important Concept)

  * Inner scope variable can override outer variable

  ```go
  var x = 100

  func main() {
      x := 50     // shadows global x
      fmt.Println(x)  // prints 50
  }
  ```

  * Global `x` still exists
  * But local `x` hides it inside function

  ⚠️ This is common in large Go codebases — must be careful

---

* Loop scope example

  ```go
  for i := 0; i < 3; i++ {
      fmt.Println(i)
  }

  fmt.Println(i)  // ❌ error
  ```

  * `i` exists only inside loop

---

* Struct field scope

  * Lowercase field → accessible only inside package
  * Capital field → accessible outside package

  ```go
  type User struct {
      Name string   // exported
      age  int      // private
  }
  ```

---

* Scope summary

  * Outside function → package-level scope
  * Inside function → function scope
  * Inside `{}` → block scope
  * Capital letter → exported outside package
  * Lowercase → restricted to same package

---

* Real-world best practices (important for your DevOps/K8s journey)

  * Avoid global variables in production apps
  * Prefer passing variables as parameters
  * Keep scope as small as possible
  * Avoid shadowing — especially in loops and error handling

---

If you want, I can next explain **lifetime vs scope difference** — many people confuse these two concepts.
