package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getRIDBData() ([]byte, error) {
	url := "https://ridb.recreation.gov/api/v1/facilities"

	// Make the HTTP GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	data, err := getRIDBData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	facility, err := facility.getFacilityByID(data, 2988)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(facility.Name) // print the name of the facility
}
