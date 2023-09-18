package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	baseUrl := "http://localhost:8000"
	data := map[string]string{"username": "archan", "email": "archan@gmail.com"}
	jsonData, err := json.Marshal(data)
	res, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonData))
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
