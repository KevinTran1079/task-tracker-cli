package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Unable to load task")
		os.Exit(1)
	}

	switch cmd := args[0]; cmd {
	case "add":
		if len(args) < 2 {
			fmt.Println("Missing description of the task")
			os.Exit(1)
		}

		description := args[1]
		tasks = AddTasks(description, tasks)

		if err := WriteFile(tasks); err != nil {
			fmt.Println("unable to write file:", err)
			os.Exit(1)
		}
	case "update":
		fmt.Println(cmd)
	case "delete":
		fmt.Println(cmd)
	case "mark-in-progress":
		fmt.Println(cmd)
	case "mark-done":
		fmt.Println(cmd)
	case "list":
		fmt.Println(cmd)
	default:
		fmt.Println("Unknown command:", cmd)
		os.Exit(1)

	}

	os.Exit(0)
}

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func AddTasks(description string, tasks []Task) []Task {
	now := time.Now()
	lastID := tasks[len(tasks)-1].ID

	task := Task{
		ID:          lastID + 1,
		Description: description,
		Status:      "in-progress",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return append(tasks, task)
}

func WriteFile(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile("tasks.json", data, 0644)
}
