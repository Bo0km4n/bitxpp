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
	inputMode   = flag.String("m", "d", "input mode \n'd: deciaml' \n'h: hex' \n'b: bit'")
	input       = flag.String("i", "0", "input text")
	consoleOpt  = flag.Bool("c", false, "console mode \ncommands: \n\t'/cd => change decimal mode' \n\t'/ch => change hex mode' \n\t'/cb => change bit mode'")
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
