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
	UpdatedAt		time.Time
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("No command provided")
		os.Exit(1)
	}


	switch cmd := args[0]; cmd {
	case "add":
		fmt.Println(cmd)
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