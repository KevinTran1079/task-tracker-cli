package main

import (
	"fmt"
	"slices"
	"time"
)

const (
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
