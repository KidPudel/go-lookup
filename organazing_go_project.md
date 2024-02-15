- _packages_: directory (folder) with go files
- _module_: collection of packages that has built in dependencies and versioning
- _internal_: directary that makes packages importable only by code rooted at parent of `internal`

## Basic package
The project consists of a single module, which consists of a single package.  
The package name matches the last path component of the module name. For a very simple package requiring a single Go file, the project structure is:
```
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  // ... other go files as needed
```
> NOTE: package could be split into multiple go files, all sharing the same package

Assuming this directory is uploaded to a GitHub repository at `github.com/someuser/modname`, the module line in the `go.mod` file should say  
```go
module github.com/someuser/modname`  
```

All code in package should declare package:
```
package modname

// ... package code here
```

Then you can import that with:
```
import "github.com/someuser/modname"
```

## Basic commands
Assuming this directory is uploaded to a GitHub repository at github.com/someuser/modname, the module line in the go.mod file should say:
```go
module github.com/someuser/modname
```
And a user should be able to install it on their machine with:
```bash
$ go install github.com/someuser/modname@latest
```

to list all commands
```bash
go help
```

## Package or command with supporting packages
When package is growing big, you can split functionality in package into other packages.  
> NOTE: Initially, it is recommended to put those packages into internal directory (read at the top) and to make it private, if we don't want to expose funtionality and support external use.
```
project-root-directory/
  internal/
    auth/
      auth.go
      auth_test.go
    hash/
      hash.go
      hash_test.go
  go.mod
  modname.go
  modname_test.go
```
If we want in `main.go` to import `auth` package: 
```go
import "github.com/someuser/modname/internal/auth"
```

## Multiple packages and importing
```
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  auth/
    auth.go
    auth_test.go
    token/
      token.go
      token_test.go
  hash/
    hash.go
  internal/
    trace/
      trace.go
```
in `go.mod`
```go
module github.com/someuser/modname
```

to import root package
```go
import "github.com/someuser/modname"
```

sub-packages could be imported 
```go
import (
  "github.com/someuser/modname/auth"
  "github.com/someuser/modname/auth/token"
  "github.com/someuser/modname/hash"
)
... but not trace, because it is resides in internal, therefore cannot be imported
```
> NOTE: It is recommended to keep modules in internal as much, as possible

## Multiple programms in a project && Multple commands
```
project-root-directory/
  go.mod
  internal/
    ... shared internal packages
  prog1/
    main.go
  prog2/
    main.go
```
- all programs will have their own main package
- top-level internal directory could have pacakges that are shared (common) accross the other apps in the repository

To intall those programms:
```bash
go install github.com/someuser/modname/prog1@latest
go install github.com/someuser/modname/prog2@latest
```
> NOTE: Common convention is to place commands into `cmd` directory. It is very useful in a mixin repository that has both commands and importable packages



## Packages and commands in the same repository
importable packages and installable commands with related functionality
```
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  auth/
    auth.go
    auth_test.go
  internal/
    ... internal packages
  cmd/
    prog1/
      main.go
    prog2/
      main.go
```

We can now do both:
1. import packages
   `import "github.com/someuser/modname/auth"`
2. install programs
   `go install github.com/someuser/modname/cmd/prog1@latest`


## Server project
There is larege variance in the structure of writing server, because of different aspects that may have and may not be (_REST, gRPC, front-end, deployment, containerization, scripts_). SO focus here is on the Go part.  

Since the server is tipically a one binary (or collection of binaries), we don't need to expose server logic, therefore we should keep all our packages in internal directory.  
Also since there is most likey other directories with non-Go code, it is the best to keep all Go commands in `cmd` directory.
```
project-root-directory/
  go.mod
  internal/
    auth/
      ...
    metrics/
      ...
    model/
      ...
  cmd/
    api-server/
      main.go
    metrics-analyzer/
      main.go
    ...
  ... the project's other directories with non-Go code
```
> ADVICE: if repository grows packages that is worth to share with other projects, it's best to sdplit them to separate modules
