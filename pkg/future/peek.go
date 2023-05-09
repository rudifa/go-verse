package future

import (
	"os/exec"
)

func OpenOpenAIPage() error {
    cmd := exec.Command("open", "https://openai.com")
    return cmd.Run()
}
