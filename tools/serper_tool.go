package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/nlpodyssey/openai-agents-go/agents"
)

type SearchResults struct {
	Organic []struct {
		Title   string `json:"title"`
		Link    string `json:"link"`
		Snippet string `json:"snippet"`
	} `json:"organic"`
}

// searchOnSerper queries the Serper API with the given query.
func searchOnSerper(_ context.Context, query string) (string, error) {
	apiKey := os.Getenv("SERPER_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("SERPER_API_KEY environment variable not set")
	}

	url := fmt.Sprintf("https://serper.dev/search?q=%s", query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-API-KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Serper API error: %s", string(body))
	}

	var searchResults SearchResults
	if err := json.Unmarshal(body, &searchResults); err != nil {
		return "", err
	}

	// Format the results into a readable string for the agent.
	var resultBuilder strings.Builder
	for _, res := range searchResults.Organic {
		resultBuilder.WriteString(fmt.Sprintf("Title: %s\n", res.Title))
		resultBuilder.WriteString(fmt.Sprintf("Link: %s\n", res.Link))
		resultBuilder.WriteString(fmt.Sprintf("Snippet: %s\n\n", res.Snippet))
	}

	return resultBuilder.String(), nil
}

// 2. Define the tool
var SerperSearchTool = agents.NewFunctionTool(
	"serper_search",
	"Searches the web for information using the Serper API. Useful for finding up-to-date information.",
	searchOnSerper,
)
