package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/tasks"
	method := "POST"

	payload := strings.NewReader(`{
    "title": "123",
    "description": "hello world"
	}`)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		req.Header.Add("Content-Type", "application/json")

		// just close it here
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		fmt.Println(string(body))
	}

}
