package bubble_modal

import "github.com/charmbracelet/lipgloss"

type Item struct {
	name  string
	path  string
	depth int

	isFolder bool
	isOpen   bool
}

type Model struct {
	items        []Item
	selected     int
	offset       int
	visibleItems int
	selectedPath string
}

type IconStyle struct {
	Icon  string
	Color lipgloss.Style
}
