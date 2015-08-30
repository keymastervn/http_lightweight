package request

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func CheckDelayedJobs() string {
	cmd := exec.Command("ruby script/delayed_jobs status")
	cmd.Stdin = strings.NewReader("N/A")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}
