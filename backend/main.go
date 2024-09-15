package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Part struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Response struct {
	Data struct {
		Model struct {
			ID       string `json:"id"`
			Title    string `json:"title"`
			Number   string `json:"number"`
			HasParts bool   `json:"hasParts"`
			Parts    struct {
				TotalCount int    `json:"totalCount"`
				Parts      []Part `json:"parts"`
			} `json:"parts"`
		} `json:"model"`
	} `json:"data"`
}

type Order struct {
	Name  string `json:"name"`
	Order string `json:"order"`
}

type Filter struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

type ParentFilter struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type Variables struct {
	ID           string         `json:"id"`
	Orders       []Order        `json:"orders"`
	Filters      []Filter       `json:"filters"`
	From         int            `json:"from"`
	Size         int            `json:"size"`
	ParentFilter []ParentFilter `json:"parentFilter"`
}

type PersistedQuery struct {
	Version    int    `json:"version"`
	Sha256Hash string `json:"sha256Hash"`
}

type Extensions struct {
	PersistedQuery PersistedQuery `json:"persistedQuery"`
}

type URL struct {
	BaseURL       string
	OperationName string     `json:"operationName"`
	Variables     Variables  `json:"variables"`
	Extensions    Extensions `json:"extensions"`
}

func (u URL) BuildURL() string {
	variablesJSON, err := json.Marshal(u.Variables)
	if err != nil {
		log.Fatal(err)
	}
	extensionsJSON, err := json.Marshal(u.Extensions)
	if err != nil {
		log.Fatal(err)
	}

	variablesEncoded := url.QueryEscape(string(variablesJSON))
	extensionsEncoded := url.QueryEscape(string(extensionsJSON))

	return fmt.Sprintf("%s/graphql?operationName=%s&variables=%s&extensions=%s",
		u.BaseURL,
		url.QueryEscape(u.OperationName),
		variablesEncoded,
		extensionsEncoded,
	)
}

func getParts(id string, from int, size int) []Part {
	orders := []Order{
		{"SELLABLE", "DESC"},
		{"RANK", "DESC"},
		{"AVAILABILITY", "DESC"},
	}

	filters := []Filter{
		{"RESTRICTION", "NOT", []string{"31", "49", "60", "12", "17"}},
	}

	parentFilter := []ParentFilter{
		{"ISMAINCATEGORY", []string{"true"}},
	}

	variables := Variables{
		ID:           id,
		Orders:       orders,
		Filters:      filters,
		From:         from,
		Size:         size,
		ParentFilter: parentFilter,
	}

	extensions := Extensions{
		PersistedQuery: PersistedQuery{
			Version:    1,
			Sha256Hash: "ec724acca5636f88b79dd75604b6812d84e8d3e091a3e614aac285ee9a836e39",
		},
	}

	urlStruct := URL{
		BaseURL:       "https://catalog.searspartsdirect.com",
		OperationName: "getModel",
		Variables:     variables,
		Extensions:    extensions,
	}

	url := urlStruct.BuildURL()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Del("Accept-Encoding")

	// Add necessary headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Origin", "https://www.searspartsdirect.com")
	req.Header.Set("Referer", "https://www.searspartsdirect.com/")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("X-Api-Key", "tV6bZZfhUh3MQmZggG6iq6LjfrZgQgcR26Yv86En")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the response body into a struct
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Access the parts
	parts := response.Data.Model.Parts.Parts

	return parts
}

func main() {

	// Get the parts
	parts := getParts("4tjp66n58g-003048", 15, 20)

	// Print the parts
	fmt.Println(parts)
}
