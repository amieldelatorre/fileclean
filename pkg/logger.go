package pkg

import (
	"fmt"
	"os"
	"strings"
)

const (
	reset  = "\x1b[0m"
	red    = "\x1b[31m"
	green  = "\x1b[32m"
	yellow = "\x1b[33m"
	blue   = "\x1b[34m"
)

func wrapArgs(colour string, args ...string) []any {
	wrappedArgs := []any{}
	for _, arg := range args {
		wrappedArgs = append(wrappedArgs, colour+arg+reset)
	}

	return wrappedArgs
}

func addNewline(format string) string {
	if strings.HasSuffix(format, "\n") {
		return format
	}
	return format + "\n"
}

func LogInformation(format string, args ...string) {
	wrappedArgs := wrapArgs(blue, args...)
	formatWithNewline := addNewline(format)
	fmt.Printf(formatWithNewline, wrappedArgs...)
}

func LogWarn(format string, args ...string) {
	wrappedArgs := wrapArgs(yellow, args...)
	formatWithNewline := addNewline(format)
	fmt.Printf(formatWithNewline, wrappedArgs...)
}

func LogError(format string, args ...string) {
	wrappedArgs := wrapArgs(red, args...)
	formatWithNewline := addNewline(format)
	fmt.Printf(formatWithNewline, wrappedArgs...)
}

func LogErrorExit(format string, args ...string) {
	LogError(format, args...)
	os.Exit(1)
}
