Just like Striner, error is a build-in interface, that fmt is looking for 
```go
type go interface {
  Error() string
}
```

Functions often return errors, so in code we should check `if error != nil` to handle it

```go
func main() {
  if error := p.failedPerson(); error != nil {
    fmt.Println(error)
  }
}
```

Now to the implementation of the interface!  
For example we want person to have a meaningful styled error message
```go
type person struct {
  name string
  age int
}

func (p *person) Error() string {
  return fmt.Sprintf("%v got in problems at age of %v", p.name, p.age)
}

func (p *person) failedPerson() error {
  if p.name == "Arnold" && p.age == 7 {
    return p
  } else {
    return nil
  }
}
```

Now when `error != nil`, we'll get
```go
  p := person{"Arnold", 7}
  if error := p.failedPerson(); error != nil {
    fmt.Println(error)
  }
```
