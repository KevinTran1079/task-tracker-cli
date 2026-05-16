package main

import "fmt"

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
