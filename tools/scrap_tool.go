package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/nlpodyssey/openai-agents-go/agents"
)

func scrapeWebsite(_ context.Context, url string) (string, error) {
	var content strings.Builder

	// Create a new collector
	c := colly.NewCollector()

	// Find the body tag and print its text
	c.OnHTML("body", func(e *colly.HTMLElement) {
		content.WriteString(e.Text)
	})

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		return "", fmt.Errorf("failed to visit URL: %w", err)
	}

	return content.String(), nil
}

var ScrapeTool = agents.NewFunctionTool(
	"scrape_website",
	"Scrapes the content from a specified website URL. Input should be the URL as a string.",
	scrapeWebsite,
)
