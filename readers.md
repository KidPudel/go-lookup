> Reader is a stream of data, from which you can read


`Reader` is an interface which represents the read **_end_** of stream of data, that is defined in `io` package and implemented by many packages in standart library, including:
- reading files
- cyphers
- network connection
- compression
and many more

`Reader` interface defines a signature of `Read` method, that is taking the byte array to **_read to_** and returning 2 values, the amount of bytes written in array and an error that is `io.EOF` which if is not nil, signals that stream has ended.  

```go
package main

import (
	"fmt"
	"io"      // where interface lives, to access some types
	"strings" // which is implements io signature method Read
)

func main() {
	message := "I love you, the place where i was born, my home, my saver.\n\nI'm sorry that so much people doesn't threat you right, hope we, people will find strength to change that, thank you for everything, Earth."
	messageReader := strings.NewReader(message)

	messagePart := [10]byte{}

	for {
		nWritten, errEnded := messageReader.Read(messagePart[:])
		fmt.Printf("%q: written in %v, ended? %v", messagePart, nWritten, errEnded)
		if errEnded == io.EOF {
			fmt.Println("bye")
			break
		}
	}
}

```


## Some guidelines

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

# Json reader
```go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.supergood.ru/getitems.php")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	var jsonData interface{}
	err = json.NewDecoder(response.Body).Decode(&jsonData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(jsonData)
}

```
