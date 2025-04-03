package placeholder

import (
	"github.com/charmbracelet/bubbles/table"
	lg "github.com/charmbracelet/log"
)

type Core struct{
	DefaultFilePath string
}

func (c *Core) FunctionA() table.Model {
	lg.SetLevel(lg.DebugLevel)
	lg.Print("Function A executing!")

	// Create columns
	columns := []table.Column{
		{Title: "Table ID", Width: 20},
		{Title: "Status", Width: 10},
		{Title: "Seats", Width: 10},
		{Title: "Reservation", Width: 20},
		{Title: "Special Requests", Width: 30},
	}

	// Create fake rows
	rows := []table.Row{
		{"table-1", "Occupied", "4", "Smith Family", "Vegan options\nWindow seat"},
		{"table-2", "Available", "2", "None", "None"},
		{"table-3", "Reserved", "6", "Johnson Group", "Birthday celebration"},
		{"table-4", "Occupied", "2", "Miller", "Anniversary"},
		{"table-5", "Cleaning", "4", "None", "None"},
	}

	// Create table model
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(30),
	)

	return t
}
