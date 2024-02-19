> Reader is a stream of data, from which you can read

# Guidelines

`os` is for general OS-agonistic wrappers of operating system functionality, such as interacting with the filesystem.

`io` is for generalized I/O interface definitions and a few utility functions that don't fit elsewhere, such as io.Copy() and io.LimitedReader.

`io/ioutil` is deprecated. All of its functionality has been replicated elsewhere. Don't use it.

`bufio` is for automatic buffering of I/O operations, plus additional functionality enabled by that buffering, such as bufio.Scanner and all the extra methods of bufio.Reader.

`bytes` is for generalized operations on []bytes. It tries to maintain parity with strings. bytes.Buffer is defined where it is because its purpose is to allow you to treat a []byte as an io.Writer.


# Bufio
bufio is the convenient package for reading data

```go
// read from input
reader := bufio.NewReader(os.Stdin)
	// getting number of lines
	strNum, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	num, err := strconv.Atoi(strNum[:len(strNum)-1])
	if err != nil {
		fmt.Println(err)
		return
	}

	results := make([]string, num)
	for i := 0; i < num; i++ {
		// getting line by line
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
```

```go
// read from api response
apiReader := bufio.NewReader(response.Body)
```
