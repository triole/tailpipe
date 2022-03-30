package mail

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func isValidJSON(i interface{}) bool {
	var str map[string]string
	data, err := getBytes(i)
	if err != nil {
		panic(err.Error())
	}
	data = data[4:]
	err = json.Unmarshal(data, &str)
	return err == nil
}

func getBytes(i interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(i)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func mktemp() string {
	f, err := ioutil.TempFile(os.TempDir(), "tailpipe_")
	if err != nil {
		log.Fatal(err)
	}
	return f.Name()
}

func readFile(filename string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %q\n", err)
	}
	return fileBytes, err
}

func readTemplate(str string) ([]byte, error) {
	data, err := tpl.ReadFile(path.Join("templates", str))
	if err != nil {
		fmt.Printf("mail template error: %s\n", err)
	}
	return data, err
}

func prettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func saveJSONFile(filename string, data string) {
	jsonData, err := prettyString(data)
	if err != nil {
		fmt.Printf("error marshal json data %q\n", err)
	} else {
		err = ioutil.WriteFile(filename, []byte(jsonData), 0644)
		if err != nil {
			fmt.Printf("error saving json file %q\n", err)
		}
	}
}
