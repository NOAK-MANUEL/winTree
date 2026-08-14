package bubble_modal

import (
	"os"
	"path/filepath"
)

func parseTree(path string, depth int, m *Model, selected int, child bool, justDir bool, justFile bool) {
	data, err := os.ReadDir(path)
	if err != nil {
		println("Couldn't open folder", err)
	}
	for _, entry := range data {
		if child {
			if justDir && entry.IsDir() {

				selected++
				m.items = append(m.items, Item{})
				copy(m.items[selected+1:], m.items[selected:])
				m.items[selected] = Item{name: entry.Name(), isFolder: entry.IsDir(), path: filepath.Join(path, entry.Name()), depth: depth + 1}
			} else if !justFile {

				selected++
				m.items = append(m.items, Item{})
				copy(m.items[selected+1:], m.items[selected:])
				m.items[selected] = Item{name: entry.Name(), isFolder: entry.IsDir(), path: filepath.Join(path, entry.Name()), depth: depth + 1}

			}
		} else {
			if justDir && entry.IsDir() {

				m.items = append(m.items, Item{name: entry.Name(), isFolder: entry.IsDir(), path: filepath.Join(path, entry.Name()), depth: depth})
			} else if justFile && !entry.IsDir() {
				m.items = append(m.items, Item{name: entry.Name(), isFolder: entry.IsDir(), path: filepath.Join(path, entry.Name()), depth: depth})

			} else {
				m.items = append(m.items, Item{name: entry.Name(), isFolder: entry.IsDir(), path: filepath.Join(path, entry.Name()), depth: depth})

			}
		}

	}
}
