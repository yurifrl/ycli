package obsidian

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	lg "github.com/charmbracelet/log"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/openai"
)

type ModelProvider string

const (
	OpenAI    ModelProvider = "openai"
	Anthropic ModelProvider = "anthropic"
)

type Obsidian struct {
	log          *lg.Logger
	vaultPath    string
	apiKey       string
	modelName    string
	provider     ModelProvider
}

func New(log *lg.Logger, vaultPath string, apiKey string, modelName string, provider ModelProvider) *Obsidian {
	return &Obsidian{
		log:       log,
		vaultPath: vaultPath,
		apiKey:    apiKey,
		modelName: modelName,
		provider:  provider,
	}
}

func (o *Obsidian) WeeklyTemplate() (string, error) {
	if o.vaultPath == "" {
		return "", fmt.Errorf("Obsidian vault path is not set")
	}
	
	homeDir, _ := os.UserHomeDir()
	templatePath := filepath.Join(homeDir, "Obsidian", "_Meta", "Templates", "Week Template.md")
	o.log.Debug("Using template at", "path", templatePath)
	
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %w", err)
	}
	
	if o.apiKey == "" {
		return "", fmt.Errorf("API key is not set")
	}
	
	var llm llms.Model
	var llmErr error

	switch o.provider {
	case OpenAI:
		llm, llmErr = openai.New(openai.WithToken(o.apiKey), openai.WithModel(o.modelName))
	case Anthropic:
		llm, llmErr = anthropic.New(anthropic.WithToken(o.apiKey), anthropic.WithModel(o.modelName))
	default:
		return "", fmt.Errorf("unsupported model provider: %s", o.provider)
	}

	if llmErr != nil {
		return "", fmt.Errorf("failed to initialize LLM: %w", llmErr)
	}
	
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
	
	prompt := systemPrompt + "\n\n" + string(templateContent)
	
	result, err := llm.Call(context.Background(), prompt)
	o.log.Info("Template processed with LLM")
	if err != nil {
		return "", fmt.Errorf("failed to process template with LLM: %w", err)
	}
	
	return result, nil
}
