### var
we can use it to create variable in package level (global to file, package), or inside a function, anywhere
```go
var a, b int = 69, 71
```

### short variable declaration
only inside of function
```go
a, b := 69, 71
```

### type omittion
like in function parameters if multiple variables with the same type go sequentially, then the type specification is needed only at the end


### conversion
```go
number := 5.7
converted := int(number)
```
converted -> 5

BUT when converting int to string like: `newThing := string(number)` it threats int like rune, meaning that it convert all numbers to unicode system
