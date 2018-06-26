package printer

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

const (
	HEX_MODE     = "hex"
	BIT_MODE     = "bit"
	DECIMAL_MODE = "decimal"

	SwitchDecimal = `/cd`
	SwitchBit     = "/cb"
	SwitchHex     = "/cx"
)

// Printer //
type Printer struct {
	Mode string
}

var ModeMap = map[string]string{
	"d": DECIMAL_MODE,
	"b": BIT_MODE,
	"x": HEX_MODE,
}

func NewPrinter(mode string) *Printer {
	return &Printer{
		Mode: mode,
	}
}

// OneShotPrint print some format from input string
func (p *Printer) OneShotPrint(input string) error {
	var v uint64
	switch p.Mode {
	case DECIMAL_MODE:
		si, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return err
		}
		v = (uint64)(si)
	case HEX_MODE:
		ui, err := strconv.ParseUint(input, 16, 64)
		if err != nil {
			return err
		}
		v = ui
	case BIT_MODE:
		ui, err := strconv.ParseUint(input, 2, 64)
		if err != nil {
			return err
		}
		v = ui
	}
	colorPrint(v)
	return nil
}

func (p *Printer) Exec(input string) {
	switch input {
	case SwitchDecimal:
		p.Mode = DECIMAL_MODE
	case SwitchHex:
		p.Mode = HEX_MODE
	case SwitchBit:
		p.Mode = BIT_MODE
	default:
		p.OneShotPrint(input)
	}
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
