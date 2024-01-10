Type assertion provides access to the concrete underlying type of the interface

```go
func main() {
  var i interface{} = "hello"

  s, ok := i.(string) // 1. s = value string, ok = true 2. s = ""(zero value), ok = false
  s2 := i.(int) // 1. s2 = value int 2. panic
}
```
