* In Go, a **list of maps** means:

  * A **slice** (dynamic list)
  * Where each element is a **map**

---

* Basic syntax

```go
[]map[keyType]valueType
```

---

* Example 1: Slice of maps (string → int)

```go
package main

import "fmt"

func main() {
	people := []map[string]int{
		{"age": 25},
		{"age": 30},
	}

	fmt.Println(people)
}
```

---

* Example 2: More practical example

```go
package main

import "fmt"

func main() {
	users := []map[string]string{
		{
			"name":  "Mahin",
			"city":  "Jaipur",
		},
		{
			"name":  "Aman",
			"city":  "Delhi",
		},
	}

	fmt.Println(users[0]["name"]) // Mahin
}
```

---

* Creating empty list of maps

```go
var data []map[string]string
```

Or

```go
data := make([]map[string]string, 0)
```

---

* Important: Initialize map before using

This will cause panic:

```go
data := make([]map[string]string, 1)
data[0]["name"] = "Mahin"  // ❌ panic (map not initialized)
```

Correct way:

```go
data := make([]map[string]string, 1)
data[0] = make(map[string]string)
data[0]["name"] = "Mahin"
```

---

* Adding new map dynamically

```go
data := []map[string]string{}

m := make(map[string]string)
m["name"] = "Mahin"

data = append(data, m)
```

---

* Iterating over list of maps

```go
for i, item := range data {
	fmt.Println("Index:", i)
	for key, value := range item {
		fmt.Println(key, "=", value)
	}
}
```

---

* Memory understanding

  * Slice → dynamic array (holds references)
  * Map → reference type
  * So slice holds references to maps

---

* Real-world use cases (very common in DevOps tools & APIs)

  * JSON-like structures
  * Kubernetes resource parsing
  * API responses
  * Dynamic configuration handling

---

* Summary

  * `[]map[string]string` → list of maps
  * Must initialize each map
  * Use `append()` to add new maps
  * Maps are reference types

---

If you want, I can also show how this connects to **JSON unmarshalling**, because APIs often return data in this format.
