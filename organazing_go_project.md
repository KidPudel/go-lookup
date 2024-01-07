- _package_: a directory of a go file
- _module_: collection of packages with built-in dependecies and versioning (modules comes with 2 additional go.mod and go.sum)
- _internal_: directory that makes packages importable only by packages that is rooted at the parent of the `internal`(this) directory


for a very simple package requiring a single Go file (package), the project structure is:
```
project-root-directory/
  go.mod
  modname.go
  modname_test.go
```
Assuming this directory is uploaded to a GitHub repository at `github.com/someuser/modname`, the module line in the go.mod file should say `module github.com/someuser/modname`.  

And a user should be able to install it on their machine with:  
`$ go install github.com/someuser/modname@latest`


## Package or command with supporting packages
_If the file is too big, it is obvious that it needs to be split into other packages._  

Initially, it's recommended to place those packages into _`internal`_ directory, this prevents other modules to see, depend on those packages, that we don't want expose or support external use. Meaning, internal makes it private to other modules.
