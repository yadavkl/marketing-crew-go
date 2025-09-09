package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nlpodyssey/openai-agents-go/agents"
)

func readFile(_ context.Context, filePath string) (string, error) {
	// Prepend a base path to ensure security and restrict file access.
	// You should configure this based on your application's needs.
	safePath := filepath.Join("resources/drafts", filePath)

	content, err := os.ReadFile(safePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file '%s': %w", safePath, err)
	}

	return string(content), nil
}

var FileReadTool = agents.NewFunctionTool(
	"read_file",
	"Reads the content of a specified file. Takes the file path as input.",
	readFile,
)
