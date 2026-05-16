package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
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
	if err := run(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {

	if len(args) < 1 {
		return fmt.Errorf("Please include arguments")
	}

	tasks, err := LoadTasks()
	if err != nil {
		return fmt.Errorf("Unable to load tasks: %w", err)
	}

	switch cmd := args[0]; cmd {
	case "add":
		if len(args) < 2 {
			return fmt.Errorf("Description not included")
		}

		description := args[1]
		tasks = AddTasks(description, tasks)

		if err := WriteFile(tasks); err != nil {
			return fmt.Errorf("Unable to write to file: %w", err)
		}

	case "update":
		if len(args[1:]) < 2 {
			fmt.Println("Provide id and description to update task")
			os.Exit(1)
		}

		taskID, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Unable to convert task ID to integer")
			os.Exit(1)
		}
		description := args[2]
		tasks, err = UpdateTask(taskID, description, tasks)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := WriteFile(tasks); err != nil {
			fmt.Println("Unable to write updated tasks to tasks.json")
			os.Exit(1)
		}

	case "delete":
		if len(args) < 2 {
			fmt.Println("Please include ID")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Unable to convert string to int")
		}

		tasks, err = DeleteTask(id, tasks)
		if err != nil {
			os.Exit(1)
		}

		if err := WriteFile(tasks); err != nil {
			os.Exit(1)
		}
	case "mark-in-progress":
		fmt.Println(cmd)
	case "mark-done":
		fmt.Println(cmd)
	case "list":
		ListAllTasks(tasks)
	default:
		return fmt.Errorf("Unknown command")

	}

	return nil
}

func ListAllTasks(tasks []Task) {
		if len(tasks) > 0 {
		for _, task := range tasks {
			fmt.Printf("id: %d\ndescription: %s\nstatus: %s\ncreatedat: %s\nupdatedat: %s\n\n",
					task.ID,
					task.Description,
					task.Status,
					task.CreatedAt,
					task.UpdatedAt)
		}
	}
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
	lastID := 0
	if len(tasks) != 0 {
		lastID = tasks[len(tasks)-1].ID
	}

	task := Task{
		ID:          lastID + 1,
		Description: description,
		Status:      "in-progress",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return append(tasks, task)
}

func UpdateTask(id int, description string, tasks []Task) ([]Task, error) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return nil, fmt.Errorf("Task with ID %d not found", id)
}

func WriteFile(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile("tasks.json", data, 0644)
}

func DeleteTask(id int, tasks []Task) ([]Task, error) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks = slices.Delete(tasks, i, i+1)
			return tasks, nil
		}
	}

	return tasks, fmt.Errorf("Task with ID: %d not found", id)
}


