package quote

import (
	"encoding/json"
	"net/http"
	"time"
	"vortex-engine/logger"
)

type ZenQuoteResponse struct {
	Text   string `json:"q"`
	Author string `json:"a"`
}

type Quote struct {
	Text   string
	Author string
}

func GetZenQuote() Quote {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Get("https://zenquotes.io/api/random")
	if err != nil {
		logger.Error.Printf("Error fetching quote: %v", err)
		return Quote{}
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logger.Error.Printf("Non-OK HTTP status: %s", response.Status)
		return Quote{}
	}

	logger.Info.Println("Response Status: " + response.Status)

	var data []ZenQuoteResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		logger.Error.Printf("Error decoding quote response: %v", err)
		return Quote{}
	}

	quote := data[0]

	logger.Info.Printf("Fetched Quote: %s - %s", quote.Author, quote.Text)

	return Quote{
		Text:   quote.Text,
		Author: quote.Author,
	}
}
