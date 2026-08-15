# 🌳 winTree

<p align="center">
  <strong>A fast, interactive terminal file explorer built with Go.</strong>
</p>

<p align="center">
  Explore your filesystem in a tree-based interface directly from your terminal.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Bubble%20Tea-TUI-000000?style=for-the-badge" alt="Bubble Tea">
  <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-blue?style=for-the-badge" alt="Platform">
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License">
</p>

---




### 🎥 Video Demo

<p align="center"> <img src="./demo.gif" alt="winTree demo" width="800"> </p>
---

## ✨ Features

- 🌳 Interactive tree-based filesystem explorer
- 📁 Expand and collapse directories
- 📄 File-type icons with color-coded highlighting
- ⌨️ Fast keyboard-driven navigation
- 📂 Open files and folders directly from the tree
- 📐 Automatically adapts to terminal size
- 🚀 Single native executable — no runtime dependencies
- 🐚 Simple shell integration
- 🔎 Directory-based navigation from any path
- 💻 Runs entirely inside your terminal

---

## 📦 Installation

### Build from source

Make sure you have [Go 1.22+](https://go.dev/dl/) installed.

```bash
git clone https://github.com/NOAK-MANUEL/winTree.git
cd winTree
go build -o wintree .
```

Then move the binary somewhere on your `PATH`:

```bash
# Linux / macOS
sudo mv wintree /usr/local/bin/

# Windows (PowerShell, as Administrator)
Move-Item .\wintree.exe C:\Windows\System32\
```



---

## 🚀 Usage

winTree is invoked from your terminal with a small set of commands:

| Command                | Description                                      |
|-------------------------|--------------------------------------------------|
| `wintree {path}`        | Open the tree explorer at the given path         |
| `wintree cd {path}`     | Navigate directly into a directory                |
| `wintree install`       | Install/register the `wintree-cd` binary on your system |

**Examples:**

```bash
# Open the explorer in the current directory
wintree .

# Open the explorer at a specific path
wintree ~/projects

# Jump straight into a folder
wintree cd ~/projects/winTree
```

---

## ⌨️ Keybindings

| Key      | Action                                      |
|----------|----------------------------------------------|
| `o`      | Open the selected file or folder              |
| `Enter`  | Show the contents of the selected folder      |
| `q`      | Quit winTree                                  |

---

## 🛠️ Built With

- [Go](https://go.dev/) — core language
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — terminal UI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — terminal styling

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!
Feel free to check the [issues page](https://github.com/NOAK-MANUEL/winTree/issues) or open a pull request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a pull request

---

## 📄 License

This project is licensed under the **MIT License** — see the [LICENSE](./LICENSE) file for details.

---
