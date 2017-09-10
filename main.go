//https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"net/url"
)

type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

func main() {
	//phone := "14158586273"
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	//safePhone := url.QueryEscape(phone)

	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=1f0ca7242582b04da8a75c1c5e0bc753&number=14158586273")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

  fmt.Println("url ", url)
	fmt.Println("req ", req)

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	fmt.Println("resp ", resp.Body)

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("record ", record)

	fmt.Println("Phone No. = ", record.InternationalFormat)
	fmt.Println("Country   = ", record.CountryName)
	fmt.Println("Location  = ", record.Location)
	fmt.Println("Carrier   = ", record.Carrier)
	fmt.Println("LineType  = ", record.LineType)

}
