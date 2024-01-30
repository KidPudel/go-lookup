function closure is a function value that references variable from outside its body.

Go functions may be closures.  

**closure functions can access values defined in outter scopes, by referencing it with a pointer**

> NOTE: function is bound to the variables (has their own), for example


```go
func adder() func (int) int {
  sum := 0
  return func (x int) int {
    sum += x
    return sum
  }
}

func main() {
  increment, decrement := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(increment(i), decrement(i*2))
  }
}
```
