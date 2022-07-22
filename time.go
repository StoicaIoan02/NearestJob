package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

// model: https://dev.to/billylkc/parse-json-api-response-in-go-10ng

func GetJsonBody(URL string) []byte {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	// fmt.Println(string(body))              // print json text
	return body
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Foo struct {
	Bar string
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

type Rezult struct {
	Distance struct {
		Text  string `json:"text"`
		Value int    `json:"value"`
	} `json:"distance"`
	Duration struct {
		Text  string `json:"text"`
		Value int    `json:"value"`
	} `json:"duration"`
	Status string `json:"status"`
}

// Generated go struct with https://mholt.github.io/json-to-go/
type Response struct {
	DestinationAddresses []string `json:"destination_addresses"`
	OriginAddresses      []string `json:"origin_addresses"`
	Rows                 []struct {
		Elements []Rezult `json:"elements"`
	} `json:"rows"`
	Status string `json:"status"`
}

func TimeResponse(origin_addresses, destination_addresses string) Rezult {
	origin_addresses = url.PathEscape(origin_addresses)
	destination_addresses = url.PathEscape(destination_addresses)
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?departure_time=1658725200&avoid=highways&destinations=%s&origins=%s&units=metric&language=ro&mode=transit&region=en&traffic_model=pessimistic&transit_mode=bus|train|tram|subway&transit_routing_preference=less_walking&key=%s", destination_addresses, origin_addresses, os.Getenv("GOOGLE_API_KEY"))
	body := GetJsonBody(url)

	// snippet only
	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	for _, rec := range result.Rows {
		for _, rec := range rec.Elements {
			return rec // Return first rezult
		}
	}

	return Rezult{}
}
