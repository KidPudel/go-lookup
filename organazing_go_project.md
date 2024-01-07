- _packages_: directory with go files
- _module_: collection of packages that has built in dependencies and versioning
- _internal_: directary that makes packages importable only by code rooted at parent of internal

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

Assuming this directory is uploaded to a GitHub repository at `github.com/someuser/modname`, the module line in the go.mod file should say  
`module github.com/someuser/modname`  

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
```
module github.com/someuser/modname
```
And a user should be able to install it on their machine with:
```
$ go install github.com/someuser/modname@latest
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
If we want in `main.go` to import `auth` package: `import "github.com/someuser/modname/internal/auth"`  

