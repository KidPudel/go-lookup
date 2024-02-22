```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	ch := make(chan map[string]any)
	go apiCall(ch)

	for n := range 10 {
		fmt.Println("do some for me ", n)
	}

	response := <-ch

	fmt.Println(response["dept"])
}

func apiCall(ch chan map[string]any) {
	response, err := http.Get("https://api.supergood.ru/getitems.php")
	if err != nil {
		ch <- nil
	}
	defer response.Body.Close()

	var jsonData interface{}
	err = json.NewDecoder(response.Body).Decode(&jsonData)
	if err != nil {
		ch <- nil
	}
	data, ok := jsonData.(map[string]any)
	if ok {
		ch <- data
	} else {
		ch <- nil
	}
}


```
