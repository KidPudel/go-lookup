Just like Striner, error is a build-in interface, that fmt is looking for 
```go
type go interface {
  Error() string
}
```

Functions often return errors, so in code we should check `if error != nil` to handle it

```go

func main() {
  if error := run(); error != nil {
    fmt.Println(error)
  }
}
```
