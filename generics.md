# type paremeters
Placed before function's parameters
```go
func Index[T comparable](sl []T, element T) {}
```

Now this function takes slice of any type and for second parameter any type.  

`comparable`: constraint that allows using `==`, `!=` on values of type `T`

```go

```
