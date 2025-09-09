package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nlpodyssey/openai-agents-go/agents"
)

type FileArgs struct {
	FilePath string `json:"file_path" jsonschema_description:"file path for writing"`
	Content  string `json:"content" jsonschema_description:"contnet of the file to be written in file path"`
}

func writeFile(_ context.Context, fileargs FileArgs) (string, error) {
	// Create directories if they don't exist
	dir := filepath.Dir(fileargs.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// Write the content to the file. os.WriteFile handles creating or truncating.
	if err := os.WriteFile(fileargs.FilePath, []byte(fileargs.Content), 0644); err != nil {
		return "", fmt.Errorf("failed to write to file '%s': %w", fileargs.FilePath, err)
	}

	return fmt.Sprintf("Successfully wrote content to %s", fileargs.FilePath), nil
}

var FileWriterTool = agents.NewFunctionTool(
	"write_file",
	"Writes content to a specified file. Takes the file path and content as input.",
	writeFile,
)
