package bubble_modal

import (
	globarvar "winTree/globarVar"

	tea "github.com/charmbracelet/bubbletea"
)

func ExecuteBubble(path string) error {
	model := Model{items: []Item{}, selected: 0}

	parseTree(path, 0, &model, 0, false, globarvar.JustDir, globarvar.JustFile)
	p := tea.NewProgram(model)
	teaModel, err := p.Run()
	if err != nil {
		return err
	}
	m := teaModel.(Model)
	println(m.selectedPath)
	return nil
}
