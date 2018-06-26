package main

import (
	"flag"
	"log"
	"os"

	"github.com/Bo0km4n/bitxpp/printer"
	"github.com/Bo0km4n/bitxpp/repl"
)

var (
	decimalMode = "d"
	hexMode     = "h"
	bitMode     = "b"
	inputMode   = flag.String("m", "d", "input mode")
	input       = flag.String("i", "0", "input")
	consoleOpt  = flag.Bool("c", false, "console mode")
)

func main() {
	flag.Parse()

	switch *consoleOpt {
	case true:
		repl.Start(os.Stdin, os.Stdout)
	case false:
		printer := printer.NewPrinter(printer.ModeMap[*inputMode])
		if err := printer.OneShotPrint(*input); err != nil {
			log.Fatal("Input format is incollect")
		}
	}
}
