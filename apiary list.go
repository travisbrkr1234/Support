package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "http://private-9b5f0-lights2.apiary-mock.com/notes", nil)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Heh, you messed up real good son!")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
