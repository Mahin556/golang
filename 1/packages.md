* In Go, multiple files work together **when they are in the same directory and same package name**
* Go automatically compiles all `.go` files in that directory together

---

* Basic example (single package, multiple files)

  * Folder structure:

    ```
    myapp/
      main.go
      math.go
      print.go
    ```

  * main.go

    ```go
    package main

    func main() {
        result := Add(5, 3)
        PrintResult(result)
    }
    ```

  * math.go

    ```go
    package main

    func Add(a int, b int) int {
        return a + b
    }
    ```

  * print.go

    ```go
    package main

    import "fmt"

    func PrintResult(value int) {
        fmt.Println("Result:", value)
    }
    ```

  * Run:

    ```
    go run .
    ```

    OR

    ```
    go build
    ```

  * Go compiles all files together because:

    * Same directory
    * Same package (`main`)

---

* Important rules

  * All files must have the same `package` name (inside one directory)
  * Only one `func main()` allowed in `package main`
  * Functions can be used across files without import (if same package)

---

* Using multiple packages (real-world style)

  * Structure:

    ```
    myapp/
      main.go
      utils/
        math.go
    ```

  * utils/math.go

    ```go
    package utils

    func Add(a int, b int) int {
        return a + b
    }
    ```

  * main.go

    ```go
    package main

    import (
        "fmt"
        "myapp/utils" //myapp should be module name not dir name
    )

    func main() {
        result := utils.Add(10, 20)
        fmt.Println(result)
    }
    ```

  * Must run:

    ```
    go mod init myapp
    go run .
    ```

---

* Key concept (very important)

  * Same package → no import needed
  * Different package → must import
  * Capital letter function name → exported
  * Small letter → private to that package

---

* How Go builds internally

  * Go groups files by package
  * Compiles package as one unit
  * Links all packages together
  * Creates binary (if package main)

---

* Best practice (especially useful for you since you're building DevOps + Kubernetes knowledge)

  * Keep `main.go` small
  * Put logic in separate packages
  * Use `internal/` for private packages
  * Use `pkg/` for reusable code

---
* In **Go**, exporting means making a function, variable, struct, or field accessible **outside its package**

* Go does NOT use `public` / `private` keywords

* Instead, it uses a simple rule:

  * **Capital letter → Exported (public)**
  * **Small letter → Not exported (private to that package)**

---

* Example 1: Not Exported (private)

  utils/math.go

  ```go
  package utils

  func add(a int, b int) int {   // small letter
      return a + b
  }
  ```

  main.go

  ```go
  package main

  import "myapp/utils"

  func main() {
      utils.add(5, 3)   // ❌ ERROR
  }
  ```

  * Error: `utils.add undefined`
  * Because `add` is not exported

---

* Example 2: Exported (public)

  utils/math.go

  ```go
  package utils

  func Add(a int, b int) int {   // Capital A
      return a + b
  }
  ```

  main.go

  ```go
  package main

  import (
      "fmt"
      "myapp/utils"
  )

  func main() {
      result := utils.Add(5, 3)
      fmt.Println(result)
  }
  ```

  * Works because `Add` is exported

---

* Exporting Variables

  ```go
  var Version = "1.0"   // exported
  var version = "1.0"   // not exported
  ```

---

* Exporting Struct

  ```go
  type User struct {     // exported struct
      Name string        // exported field
      age  int           // not exported field
  }
  ```

  Important:

  * If struct field is lowercase → it cannot be accessed outside the package
  * Even if struct itself is exported

---

* Real-world analogy

  * Package = folder
  * Capital letter = door open to other packages
  * Small letter = internal room (only that package can access)

---

* Very important (used heavily in Kubernetes & large Go projects)

  * `internal/` directory → cannot be imported outside the module
  * Exporting still follows capital letter rule

---

* Quick Summary

  * `func Add()` → exported
  * `func add()` → private
  * `type Server struct` → exported
  * `Name string` → exported field
  * `name string` → private field

---

If you want, I can also explain **why Go designers chose capital-letter exporting instead of keywords** (there’s an interesting design philosophy behind it).

