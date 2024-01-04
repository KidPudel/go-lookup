# For
```go
for i: 0; i < 10; i++ {
  fmt.Println(i)
}
```

# For continues
The init and post statemnt are optional
```go
sum := 1
for ; sum < 100; {
  sum++
}
```

# Go's while
```go
sum := 0
for sum < 1000 {
  sum++
}
```

# Forever
equivalent to` while(true)` in c
```go
fop {
}
```

---


# If
```go
if isReady {
}
```

# If with a short statement
like in loop you can declare a local variable that is bound to the condition/loop's scopes
```go
if x := math.Pow(5, x); x < 25 {
  fmt.Printf("%v is not enough\n", x)
}
```

# if and else
else also share variables that declared in if
```go
if x := math.Pow(5, x); x < 25 {
  fmt.Printf("%v is not enough\n", x)
} else {
  fmt.Println("%v is enough\n", x)
}
```
