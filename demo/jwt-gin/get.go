package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//url := "http://controller:5000/v3/projects"
	url := "http://controller:8774/v2.1/servers/detail"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", "gAAAAABea7R6EwKPXXVqDvvDfh9aBjGytgmSg65Hdkmyk5rUuR5PMSJ4_1e2--5JFQu-unUEPR08dCXzyPQk6h2pIf3YoXIw_Rk4Zyi9uc5uDP4q1lC7OhfeFUcmNTmYYITrK-SG6XXCGSZ5QSKDq1JTDGUQXQzv9CzLGOW4DN-MzcKSeTyEUXQ")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var str bytes.Buffer
	_ = json.Indent(&str, []byte(body), "", "    ")
	fmt.Println("formated: ", str.String())
}
