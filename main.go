package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"main/app"

	tea "github.com/charmbracelet/bubbletea"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	clearScreen()
	p := tea.NewProgram(app.NewMenuModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
