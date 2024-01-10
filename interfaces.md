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
