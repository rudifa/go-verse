package future

import (
	"os/exec"
)

func OpenOpenAIPage() error {
	cmd := exec.Command("open", "https://openai.com")
	return cmd.Run()
}

func OpenUzufly() error {
	cmd := exec.Command("open", "https://uzufly.com")
	return cmd.Run()
}
