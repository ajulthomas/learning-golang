package main

import "fmt"

type JSON map[string]string

func main() {
	var m = make(map[string]string)
	m["name"] = "My Name"
	address := map[string]string{
		"house":   "House name/ House number",
		"street":  "My street name",
		"country": "My Counry",
		"state":   "My state",
		"city":    "My City",
		"pincode": "pincode",
	}
	fmt.Println("Looping through Map")
	fmt.Println("====================")
	for key, value := range address {
		fmt.Printf("key: %v, value: %v\n\n", key, value)
	}
	apiResponse := JSON{
		"apiName":    "test",
		"status":     "SUCCESS",
		"statusCode": "200",
	}
	fmt.Printf("API response ===> %v", apiResponse)
}
