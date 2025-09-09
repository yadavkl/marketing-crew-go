package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nlpodyssey/openai-agents-go/agents"
)

func readDirectory(_ context.Context, path string) (string, error) {
	// Prepend a base path to ensure security and prevent arbitrary file access.
	// You should configure this based on your application's needs.
	safePath := filepath.Join("resources/drafts", path)

	entries, err := os.ReadDir(safePath)
	if err != nil {
		return "", fmt.Errorf("failed to read directory '%s': %w", safePath, err)
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Contents of directory '%s':\n", safePath))
	for _, entry := range entries {
		entryType := "File"
		if entry.IsDir() {
			entryType = "Directory"
		}
		builder.WriteString(fmt.Sprintf("- %s: %s\n", entryType, entry.Name()))
	}
	return builder.String(), nil
}

var DirectoryReadTool = agents.NewFunctionTool(
	"read_directory",
	"Reads the entries (files and subdirectories) within a specified directory path. Input should be the path as a string.",
	readDirectory,
)
