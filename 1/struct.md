* In Go, a **struct** is a custom data type
* It groups multiple fields into one logical unit
* Similar to a class in other languages (but without inheritance)

---

* Basic syntax

```go
type User struct {
	Name string
	Age  int
}
```

---

* Creating a struct value

```go
u1 := User{
	Name: "Mahin",
	Age:  22,
}
```

OR

```go
var u2 User
u2.Name = "Aman"
u2.Age = 25
```

---

* Accessing fields

```go
fmt.Println(u1.Name)
fmt.Println(u1.Age)
```

---

* Struct with functions (methods)

```go
type User struct {
	Name string
	Age  int
}

func (u User) Greet() string {
	return "Hello " + u.Name
}
```

Usage:

```go
u := User{Name: "Mahin", Age: 22}
fmt.Println(u.Greet())
```

---

* Pointer receiver (important concept)

```go
func (u *User) UpdateAge(newAge int) {
	u.Age = newAge
}
```

* Use pointer when:

  * You want to modify original struct
  * Avoid copying large structs

---

* Exporting struct

```go
type User struct {      // Exported
	Name string          // Exported
	age  int             // Not exported
}
```

* Capital letter → accessible outside package
* Lowercase → private to package

---

* Anonymous struct

```go
person := struct {
	Name string
	Age  int
}{
	Name: "Mahin",
	Age:  22,
}
```

* Used for quick, temporary structures

---

* Struct inside struct (composition)

```go
type Address struct {
	City string
}

type User struct {
	Name    string
	Address Address
}
```

Access:

```go
u.Address.City
```

---

* Slice of structs

```go
users := []User{
	{Name: "Mahin", Age: 22},
	{Name: "Aman", Age: 25},
}
```

---

* Why structs are important (especially for you)

  * JSON parsing
  * Kubernetes objects (Pods, Services, etc. are structs)
  * API request/response models
  * Configuration management tools
  * Terraform & Cloud SDKs use structs heavily

---

* Struct vs Map (important difference)

  * Struct

    * Fixed structure
    * Type-safe
    * Better performance
    * Compile-time checking

  * Map

    * Dynamic keys
    * Flexible
    * Less type safety

---

* Summary

  * `type Name struct {}` defines new type
  * Fields can be exported or private
  * Methods attach behavior
  * Use pointer receivers for modification
  * Structs are core building block of Go applications

---

If you want, I can next explain **how structs are used in JSON with tags**, which is extremely important for Kubernetes and API development.
