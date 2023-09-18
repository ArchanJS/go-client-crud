package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const BaseUrl string = "http://localhost:8000"

func main() {
	fmt.Println("Printing GET response: ")
	getReq()
	fmt.Println("Printing POST response: ")
	postReq()
}

func getReq() {
	res, err := http.Get(BaseUrl)
	handleNilError(err)
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	fmt.Println(string(data))
}
func postReq() {
	requestBody := strings.NewReader(`
	{
		"username": "archan",
		"email": "archan@gmail.com"
	}
	`)
	res, err := http.Post(BaseUrl, "application/json", requestBody)
	handleNilError(err)
	defer res.Body.Close()

	userdata, err := io.ReadAll(res.Body)
	fmt.Println(string(userdata))
}

func handleNilError(err error) {
	if err != nil {
		panic(err)
	}
}
