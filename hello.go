package main

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
)

// type Response struct {
// 	ID     string `json:"id"`
// 	Joke   string `json:"joke"`
// 	Status int    `json:"status"`
// }

func main() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	var url	string = "https://secure.runescape.com/m=hiscore/index_lite.ws?player=zee_pk"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "text/html")
	req.Header.Add("Content-Type", "text/html")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
	// var responseObject Response
	// json.Unmarshal(bodyBytes, &responseObject)
		fmt.Print(url)
		fmt.Print("\n")
	fmt.Printf("API Response as struct %+v\n", resp)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("\n")
	fmt.Printf("API Response as struct %+v\n", resp.Body.body)
}
