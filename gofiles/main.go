package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	data := map[string]string{"username": "archan", "email": "archan@gmail.com"}
	jsonData, err := json.Marshal(data)
	res, err := http.Post(BaseUrl, "application/json", bytes.NewBuffer(jsonData))
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
