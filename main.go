package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

var (
	decimalMode = "d"
	hexMode     = "h"
	bitMode     = "b"
	inputMode   = flag.String("m", "d", "input mode")
	input       = flag.String("i", "0", "input")
	consoleOpt  = flag.Bool("c", false, "console mode")

	currentMode string
)

func init() {
	currentMode = decimalMode
}

func main() {
	flag.Parse()

	switch *consoleOpt {
	case true:
		// TODO: implement console mode
		fmt.Println("sorry, console mode has not implemented yet.")
	case false:
		oneShotPrint(*inputMode, *input)
	}
}

func oneShotPrint(mode, input string) error {
	var v uint64
	switch mode {
	case decimalMode:
		si, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return err
		}
		v = (uint64)(si)
	case hexMode:
		ui, err := strconv.ParseUint(input, 16, 64)
		if err != nil {
			return err
		}
		v = ui
	case bitMode:
		ui, err := strconv.ParseUint(input, 2, 64)
		if err != nil {
			return err
		}
		v = ui
	}
	colorPrint(v)
	return nil
}

func colorPrint(in uint64) error {
	formatD := "(decimal): " + "%" + fmt.Sprintf("%d", getMaxStringLen(in))
	formatX := "(hex)    : " + "%" + fmt.Sprintf("%d", getMaxStringLen(in))
	formatB := "(bit)    : " + "%" + fmt.Sprintf("%d", getMaxStringLen(in))
	color.Cyan(formatD+"d", in)
	color.Red(formatX+"x", in)
	color.Yellow(formatB+"b", in)
	return nil
}

func getMaxStringLen(in uint64) int {
	buf := make([]int, 0)
	max := 0
	buf = append(buf, len(fmt.Sprintf("%d", in)))
	buf = append(buf, len(fmt.Sprintf("%x", in)))
	buf = append(buf, len(fmt.Sprintf("%b", in)))

	for i := range buf {
		if max < buf[i] {
			max = buf[i]
		}
	}
	return max
}
