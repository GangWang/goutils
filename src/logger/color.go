package logger

import (
	"fmt"
	"runtime"
	"strings"
)

type ColorCode string

const (
	red    ColorCode = "1;31"
	green            = "1;32"
	yellow           = "1;33"
	blue             = "1;34"
	pink             = "1;35"
	cyan             = "1;36"
	white            = "1;37"
)

func colorize(row string, color ColorCode) string {
	switch runtime.GOOS {
	case "linux":
		return fmt.Sprintf("%s%s%s", "\033["+color+"m", row, "\033[0m")
	}
	return row
}

func Red(row string) string {
	return colorize(row, red)
}

func Green(row string) string {
	return colorize(row, green)
}

func Yellow(row string) string {
	return colorize(row, yellow)
}

func Blue(row string) string {
	return colorize(row, blue)
}

func Pink(row string) string {
	return colorize(row, pink)
}

func Cyan(row string) string {
	return colorize(row, cyan)
}

func White(row string) string {
	return colorize(row, white)
}

func Indent(v interface{}) string {
	return strings.Replace(fmt.Sprintf("%v", v), "\n", "\n    ", -1)
}
