package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const BaseUrl string = "http://localhost:8000"

func main() {
	fmt.Println("Printing GET response: ")
	getReq()
	fmt.Println("Printing POST response: ")
	postReq()
	fmt.Println("Printing POST Form response: ")
	postFormReq()
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

func postFormReq() {
	requestBody := url.Values{}
	requestBody.Add("username", "archan")
	requestBody.Add("email", "archan@gmail.com")
	res, err := http.PostForm(BaseUrl, requestBody)
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
