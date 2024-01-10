> Interface is type that consists of method signatures

A value of interface type can be assigned to any type that implements those signatures

define interface type
```go
type Abser interface {
  Abs() float64 // method signature
}
```

define types that implement signatures of interface
```go
// type 1
type Vetex struct {
  x float64
  y float64
}

func (v *Vertex) Abs() float64 {
  if v.x < 0 {
    v.x = v.x * -1
  }
  if v.y < 0 {
    v.y = v.y * -1
  }
  return 1.0 // ok
}

// type 2
type CustomFloat float64

func (f CustomFloat) Abs() float64 {
  if f < 0 {
    f = f * -1
  }
  return f
}
```


using it
```go
func main() {
  var a Abser
  v := Vertex{3, 4}

  a = &v // OK, implements all signatures (Abs() -> func (v *Vertex) Abs() float64)  = (&{3, 4}, *main.Vetex)
  a = v // ERROR, v Vertex does not implement Abs, because of (v *Vertex)

  f := CustomFloat(-5.4)

  a = f // OK, CustomFloat type implents Abs 
}
```
> NOTE: In general, all methods of the same type should have a value OR pointer receiver. Not the mixture of both

> There is no explicit declaration of intent that something implements interface, no `implements` keyboard.
> So type impments interface just by implementing its methods


## Interface's _true_ values
Interface values (underlying type) could be thought as tuple (`value of underlying type`, `concrete underlying type`).  
> **_So calling interface method, we call underlying method that has the same name_**


### Underlying nil values
In other languages that would cause a null pointer, but in Go it's common to write methods that gracefully handle nil value
```go
type I interface {
  Test()
}

type Person struct {
  name string
  age int
}

func (p *Person) Test() {
  if p == nil {
    fmt.Println("p is <nil>")
    return
  }
  fmt.Println(p.name, p.age)
}

func main() {
  var i I
  var p Person

  i = &p
  describe(i) // (<nil>, *main.Person)
  i.Test() // p is <nil>
}
```

### Interface's nil values
If interface's value itself is nil, meaning that it has nor value, nor underlying type, when calling method, it would be a run-time error, since no method is to call
