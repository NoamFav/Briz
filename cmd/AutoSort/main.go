package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/NoamFav/AutoSort/internal/llm"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: autosort <path>")
		return
	}

	path, err := filepath.Abs(args[0])
	if err != nil {
		fmt.Println("Path incorrect or does not exist")
		return
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Path does not exist")
		return
	}
	out, err := exec.Command("/opt/homebrew/bin/eza", "--tree", path).Output()
	if err != nil {
		fmt.Println("Error running ls:", err)
		return
	}

	response := llm.LlmQuery("Reorganize...: " + string(out))
	os.WriteFile("autosort_suggestion.md", []byte(response), 0644)
}
