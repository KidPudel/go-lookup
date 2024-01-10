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

### Not only structs
You can declare a method on receiver which is defined in the same package

> NOTE: receiver value here is just a copy

```go
type customFloat float64

func (customF customFloat) Absolute() customFloat {
  if customF < 0 {
    return -customF
  } else {
    return customF
  }
}

func main() {
  var number customFloat = -1.5
  number = number.Absolute() // 1.5
}
```

### Pointer receivers
> NOTE: this is more common, since you more often use methods to change the value of receiver itself, rather than just have a copy
```go
func (customF *customFloat) Absolute() {
	if *customF < 0 {
		*customF = -*customF
	}
}

func main() {
	var number customFloat = -1.5
	number.Absolute() // 1.5
	fmt.Println(number)
}
```


# What receiver all methods of type should have

> NOTE: In general, all methods of the same type should have a value OR pointer receiver. Not the mixture of both
