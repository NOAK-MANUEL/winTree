package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	globarvar "winTree/globarVar"
)

func installShell(name string) error {
	var (
		file string
		dest string
	)

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	switch name {
	case "powershell":
		file = "shells/shell.ps1"
		dest = filepath.Join(
			home, "Documents", "PowerShell", "Microsoft.PowerShell_profile.ps1",
		)
	case "bash":
		file = "shells/shell.bashrc"
		dest = filepath.Join(home, ".bashrc")
	case "zsh":
		file = "shells/shell.sh"
		dest = filepath.Join(home, "zshrc")
	case "fish":
		file = "shells/shell.fish"
		dest = filepath.Join(
			home, ".config", "fish", "config.fish",
		)

	default:
		return fmt.Errorf("unsupported shell: %s", name)
	}
	return installScript(file, dest)
}

func installScript(file, dest string) error {
	data, err := globarvar.ShellScripts.ReadFile(file)
	if err != nil {
		return err
	}

	fmt.Println("Installing shell integration to:", dest)
	return appendFile(dest, string(data))
}

func appendFile(dest, data string) error {
	dir := filepath.Dir(dest)
	if err := os.MkdirAll(dir, 0355); err != nil {
		return err
	}
	if strings.Contains(data, globarvar.HasScript) {
		println("Script already installed")
		return nil
	}

	file, err := os.OpenFile(dest, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	strings.ReplaceAll(data, "_winTree_", exe)
	_, err = file.WriteString("\n" + data + "\n")
	return err
}

func validPath(path string) string {

	if path == "" {
		println("Missing path. Reverting to default path")
		dir, err := os.Getwd()
		if err != nil {

			println(err)
			println("Couldn't get current directory. Reverting to ' . '")
			path = "."
		} else {
			path = dir
		}
	}
	return path
}
