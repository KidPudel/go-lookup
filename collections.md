## Array (fixed size collection)
syntax: `[n]T`, where `n` is size and `T` type
```go
var initalMatches [10]int
fmt.Println(initialMatches) // [0 0 0 0 0 0 0 0 0 0]
initialMatches[1] = 100
fmt.Println(initialMatches) // [0 100 0 0 0 0 0 0 0 0]

citiesVisitedInPast := [4]string{"mexico", "new-york", "tokyo", "shanghai"}
```
> NOTE: by default, array is zero-valued, so for int's it's 0

#### length of an arrays is a built-in function, `len`
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


## Slice (dynamically sized collection)
> NOTE: much more common than the array
syntax: `[]T`  
Unlike array, uninitialized slice is equal to `nil` and its length is 0
```go
var scores []int
fmt.Println(scores, scores == nil, len(scores)) // [] true 0

slice := []int{1, 2, 3, 4, 5} // [1 2 3 4 5]
```

#### To make slice not nil and non-zero length, use `make`
```go
scores := make([]int, 3) // [0 0 0], cap (capacity) = 3, len = 3
```
#### Appending a Slice
Slices support appending new values to the array (dynamicly sized) with function `append`.  

`append` works by taking slice and a new value, then return a new slice
```go
numbers := make([]int, 2) // [0 0]
numbers = append(numbers, 3) // [0 0 3]
```

#### Coping slice
```go
copiedNumbers := make([]int, len(numbers))
copy(copiedNumbers, numbers)
fmt.Println(copiedNumbers) // [0 0 3]
```

#### Slice a slices
Slices support _slice_, meaning, another way we can get a slice, is by slicing some array or another slice, by specifying `[low:high]` where low is included index and high is excluded
```go
middleMatches := initialMatches[2:5] // [0 0 0]
a := collection[:5] // slices to 5th index excluding
b := colllection[2:] // slices from 2nd till the end
```

#### `"slices"` package's useful utility functions
- `slices.Equal(s1, s2)` // check if slices are equal

#### multi-dimentional slices
```go
twoD := make([][]int, 3)
for i := 0; i < len(twoD); i++ {
  twoD[i] = make([]int, i + 1)
  for j := 0; j < i + 1; j++ {
    twoD[i][j] = j
  }
}
```
