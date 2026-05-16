package main

import (
	"errors"
	"fmt"
	"io"
)

const usageText = `Usage:
  task-tracker-cli add <description>
  task-tracker-cli update <id> <description>
  task-tracker-cli delete <id>
  task-tracker-cli mark-in-progress <id>
  task-tracker-cli mark-done <id>
  task-tracker-cli list [all|todo|in-progress|done]
  task-tracker-cli help`

type usageError struct {
	err error
}

func (e *usageError) Error() string {
	return e.err.Error()
}

func (e *usageError) Unwrap() error {
	return e.err
}

func usageErr(format string, args ...any) error {
	return &usageError{err: fmt.Errorf(format, args...)}
}

func isUsageError(err error) bool {
	var usageErr *usageError
	return errors.As(err, &usageErr)
}

func isHelpCommand(cmd string) bool {
	return cmd == "help" || cmd == "-h" || cmd == "--help"
}

func PrintUsage(w io.Writer) {
	fmt.Fprintln(w, usageText)
}
