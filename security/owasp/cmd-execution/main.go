package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	in := flag.String("cmd", "", "echo Hello World")
	flag.Parse()

	// 1. Receive Command

	inSplitter := strings.Split(*in, " ")

	cmd := exec.Command(inSplitter[0], inSplitter[1:]...)
	cmd.Stdout = os.Stdout

	// 2. Execute Command
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
