package placeholder

import (
	"fmt"
	"os"

	lg "github.com/charmbracelet/log"
)

func (c *Core) SummarizeFile(filePath, prompt string) (string, error) {
	lg.SetLevel(lg.DebugLevel)
	lg.Print("Function B executing!")

	// Use DefaultFilePath if available and filePath is empty
	if filePath == "" && c.DefaultFilePath != "" {
		filePath = c.DefaultFilePath
	}

	// Read the file content
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Prepare the content to be summarized along with the optional prompt
	fullPrompt := fmt.Sprintf("Summarize the following text:\n\n%s", string(fileContent))
	if prompt != "" {
		fullPrompt += fmt.Sprintf("\n\nAdditional instructions: %s", prompt)
	}

	// Get the OpenAI API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	// Create an OpenAI client
	// client := openai.NewClient(option.WithAPIKey(apiKey))

	// // Make a request to OpenAI API for summarization
	// response, err := client.Chat.Completions.Create(openai.ChatCompletionRequest{
	// 	Engine:    "davinci",
	// 	Prompt:    fullPrompt,
	// 	MaxTokens: 150,
	// })
	// if err != nil {
	// 	return "", err
	// }

	return "Hello", nil
	// Extract and return the summary
	// if len(response.Choices) > 0 {
	// 	return response.Choices[0].Text, nil
	// }
	// return "", fmt.Errorf("no response from OpenAI API")
}
