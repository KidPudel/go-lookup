package main

import (
	"io"
	"os"
	"strings"
)

// Implemnt rot13Reader - altering Reader
// 1. Read data from taking reader
// 2. Change data and write it to the array

type Rot13Reader struct {
	nestedReader io.Reader
}

func (r13 Rot13Reader) Read(arr []byte) (int, error) {
	// to know a size we read (if arr is greater than file, length is the file size)
	readLength := 0
	var totalError error = nil
	for readLength < len(arr) {
		numRead, err := r13.nestedReader.Read(arr)
		readLength += numRead
		if err != nil {
			// reached the end
			totalError = io.EOF
			break
		}
	}
	
	for i := 0; i < readLength; i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			if arr[i] > 'M' {
				arr[i] -= 13
			} else {
				arr[i] += 13
			}
		}
		if arr[i] >= 'a' && arr[i] <= 'z' {
			if arr[i] > 'm' {
				arr[i] -= 13
			} else {
				arr[i] += 13
			}
		}
		
	}
	
	return readLength, totalError
	
	
}

func main() {
	stringReader := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot13Reader := Rot13Reader{stringReader}
	
	io.Copy(os.Stdout, rot13Reader) // You cracked the code!
}
