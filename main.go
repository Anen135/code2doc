package main

import (
	"code2doc/internal/llm"
	"code2doc/internal/logger"
	"code2doc/internal/scanner.go"
	"code2doc/internal/utils"
	"fmt"
	"os"
)

func main() {
	utils.LoadEnv(".env")
	logger.Init()
	defer logger.Close()

	rootPath := "."
	apiKey := os.Getenv("OPENROUTER_KEY")
	model := os.Getenv("LLM_MODEL")
	if apiKey == "" || model == "" {
		fmt.Println("Please set OPENROUTER_KEY and LLM_MODEL environment variables")
		logger.Log("ERROR", "Missing environment variables")
		return
	}

	logger.Log("INFO", fmt.Sprintf("API Key configured, model: %s", model))

	llmClient := &llm.Client{
		APIKey: apiKey,
		Model:  model,
	}

	fmt.Println("Scanning...")
	logger.Log("INFO", "Starting code scan")

	files, err := scanner.Scan(rootPath, ".go")
	if err != nil {
		fmt.Printf("Failed to scan: %v\n", err)
		logger.Log("ERROR", fmt.Sprintf("Scan failed: %v", err))
		return
	}

	logger.Log("INFO", fmt.Sprintf("Found %d files to process", len(files)))

	for _, file := range files {
		fmt.Printf("Processing: %s\n", file.Path)
		logger.Log("INFO", fmt.Sprintf("Processing file: %s", file.Path))
		summary, err := scanner.AnalyzeCode(file.Path, file.Content)
		if err != nil {
			fmt.Printf("Failed to analyze code: %v\n", err)
			logger.Log("ERROR", fmt.Sprintf("Failed to analyze code: %v", err))
			continue
		}
		prompt := fmt.Sprintf("Write the technical documentation for this code. "+
			"Describe the basic structures and functions:\n\n%s", summary)

		doc, err := llmClient.Ask(prompt)
		if err != nil {
			fmt.Printf("Failed to generate documentation for %s: %v\n", file.Path, err)
			logger.Log("ERROR", fmt.Sprintf("Failed to generate doc for %s: %v", file.Path, err))
			continue
		}

		docPath := file.Path + ".md"
		os.WriteFile(docPath, []byte(doc), 0644)
		logger.Log("INFO", fmt.Sprintf("Documentation saved: %s", docPath))
		fmt.Printf("Documentation saved: %s\n", docPath)
	}

	logger.Log("INFO", "Documentation generation complete")
}
