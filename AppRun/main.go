package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	self, err := os.Executable()
	if err != nil { panic(err) }
	base := filepath.Dir(self)

	cmd := exec.Command(base + "/bin/redis-server")
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cmd.Wait()

	time.Sleep(time.Second)
	cmd2 := exec.Command(base + "/bin/redis-cli", "set", "hello", "goodbye")
	cmd2.Stdout = os.Stdout
	if err := cmd2.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
