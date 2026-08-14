package bubble_modal

import "github.com/charmbracelet/lipgloss"

var FileIcons = map[string]IconStyle{
	// --- Folders ---
	"folder":      {"📁", lipgloss.NewStyle().Foreground(lipgloss.Color("#79B8FF"))},
	"folder_open": {"📂", lipgloss.NewStyle().Foreground(lipgloss.Color("#79B8FF"))},

	// --- Archives / Zips ---
	".zip": {"📦", lipgloss.NewStyle().Foreground(lipgloss.Color("#E8B04B"))},
	".tar": {"📦", lipgloss.NewStyle().Foreground(lipgloss.Color("#E8B04B"))},
	".gz":  {"📦", lipgloss.NewStyle().Foreground(lipgloss.Color("#E8B04B"))},
	".rar": {"📦", lipgloss.NewStyle().Foreground(lipgloss.Color("#E8B04B"))},
	".7z":  {"📦", lipgloss.NewStyle().Foreground(lipgloss.Color("#E8B04B"))},
	".xz":  {"📦", lipgloss.NewStyle().Foreground(lipgloss.Color("#E8B04B"))},

	// --- Programming Languages ---
	".go":    {"🐹", lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8"))}, // go gopher
	".py":    {"🐍", lipgloss.NewStyle().Foreground(lipgloss.Color("#3572A5"))}, // python
	".js":    {"📜", lipgloss.NewStyle().Foreground(lipgloss.Color("#F1E05A"))},
	".ts":    {"📘", lipgloss.NewStyle().Foreground(lipgloss.Color("#3178C6"))},
	".jsx":   {"⚛️", lipgloss.NewStyle().Foreground(lipgloss.Color("#61DAFB"))},
	".tsx":   {"⚛️", lipgloss.NewStyle().Foreground(lipgloss.Color("#61DAFB"))},
	".rs":    {"🦀", lipgloss.NewStyle().Foreground(lipgloss.Color("#DEA584"))}, // rust crab
	".java":  {"☕", lipgloss.NewStyle().Foreground(lipgloss.Color("#B07219"))}, // coffee
	".c":     {"🔧", lipgloss.NewStyle().Foreground(lipgloss.Color("#599EFF"))},
	".cpp":   {"🔧", lipgloss.NewStyle().Foreground(lipgloss.Color("#F34B7D"))},
	".h":     {"📐", lipgloss.NewStyle().Foreground(lipgloss.Color("#A074C4"))},
	".cs":    {"🔷", lipgloss.NewStyle().Foreground(lipgloss.Color("#68217A"))},
	".php":   {"🐘", lipgloss.NewStyle().Foreground(lipgloss.Color("#8892BF"))}, // elephant
	".rb":    {"💎", lipgloss.NewStyle().Foreground(lipgloss.Color("#CC342D"))}, // ruby gem
	".swift": {"🦅", lipgloss.NewStyle().Foreground(lipgloss.Color("#F05138"))},
	".kt":    {"🟣", lipgloss.NewStyle().Foreground(lipgloss.Color("#A97BFF"))},
	".lua":   {"🌙", lipgloss.NewStyle().Foreground(lipgloss.Color("#000080"))}, // moon
	".sh":    {"🐚", lipgloss.NewStyle().Foreground(lipgloss.Color("#89E051"))}, // shell

	// --- Web / Markup ---
	".html": {"🌐", lipgloss.NewStyle().Foreground(lipgloss.Color("#E44D26"))},
	".css":  {"🎨", lipgloss.NewStyle().Foreground(lipgloss.Color("#563D7C"))},
	".scss": {"🎨", lipgloss.NewStyle().Foreground(lipgloss.Color("#C6538C"))},
	".md":   {"📝", lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))},
	".json": {"🧾", lipgloss.NewStyle().Foreground(lipgloss.Color("#CBCB41"))},
	".yaml": {"⚙️", lipgloss.NewStyle().Foreground(lipgloss.Color("#CB171E"))},
	".yml":  {"⚙️", lipgloss.NewStyle().Foreground(lipgloss.Color("#CB171E"))},
	".toml": {"⚙️", lipgloss.NewStyle().Foreground(lipgloss.Color("#9C4221"))},
	".xml":  {"📰", lipgloss.NewStyle().Foreground(lipgloss.Color("#E37933"))},

	// --- Docs / Office ---
	".pdf":  {"📕", lipgloss.NewStyle().Foreground(lipgloss.Color("#E34F26"))},
	".doc":  {"📄", lipgloss.NewStyle().Foreground(lipgloss.Color("#2B579A"))},
	".docx": {"📄", lipgloss.NewStyle().Foreground(lipgloss.Color("#2B579A"))},
	".xls":  {"📊", lipgloss.NewStyle().Foreground(lipgloss.Color("#217346"))},
	".xlsx": {"📊", lipgloss.NewStyle().Foreground(lipgloss.Color("#217346"))},
	".ppt":  {"📈", lipgloss.NewStyle().Foreground(lipgloss.Color("#D24726"))},
	".pptx": {"📈", lipgloss.NewStyle().Foreground(lipgloss.Color("#D24726"))},
	".txt":  {"📄", lipgloss.NewStyle().Foreground(lipgloss.Color("#89E051"))},

	// --- Media ---
	".png":  {"🖼️", lipgloss.NewStyle().Foreground(lipgloss.Color("#A074C4"))},
	".jpg":  {"🖼️", lipgloss.NewStyle().Foreground(lipgloss.Color("#A074C4"))},
	".jpeg": {"🖼️", lipgloss.NewStyle().Foreground(lipgloss.Color("#A074C4"))},
	".gif":  {"🖼️", lipgloss.NewStyle().Foreground(lipgloss.Color("#A074C4"))},
	".svg":  {"🎭", lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB13B"))},
	".mp3":  {"🎵", lipgloss.NewStyle().Foreground(lipgloss.Color("#E8B04B"))},
	".mp4":  {"🎬", lipgloss.NewStyle().Foreground(lipgloss.Color("#E8B04B"))},
	".mov":  {"🎬", lipgloss.NewStyle().Foreground(lipgloss.Color("#E8B04B"))},

	// --- Config / Misc ---
	".env":       {"🔐", lipgloss.NewStyle().Foreground(lipgloss.Color("#FAF743"))},
	".gitignore": {"🚫", lipgloss.NewStyle().Foreground(lipgloss.Color("#F14E32"))},
	".lock":      {"🔒", lipgloss.NewStyle().Foreground(lipgloss.Color("#BBBBBB"))},
	".sql":       {"🗃️", lipgloss.NewStyle().Foreground(lipgloss.Color("#DAD8D8"))},
	".db":        {"🗃️", lipgloss.NewStyle().Foreground(lipgloss.Color("#DAD8D8"))},
	".exe":       {"⚡", lipgloss.NewStyle().Foreground(lipgloss.Color("#83CD29"))},

	// --- Fallback ---
	"default": {"📄", lipgloss.NewStyle().Foreground(lipgloss.Color("#7A7A7A"))},
}
