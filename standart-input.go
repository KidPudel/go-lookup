package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
}
