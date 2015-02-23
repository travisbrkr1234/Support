package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	body := []byte("{\n  \"title\": \"Buy cheese and bread for breakfast.\"\n}")

	req, _ := http.NewRequest("POST", "http://private-9b5f0-lights2.apiary-mock.com/notes", bytes.NewBuffer(body))

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error, error Will Robinson... when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
