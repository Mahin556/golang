```go
package main

import "fmt"

func main() {
	fmt.Print("Hello, World!") //Statement
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
}
```

### We can also use the semi colon to saperate the statements
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello");
	fmt.Println(" World");
	fmt.Print("How are you?");
}
```

### Go code in single line
```bash
package main; import ("fmt"); func main() { fmt.Println("Hello World!");}
```

### Var
```go
package main

import (
	"fmt"
)

func main() {
	var var1 string = "Hello, World!"
	fmt.Println(var1)
	
	var var2 int = 3
	fmt.Println(var2)

	var var3 float32 = 34.56
	fmt.Println(var3)
	
	var var4 float64 = 34.564
	fmt.Println(var4)

	var var5 bool = true
	fmt.Println(var5)

	var6 := 5 //always need to define value //Syntactic sugar
	fmt.Println(var6)

	var var7 string //can assign value later
	var7 = "Mahin"
	fmt.Println(var7)

	var student2 = "Jane" //type is inferred
	fmt.Println(student2)
}
```

### If var value not setted go pick default values
```go
package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```

### Difference Between var and :=
```go
// = can be used inside or outside the function
// := only used inside the function
package main
import ("fmt")

var a string = "demo"

func main() {
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```

### Declaring multiple variable in single line(with type keyword)
```go
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int = 1,2,3,4

	fmt.Println(a,b,c,d)
}
```
### Declaring multiple variable in single line(without type keyword)
```go
package main

import (
	"fmt"
)

func main() {
	var a, b = 1,2
	var c, d = 3,"hello"

	fmt.Println(a,b,c,d)
}
```

### var() block
```go
package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```

### Variable naming
```bash
GO VARIABLE NAMING RULES (Simple Version)

• A variable name must start with:
  - A letter (a-z, A-Z)
  - Or an underscore (_)

• It CANNOT start with a number

• It can contain:
  - Letters (a-z, A-Z)
  - Numbers (0-9)
  - Underscore (_)

• It CANNOT contain:
  - Spaces
  - Special characters (!, @, #, etc.)
  - Go keywords (like var, func, package, if, etc.)

• Variable names are case-sensitive:
  age, Age, and AGE are three different variables

• There is no length limit

--------------------------------------------

MULTI-WORD VARIABLE STYLES

1) Camel Case (most common in Go)
   myVariableName

2) Pascal Case (used for exported/public variables)
   MyVariableName

3) Snake Case (rare in Go, but valid)
   my_variable_name
```

---

### Bool data type
```go
package main
import ("fmt")

func main() {
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
}
```

### Int data type
```go
package main
import ("fmt")

func main() {
  var x int = 500
  var y int = -4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}
```
```go
package main

import "fmt"

func main() {

    // int (depends on system architecture: 32-bit or 64-bit)
    var a int = 100
    fmt.Println("int:", a)

    // int8 (8 bits, range: -128 to 127)
    var b int8 = 127
    fmt.Println("int8:", b)

    // int16 (16 bits, range: -32768 to 32767)
    var c int16 = 32000
    fmt.Println("int16:", c)

    // int32 (32 bits, range: -2147483648 to 2147483647)
    var d int32 = 2000000000
    fmt.Println("int32:", d)

    // int64 (64 bits, range: very large numbers)
    var e int64 = 9000000000000000000
    fmt.Println("int64:", e)
}
```

---

### Float data type
```go
package main
import ("fmt")

func main() {
  var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}
```
```go
package main
import ("fmt")

func main() {
  var x float64 = 1.7e+308
  fmt.Printf("Type: %T, value: %v", x, x)
}
```

### String Data type
```go
package main
import ("fmt")

func main() {
  var txt1 string = "Hello!"
  var txt2 string
  txt3 := "World 1"

  fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
```

### Format specifiers
```bash
GO FORMATTING VERBS (Printf)

==============================
GENERAL (All Data Types)
==============================

%v    → Default format
%+v   → Struct with field names
%#v   → Go-syntax format
%T    → Type of value
%%    → Print % sign

Example:
fmt.Printf("%v\n", 15.5)    // 15.5
fmt.Printf("%#v\n", 15.5)   // 15.5
fmt.Printf("%T\n", 15.5)    // float64
fmt.Printf("%%\n")          // %

==============================
INTEGER VERBS
==============================

%b    → Binary (base 2)
%d    → Decimal (base 10)
%+d   → Decimal with sign
%o    → Octal (base 8)
%O    → Octal with 0o prefix
%x    → Hex lowercase
%X    → Hex uppercase
%c    → Unicode character
%#x   → Hex with 0x prefix
%U    → Unicode format (U+1234)

Width & Padding:
%4d   → Width 4, right aligned
%-4d  → Width 4, left aligned
%04d  → Width 4, zero padded

Example (i := 15):
%b    → 1111
%d    → 15
%+d   → +15
%o    → 17
%#x   → 0xf
%04d  → 0015

==============================
STRING VERBS
==============================

%s    → Plain string
%q    → Double-quoted string
%8s   → Width 8, right aligned
%-8s  → Width 8, left aligned
%x    → Hex dump
% x   → Hex dump with spaces

Example ("Hello"):
%s    → Hello
%q    → "Hello"
%x    → 48656c6c6f
% x   → 48 65 6c 6c 6f

==============================
BOOLEAN VERB
==============================

%t    → true / false

Example:
fmt.Printf("%t\n", true)   // true

==============================
FLOAT VERBS
==============================

%e     → Scientific notation
%E     → Scientific notation (E)
%f     → Decimal format
%.2f   → Precision 2
%6.2f  → Width 6, precision 2
%g     → Compact format (auto)

Example (3.141):
%e     → 3.141000e+00
%f     → 3.141000
%.2f   → 3.14
%6.2f  →   3.14
%g     → 3.141


==============================
POINTER
==============================

%p   → Memory address

----------------------------------------

EXAMPLE

package main

import "fmt"

func main() {
    num := 25
    float := 12.3456
    text := "Go"
    flag := true

    fmt.Printf("Decimal: %d\n", num)
    fmt.Printf("Binary: %b\n", num)
    fmt.Printf("Float: %f\n", float)
    fmt.Printf("Scientific: %e\n", float)
    fmt.Printf("String: %s\n", text)
    fmt.Printf("Quoted: %q\n", text)
    fmt.Printf("Boolean: %t\n", flag)
    fmt.Printf("Type: %T\n", num)
}
```

### Escape seq
```bash
GO ESCAPE SEQUENCES (Used inside double-quoted strings)

\n   → New line
\t   → Tab space
\\   → Backslash
\"   → Double quote
\'   → Single quote
\r   → Carriage return
\b   → Backspace
\f   → Form feed
\v   → Vertical tab
\a   → Bell (alert sound)

----------------------------------------

EXAMPLE

package main

import "fmt"

func main() {

    fmt.Println("Hello\nWorld")
    fmt.Println("Go\tLang")
    fmt.Println("This is a backslash: \\")
    fmt.Println("He said \"Go is awesome\"")
}
```

### GO NUMERIC DATA TYPES
```bash
GO NUMERIC DATA TYPES

==============================
INTEGER TYPES (Whole Numbers)
==============================

UNSIGNED INTEGERS (Only Positive)

uint8   → 8 bits   → 0 to 255
uint16  → 16 bits  → 0 to 65535
uint32  → 32 bits  → 0 to 4294967295
uint64  → 64 bits  → 0 to 18446744073709551615

SIGNED INTEGERS (Positive + Negative)

int8    → 8 bits   → -128 to 127
int16   → 16 bits  → -32768 to 32767
int32   → 32 bits  → -2147483648 to 2147483647
int64   → 64 bits  → -9223372036854775808 to 9223372036854775807

int     → Depends on system
          32-bit system → 32 bits
          64-bit system → 64 bits

==============================
FLOATING TYPES (Decimal Numbers)
==============================

float32   → 32-bit IEEE-754 floating point
             Range: ~ -3.4e+38 to 3.4e+38

float64   → 64-bit IEEE-754 floating point
             Range: ~ -1.7e+308 to 1.7e+308
             (Default float type in Go)

==============================
COMPLEX TYPES (Real + Imaginary)
==============================

complex64   → float32 real + float32 imaginary
complex128  → float64 real + float64 imaginary
```

### Const
```go
// GO CONSTANTS (Simple Explanation)

// • const keyword is used to declare constants
// • Constant = fixed value (cannot be changed)
// • Must assign value at the time of declaration
// • Constants are read-only

// ----------------------------------------

// SYNTAX

// const NAME type = value

// Example:

const PI = 3.14
const float_value float = 3.14

// ----------------------------------------

// BASIC EXAMPLE

package main
import "fmt"

const PI = 3.14

func main() {
    fmt.Println(PI)
}

// ----------------------------------------

// CONSTANT RULES

// • Same naming rules as variables
// • Usually written in UPPERCASE (by convention)
// • Can be declared inside or outside a function
// • Value cannot be changed after declaration

// ----------------------------------------

// TYPED CONSTANT

// Type is explicitly mentioned

const A int = 10

// ----------------------------------------

// UNTYPED CONSTANT

// Type is NOT mentioned
// Compiler decides type automatically

const B = 10

// ----------------------------------------

// CONSTANTS ARE UNCHANGEABLE

const A = 1
A = 2   // ❌ ERROR

// Error:
// cannot assign to A

// ----------------------------------------

// MULTIPLE CONSTANTS (Block Style)

const (
    A int = 1
    B = 3.14
    C = "Hi!"
)

// ----------------------------------------

// IMPORTANT

// • Untyped constants are flexible
// • They can adapt to different types when used
// • Constants are evaluated at compile time
```

### Input
```go

```