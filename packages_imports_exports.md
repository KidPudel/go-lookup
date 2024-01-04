# Packages
Go program is made up of packages, so every collection of source code in the same directory is a package therefore you must at the top of the file include `package name`.  

Program starts with main package, so entry point of the app is going to be whaterver file you've specifield as `package main`.

## nested packages
Logically if we have package `rand` for random that is LOGICALLY nested inside of a `math` folder, if we want to access it in our main package, we'll import it like
`import "math/rand"`  


# Imports

## multiple imports
```go
import (
  "fmt" // package that is responsible for ForMaTing, writing, reading inputs, files
  "math/rand" // random
)
```

## internal
if package placed inside of "_internal_" folder, then only packages that are placed at the root of the internal could import it


# Exports
Exported must start with a capital latter, so `pizza`, `pi` won't be exportable, and `Pizza`, `Pi` will
```go
menu.pizza // error
menu.Pizza // ok
```
