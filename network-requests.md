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
