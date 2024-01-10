Type assertion provides access to the concrete underlying type of the interface

```go
func main() {
  var i interface{} = "hello"

  s, ok := i.(string) // 1. s = value string, ok = true 2. s = ""(zero value), ok = false
  s2 := i.(int) // 1. s2 = value int 2. panic
}
```

## Type switch

Permits several type assertions in series.
```go
func check(i interface{}) {
  switch v := i.(type) {
    case string:
      fmt.Println("i know you are an array of characters!")
    case int:
      fmt.Println("aren't you a pretty whole number")
    default:
      fmt.Println("- who are you suppose to be? - I'm vengeance"
  }  
}


type Superhero struct {
  name string
}

func main() {
  check("hey")
  check(10)
  check(Superhero{name: "Batman"})
}
```
