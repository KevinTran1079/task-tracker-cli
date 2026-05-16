package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	tasksFile        = "tasks.json"
	statusTodo       = "todo"
	statusInProgress = "in-progress"
	statusDone       = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {

	if len(args) < 1 {
		return fmt.Errorf("please include a command")
	}

	tasks, err := LoadTasks()
	if err != nil {
		return fmt.Errorf("unable to load tasks: %w", err)
	}

	switch cmd := args[0]; cmd {
	case "add":
		if len(args) < 2 {
			return fmt.Errorf("description not included")
		}

		description := strings.Join(args[1:], " ")
		tasks = AddTask(description, tasks)

		if err := WriteTasks(tasks); err != nil {
			return fmt.Errorf("unable to write tasks: %w", err)
		}

	case "update":
		if len(args) < 3 {
			return fmt.Errorf("provide id and description to update task")
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
			return fmt.Errorf("please include ID")
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
			return fmt.Errorf("please include ID")
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
			return fmt.Errorf("please include ID")
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
		return fmt.Errorf("unknown command: %s", cmd)

	}

	return nil
}

func ListTasks(tasks []Task, status string) error {
	switch status {
	case "all":
		ListAllTasks(tasks)
	case statusTodo, statusInProgress, statusDone:
		ListTasksByStatus(tasks, status)
	default:
		return fmt.Errorf("unknown list status: %s", status)
	}

	return nil
}

func ListAllTasks(tasks []Task) {
	for _, task := range tasks {
		PrintTask(task)
	}
}

func ListTasksByStatus(tasks []Task, status string) {
	for _, task := range tasks {
		if task.Status == status {
			PrintTask(task)
		}
	}
}

func PrintTask(task Task) {
	fmt.Printf("id: %d\ndescription: %s\nstatus: %s\ncreatedat: %s\nupdatedat: %s\n\n",
		task.ID,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt)
}

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(tasksFile)
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

func AddTask(description string, tasks []Task) []Task {
	now := time.Now()
	lastID := 0
	if len(tasks) != 0 {
		lastID = tasks[len(tasks)-1].ID
	}

	task := Task{
		ID:          lastID + 1,
		Description: description,
		Status:      statusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return append(tasks, task)
}

func UpdateTask(id int, description string, tasks []Task) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func WriteTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(tasksFile, data, 0644)
}

func DeleteTask(id int, tasks []Task) ([]Task, error) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks = slices.Delete(tasks, i, i+1)
			return tasks, nil
		}
	}

	return tasks, fmt.Errorf("task with ID %d not found", id)
}

func UpdateStatus(id int, status string, tasks []Task) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func parseTaskID(value string) (int, error) {
	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid task ID %q: %w", value, err)
	}

	return id, nil
}
