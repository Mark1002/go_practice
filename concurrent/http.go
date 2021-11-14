package concurrent

import (
	"fmt"
	"io"
	"net/http"
)

func ExecuteHttp() {
	resp, err := http.Get("https://gobyexample.com/hello-world")
	if err != nil {
		panic(err)
	}
	fmt.Println("Response status:", resp.Status)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
