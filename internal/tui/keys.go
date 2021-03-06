package tui

import "github.com/charmbracelet/bubbles/key"

var (
	deleteKey      = key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "delete branch(es)"))
	selectKey      = key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "select branch"))
	selectAllKey   = key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "select all branches"))
	unselectAllKey = key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "unselect all branches"))
)
