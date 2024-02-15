package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/samber/lo"
	"github.com/tcodes0/jaq/main/cmd"
)

const (
	escolaSP = "escolaSP"
)

func main() {
	command := flag.String("command", "", fmt.Sprintf("Command to run: %s", strings.Join(commands, ", ")))

	flag.Parse()

	if !lo.Contains(commands, *command) {
		log.Printf("error: invalid command %s", *command)
		flag.Usage()
		return
	}

	switch *command {
	case escolaSP:
		cmd.EscolaSP()
	}
}

// slice of supported commands
var commands = []string{escolaSP}
