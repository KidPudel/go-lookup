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
