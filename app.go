package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run(args []string) error {
	if len(args) < 1 {
		return usageErr("please include a command")
	}

	if isHelpCommand(args[0]) {
		PrintUsage(os.Stdout)
		return nil
	}

	tasks, err := LoadTasks()
	if err != nil {
		return fmt.Errorf("unable to load tasks: %w", err)
	}

	switch cmd := args[0]; cmd {
	case "add":
		if len(args) < 2 {
			return usageErr("description not included")
		}

		description := strings.Join(args[1:], " ")
		tasks = AddTask(description, tasks)

		if err := WriteTasks(tasks); err != nil {
			return fmt.Errorf("unable to write tasks: %w", err)
		}

	case "update":
		if len(args) < 3 {
			return usageErr("provide id and description to update task")
		}

		taskID, err := parseTaskID(args[1])
		if err != nil {
			return err
		}
		description := strings.Join(args[2:], " ")
		if err := UpdateTask(taskID, description, tasks); err != nil {
			return err
		}
		if err := WriteTasks(tasks); err != nil {
			return fmt.Errorf("unable to write updated tasks: %w", err)
		}

	case "delete":
		if len(args) < 2 {
			return usageErr("please include ID")
		}

		id, err := parseTaskID(args[1])
		if err != nil {
			return err
		}

		tasks, err = DeleteTask(id, tasks)
		if err != nil {
			return err
		}

		if err := WriteTasks(tasks); err != nil {
			return fmt.Errorf("unable to write tasks: %w", err)
		}

	case "mark-in-progress":
		if len(args) < 2 {
			return usageErr("please include ID")
		}
		id, err := parseTaskID(args[1])
		if err != nil {
			return err
		}
		if err := UpdateStatus(id, statusInProgress, tasks); err != nil {
			return err
		}
		if err := WriteTasks(tasks); err != nil {
			return fmt.Errorf("unable to write tasks: %w", err)
		}

	case "mark-done":
		if len(args) < 2 {
			return usageErr("please include ID")
		}
		id, err := parseTaskID(args[1])
		if err != nil {
			return err
		}
		if err := UpdateStatus(id, statusDone, tasks); err != nil {
			return err
		}

		if err := WriteTasks(tasks); err != nil {
			return fmt.Errorf("unable to write tasks: %w", err)
		}

	case "list":
		status := "all"
		if len(args) > 1 {
			status = args[1]
		}

		if err := ListTasks(tasks, status); err != nil {
			return err
		}

	default:
		return usageErr("unknown command: %s", cmd)
	}

	return nil
}

func parseTaskID(value string) (int, error) {
	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, usageErr("invalid task ID %q: %w", value, err)
	}

	return id, nil
}
