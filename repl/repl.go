package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Bo0km4n/bitxpp/printer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Reader) {
	scanner := bufio.NewScanner(in)
	printer := printer.Printer{
		Mode: printer.DECIMAL_MODE,
	}
	for {
		fmt.Printf("(%s) %s", printer.Mode, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		// TODO:
		// implement
		printer.Exec(line)
	}
}
