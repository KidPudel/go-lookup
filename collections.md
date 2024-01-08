# Array (fixed size collection)
syntax: `[n]T`, where `n` is size and `T` type
```go
var initalMatches [10]int
fmt.Println(initialMatches) // [0 0 0 0 0 0 0 0 0 0]
initialMatches[1] = 100
fmt.Println(initialMatches) // [0 100 0 0 0 0 0 0 0 0]

// array literal
citiesVisitedInPast := [4]string{"mexico", "new-york", "tokyo", "shanghai"}
```
> NOTE: by default, array is zero-valued, so for int's it's 0

### length of an arrays is a built-in function, `len`
```go
len(matches) // 10
```

arrays are one-dimentional, but we can compose types
```go
var twoDm [2][3]int
for i := 0; i < 2; i++ {
  for j := 0; j < 3; j++ {
    twoDm[i][j] = i + j
  }
}
```


# Slice (dynamically sized collection)
> NOTE: much more common than the array

> **IMPORTANT**: Unlike array, slice doesn't actually store any data, it's just describes a section of underlying array. **_Slices are references to arrays_** so by changing value of the slice, you change value of the array

syntax: `[]T`  
Unlike array, uninitialized slice is equal to `nil` and its length is 0
```go
var scores []int
fmt.Println(scores, scores == nil, len(scores)) // [] true 0

// slice literal
slice := []int{1, 2, 3, 4, 5} // [1 2 3 4 5]
```

### Slice a slices
Slices support _slice_, meaning, another way we can get a slice, is by slicing some array or another slice, by specifying `[low:high]` where low is included index and high is excluded
```go
middleMatches := initialMatches[2:5] // [0 0 0]
a := collection[:5] // slices to 5th index excluding
b := colllection[2:] // slices from 2nd till the end
c := collection[:] // slices whole collection
```

### length vs capacity
- length is the number of elements in the slice
- capacity is the number of elements in the underlying array (how much memory is allocated, and if we exceed, create new)


### Creating slice with `make`
this creates actually dymanically sized array, to
```go
scores := make([]int, 3) // [0 0 0], cap (capacity) = 3, len = 3
```
also we can pass 3rd argument, to specify capacity

### Appending a Slice
Slices support appending new values to the array (dynamicly sized) with function `append`.  

`append` works by taking slice and a new value, then return a new slice
```go
numbers := make([]int, 2) // [0 0]
numbers = append(numbers, 3) // [0 0 3]
```

### Coping slice
```go
copiedNumbers := make([]int, len(numbers))
copy(copiedNumbers, numbers)
fmt.Println(copiedNumbers) // [0 0 3]
```



### `"slices"` package's useful utility functions
- `slices.Equal(s1, s2)` // check if slices are equal

### multi-dimentional slices
```go
twoD := make([][]int, 3)
for i := 0; i < len(twoD); i++ {
  twoD[i] = make([]int, i + 1)
  for j := 0; j < i + 1; j++ {
    twoD[i][j] = j
  }
}
```


### Range
> form of for loop, that iterates over slice or map.

when ranging, 2 values are returned, index and copy of value
```go
for i, v := range scores {
  fmt.Println(i, v)
}
```

also we can scip on of the values
```go
for i, _ := range pow
for _, value := range pow

for i := range pow
```


# Maps
Uninitialized map is a `nil` value. It has no keys and they cannot be added.  
Solution is to initialize with `make`, this will initialize a map and make it ready to use
```go
type Vertex struct {
  Lat, Long float64
}

// uninitialized
var coordinates map[string]Vertex

func main() {
  // inirtialized
  coordinates = nmake(map[string]Vertex)
  coordinates["Canada"] = Vertex{Lat: 56.1304, Long: 106.3468}
  fmt.Println(coordinates["Canada"])

  // map literals
  people := map[string]int{
    "Jinny" : 13,
    "Harry" : 14,
  }
  
}
```

### mutating
insert or update value
```go
m[key] = value
```

retrieve an element
```go
answer := m[key]
```

delete an element
```go
delete(m, key)
```

key is present, with 2 value asignments
```go
elem, ok := m[key]
```
`ok`, true or false, if false, then `elem` is zero value of value type

