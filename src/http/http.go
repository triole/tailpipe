package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Post does send a post request to a certain address
func Post(msg, targetAddr string) {
	data, err := json.Marshal(map[string]string{
		"message": msg,
	})
	if err != nil {
		fmt.Printf("Error. Json marshal failed: %q\n", err)
	}
	resp, err := http.Post(targetAddr,
		"application/json", bytes.NewBuffer(data))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}
