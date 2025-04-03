package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	lg "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/yurifrl/cli/pkg/obsidian"
	"github.com/yurifrl/cli/pkg/placeholder"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "",
	Example: "",
	Long:    "",
	Run:     run,
	Args:    cobra.MaximumNArgs(1),
}
var selected string

func init() {
	rootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	lg.Debug("Hy 🍪")

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your option").
				Options(
					huh.NewOption("Table Data", "FunctionA"),
					huh.NewOption("Simple Message", "FunctionB"),
					huh.NewOption("Obsidian Weekly Template", "ObsidianWeeklyTemplate"),
				).
				Value(&selected),
		),
	)

	err := form.Run()
	if err != nil {
		lg.Fatal(err)
	}

	core := placeholder.Core{}
	
	obsidian := obsidian.New(
		lg.Default(),
		appConfig.ObsidianVaultPath,
		appConfig.AnthropicAPIKey,
		"claude-3-5-sonnet-20240620",
		obsidian.Anthropic,
	)

	switch selected {
	case "FunctionA":
		t := core.FunctionA()

		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false)
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false)
		t.SetStyles(s)

		m := model{table: t}
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	case "FunctionB":
		defaultFile := "~/Downloads/8\\ Aug\\ 13.50.48\\ System\\ Audio_Microphone.txt"
		message, err := core.SummarizeFile(defaultFile, "")
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(message)
	case "ObsidianWeeklyTemplate":
		message, err := obsidian.WeeklyTemplate()
		if err != nil {
			lg.Fatal(err)
		}
		fmt.Println(message)
	default:
		fmt.Println("Invalid selection")
	}
}