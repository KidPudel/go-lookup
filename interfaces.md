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
}
```


using it
```go
func main() {
  var a Abser
  v := Vertex{3, 4}

  a = &v // OK, implements all signatures (Abs() -> func (v *Vertex) Abs() float64)
  a = v // ERROR, v Vertex does not implement Abs, because of (v *Vertex)
}
```
> NOTE: In general, all methods of the same type should have a value OR pointer receiver. Not the mixture of both
