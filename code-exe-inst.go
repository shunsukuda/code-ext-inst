package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	fn := os.Args[1]
	b, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(b, []byte{'\n'})
	for i := range lines {
		if len(lines[i]) > 0 {
			result, err := exec.Command("code.cmd", "--install-extension", string(lines[i])).Output()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(result))
			}
		}
	}
}
