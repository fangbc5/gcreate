package flow

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func execCommand(projectName string, command string, args ...string) error {
	gopath := os.Getenv("GOPATH") // GOPATH环境变量指定的目录
	cmd := exec.Command(command, args...)
	cmd.Dir = gopath + "/src/" + projectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return err
	}
	log.Println(string(output))
	return nil
}
