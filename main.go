package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"main/app"
	"main/config"

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

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Using default theme (light):", err)
		cfg = &config.Config{Theme: "light"}
	}

	if cfg.Theme != "dark" && cfg.Theme != "light" {
		log.Fatal("Chosen theme not known. Please use 'light' or 'dark'")
	}

	os.Setenv("theme", cfg.Theme)

	p := tea.NewProgram(app.NewMenuModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
