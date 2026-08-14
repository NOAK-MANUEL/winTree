package bubble_modal

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	globarvar "winTree/globarVar"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func Lookup(ext string) IconStyle {
	if style, ok := FileIcons[ext]; ok {
		return style
	}
	return FileIcons["default"]
}

func (m Model) View() string {
	var output strings.Builder

	selectedStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205"))

	end := min(m.offset+m.visibleItems, len(m.items))

	for index := m.offset; index < end; index++ {
		item := m.items[index]
		var icon IconStyle

		output.WriteString(strings.Repeat("  ", item.depth))
		if item.isFolder {
			if item.isOpen {

				icon = Lookup("folder_open")
			} else {

				icon = Lookup("folder")
			}
			if m.selected == index {

				output.WriteString(selectedStyle.Render("> " + icon.Icon + " " + item.name))
				output.WriteString("\n")
			} else {
				output.WriteString(icon.Color.Render(icon.Icon + " " + item.name))
				output.WriteString("\n")

			}
		} else {
			icon = Lookup(filepath.Ext(item.path))
			if m.selected == index {

				output.WriteString(selectedStyle.Render("> " + icon.Icon + " " + item.name))
				output.WriteString("\n")
			} else {
				output.WriteString(icon.Color.Render(icon.Icon + " " + item.name))
				output.WriteString("\n")

			}

		}

	}

	return output.String()
}

func (m Model) Init() tea.Cmd {
	render := lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	println(render.Render("These are the following actions\nenter: to show the content of a folder\no: to open file/folder\nq: to quite program\n Happy exploring!\n"))
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.selected > 0 {
				m.selected--
				if m.selected < m.offset {
					m.offset--
				}
			}

		case "down":
			if m.selected < len(m.items)-1 {
				m.selected++
				if m.selected >= m.visibleItems+m.offset {
					m.offset++
				}
			}
		case "enter":
			item := &m.items[m.selected]
			if item.isFolder {
				if item.isOpen {
					index := m.selected + 1

					for index < len(m.items) &&
						item.depth < m.items[index].depth {

						copy(m.items[index:], m.items[index+1:])
						m.items = m.items[:len(m.items)-1]
					}

					item.isOpen = false

				} else {
					if m.selected >= m.visibleItems+m.offset {
						m.offset++
					}
					item.isOpen = true
					parseTree(item.path, item.depth, &m, m.selected, true, globarvar.JustDir, false)
				}
			}

		case "o":
			item := &m.items[m.selected]
			if globarvar.ON_CD_COMMAND {
				m.selectedPath = item.path
				return m, tea.Quit
			} else {

				switch runtime.GOOS {
				case "windows":
					exec.Command("cmd", "/c", "start", "", item.path).Start()
				case "darwin":
					exec.Command("open", item.path).Start()
				case "linux":
					exec.Command("xdg-open", item.path).Start()
				}

			}

		case "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.visibleItems = msg.Height - 1

	}
	return m, nil
}
