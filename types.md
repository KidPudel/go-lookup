### basic
- `bool`
- `string`
- `int`, if os 64x => `int64` (int8,16,32,64)
- `uint`, the same as int, but additionaly `uintptr`
- `float32`, `float64`
- `complex64`, `complex128`

### aliases
- `byte`, is for `uint8`
- `rune`, is for `int32` that represents _unicode code point_


## Create your own types
To define a new type:
1. just start with `type`
2. give it a name `foo`
3. specify a type on which a new type is going to be based

```go
type newInt int

type person struct {
  name string
  age int
}

type (
  newString string // type definition
  newFloat = float64 // type aleas (doen't create a new type and doesn't need a conversion)
)
```


## Stringers
One of the most ubiquitous interfaces is `Stringer` _already defined_ by `fmt`.  
`Stringer` is type that is responsible for descrybing underlying, implemention type as a string

```go
type person struct {
  name string
  age int
}

func (p *person) String() string {
  // here we can override default behaiviour
  return fmt.Sprintf("This person's name is %v, their age is %v", p.name, p.age)
}

func main() {
  p := person{name: "Ron", age: 20}
  fmt.Println(p) // This person's name is Ron, their age is 20
}
```
