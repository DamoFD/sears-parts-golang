package search

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Model struct {
	ID        string   `json:"id"`
	Number    string   `json:"number"`
	Title     string   `json:"title"`
	Images    []string `json:"images"`
	Brand     Brand    `json:"brand"`
	PartCount int      `json:"partCount"`
}

type Brand struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Part struct {
	ID                string         `json:"id"`
	Number            string         `json:"number"`
	Title             string         `json:"title"`
	TopSoldPosition   int            `json:"topSoldPosition"`
	Media             Media          `json:"media"`
	Subscribable      bool           `json:"subscribable"`
	Subscriptions     []Subscription `json:"subscriptions"`
	Division          Division       `json:"division"`
	SourceID          string         `json:"sourceId"`
	Restrictions      []Restriction  `json:"restrictions"`
	Pricing           Pricing        `json:"pricing"`
	SubstitutedByList struct {
		Parts []Part `json:"parts"`
	} `json:"substitutedByList"`
}

type Media struct {
	Image struct {
		Urls []string `json:"urls"`
	} `json:"image"`
}

type Subscription struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	Interval    string `json:"interval"`
}

type Division struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Restriction struct {
	ID string `json:"id"`
}

type Pricing struct {
	StrikeThroughPriceType string  `json:"strikeThroughPriceType"`
	Sell                   float64 `json:"sell"`
	List                   float64 `json:"list"`
	AvailabilityInfo       struct {
		Status    string `json:"status"`
		Inventory int    `json:"inventory"`
	} `json:"availabilityInfo"`
}

type Response struct {
	Data struct {
		ModelSearch struct {
			TotalCount int     `json:"totalCount"`
			Models     []Model `json:"models"`
		} `json:"modelSearch"`
	} `json:"data"`
}

type PartResponse struct {
	Data struct {
		PartSearch struct {
			TotalCount int    `json:"totalCount"`
			Parts      []Part `json:"parts"`
		} `json:"partSearch"`
	} `json:"data"`
}

type Variables struct {
	Query       string      `json:"q"`
	Page        Page        `json:"page"`
	PriceFilter PriceFilter `json:"priceFilter"`
	Filters     []string    `json:"filters"`
}

type PartVariables struct {
	Query                   string   `json:"q"`
	Page                    Page     `json:"page"`
	Orders                  []Order  `json:"orders"`
	Filters                 []Filter `json:"filters"`
	SubstitutedByListFilter []Filter `json:"substitutedByListFilter"`
}

type Page struct {
	From int `json:"from"`
	Size int `json:"size"`
}

type PriceFilter struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []string `json:"values"`
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

type PartURL struct {
	BaseURL       string
	OperationName string        `json:"operationName"`
	Variables     PartVariables `json:"variables"`
	Extensions    Extensions    `json:"extensions"`
}

func (u PartURL) BuildURL() string {
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

func searchPart(query string, from int, size int) []Part {
	page := Page{
		From: from,
		Size: size,
	}

	orders := []Order{
		{"SCORE", "DESC"},
		{"SYNONYMS", "ASC"},
		{"SELLABLE", "DESC"},
		{"RANK", "DESC"},
		{"AVAILABILITY", "DESC"},
	}

	filters := []Filter{
		{
			Name:   "RESTRICTION",
			Type:   "NOT",
			Values: []string{"31", "49", "60", "12", "17"},
		},
		{
			Name:   "SELLABLE",
			Type:   "MATCH",
			Values: []string{"true"},
		},
	}

	substitutedByListFilter := []Filter{
		{
			Name:   "AVAILABILITY",
			Type:   "MATCH",
			Values: []string{"PIA", "BORD", "NLO"},
		},
		{
			Name:   "PRICE",
			Type:   "RANGE",
			Values: []string{">1"},
		},
		{
			Name:   "SELLABLE",
			Type:   "MATCH",
			Values: []string{"true"},
		},
		{
			Name:   "RESTRICTION",
			Type:   "NOT",
			Values: []string{"31", "49", "60", "12", "17"},
		},
	}

	variables := PartVariables{
		Query:                   query,
		Page:                    page,
		Orders:                  orders,
		Filters:                 filters,
		SubstitutedByListFilter: substitutedByListFilter,
	}

	extensions := Extensions{
		PersistedQuery: PersistedQuery{
			Version:    1,
			Sha256Hash: "9c025907173f3568d9019e250fceb5f0f92bf26a63dce5ca256f77a631bd4788",
		},
	}

	urlStruct := PartURL{
		BaseURL:       "https://catalog.searspartsdirect.com",
		OperationName: "partSearch",
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
	var response PartResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Access the parts
	parts := response.Data.PartSearch.Parts

	return parts
}

func searchModel(query string, from int, size int) []Model {
	page := Page{
		From: from,
		Size: size,
	}

	priceFilter := PriceFilter{
		Name:   "PRICE",
		Type:   "RANGE",
		Values: []string{">1"},
	}

	variables := Variables{
		Query:       query,
		Page:        page,
		PriceFilter: priceFilter,
		Filters:     []string{},
	}

	extensions := Extensions{
		PersistedQuery: PersistedQuery{
			Version:    1,
			Sha256Hash: "eadec1e2e8cbfc0b7c3a4b87de9af960a4aa14df1cbea852fdff503e9740ad67",
		},
	}

	urlStruct := URL{
		BaseURL:       "https://catalog.searspartsdirect.com",
		OperationName: "modelSearch",
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

	// Access the models
	models := response.Data.ModelSearch.Models

	return models
}

func SearchParts(query string, from int, size int) []Part {
	return searchPart(query, from, size)
}

func SearchModels(query string, from int, size int) []Model {
	return searchModel(query, from, size)
}
