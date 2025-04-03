package obsidian

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/huh"
	lg "github.com/charmbracelet/log"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Obsidian struct {
	log *lg.Logger
	vaultPath string
	apiKey string
}

func New(log *lg.Logger, vaultPath string, apiKey string) *Obsidian {
	return &Obsidian{
		log: log,
		vaultPath: vaultPath,
		apiKey: apiKey,
	}
}

func (o *Obsidian) Picker() (string, error) {
	selected := ""

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your option").
				Options(
					huh.NewOption("Create Weekly Template", "CreateWeeklyTemplate"),
					huh.NewOption("Create Daily Template", "CreateDailyTemplate"),
				).
				Value(&selected),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	switch selected {
	case "CreateWeeklyTemplate":
		o.log.Info("Creating weekly template...")
		return o.processWeeklyTemplate()
	case "CreateDailyTemplate":
		return "Creating daily template...", nil
	}

	return "", nil
}

func (o *Obsidian) processWeeklyTemplate() (string, error) {
	if o.vaultPath == "" {
		return "", fmt.Errorf("Obsidian vault path is not set")
	}
	
	templatePath := filepath.Join(o.vaultPath, "_Meta", "Templates", "Week Template")
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %w", err)
	}
	
	if o.apiKey == "" {
		return "", fmt.Errorf("API key is not set")
	}
	
	client := openai.NewClient(option.WithAPIKey(o.apiKey))
	
	today := time.Now().Format("02/01/2006")
	systemPrompt := fmt.Sprintf(`
- Hoje é dia %s
- Compile o template de rotinas em um checklist :
- Compile apenas a semana atual
- Use o formato - [ ] para checklist
- Mantenha os textos exatamente como estão no template
- Formate as datas como 'Dia dd/mm' (ex: 'Segunda 01/04')
- Agrupe por semanas numeradas com seus intervalos (ex: Semana 1 (1-6))
- Mantenha a ordem cronológica dentro de cada semana
- Repita os eventos que acontecem no mesmo dia
`, today)
	
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemPrompt),
		openai.UserMessage(string(templateContent)),
	}
	
	resp, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    "gpt-4",
	})
	o.log.Info("Template processed with OpenAI")
	if err != nil {
		return "", fmt.Errorf("failed to process template with OpenAI: %w", err)
	}
	
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}
	
	return resp.Choices[0].Message.Content, nil
}
