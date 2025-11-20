package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

	prompt := "Give me only the bash commands to restructure the following tree. Start with '#!/bin/bash' and output ONLY code, no comments, markdown, or explanations. Absolutely no text outside the bash script:\n\n" + string(out) + "\n\n" + response
	commands := llm.LlmQuery(prompt)
	if strings.Contains(commands, "rm ") || strings.Contains(commands, ":(){") {
		fmt.Println(" Dangerous commands detected in output. Not saving.")
		return
	}
	start := strings.Index(commands, "#!/bin/bash")
	if start == -1 {
		fmt.Println("No bash script found in output")
		return
	}

	end := len(commands)
	if idx := strings.LastIndex(commands, "```"); idx != -1 && idx > start {
		end = idx
	}

	script := commands[start:end]

	if err := os.WriteFile("commands.sh", []byte(script), 0755); err != nil {
		fmt.Println("Error writing script:", err)
		return
	}
	if err := os.Chmod("commands.sh", 0755); err != nil {
		fmt.Println("Error chmod script:", err)
		return
	}
	os.WriteFile("commands.sh", []byte(script), 0755)
	os.Chmod("commands.sh", 0755)
	fmt.Print("Run the generated script? (y/N): ")
	var resp string
	fmt.Scanln(&resp)
	if strings.ToLower(resp) == "y" {
		cmd := exec.Command("bash", "commands.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
