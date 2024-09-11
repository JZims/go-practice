package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url string = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	req.Header = http.Header{
		"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"},
		"Accept":     {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
	}

	resp, err := client.Do(req)
	checkError(err)

	fmt.Printf("Reponse status code is: %v\n", resp.StatusCode)
	// fmt.Printf("Request Header is: %v\n", req.Header.Get("Host"))
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)

	// fmt.Print(content)

	tours := ToursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Converts incoming JSON data into a Tours map
func ToursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
