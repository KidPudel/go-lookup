> in go ther is no classes, but you can create a struct type, and that struct you can append with methods, just like in classes.  

To create a method, before name of the function, you need to specify receiver (for who is this function)
```go
func (a animal) makeSound() string {}
```

```go
type person struct {
  name string
  age int
}

func (p person) isAdult() bool {
  return p.age >= 18
}

func main() {
  p := person{"Harry", 17}
  fmt.Println(p.isAdult()) // false
}
```
