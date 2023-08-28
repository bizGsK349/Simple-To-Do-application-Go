package main

import (
	"fmt"
)

var tasks []Task

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

func addTask(title, description string) {
	task := Task{
		ID:          len(tasks) + 1,
		Title:       title,
		Description: description,
		Done:        false,
	}
	tasks = append(tasks, task)
}

func listTasks() {
	fmt.Println("Tasks:")
	for _, task := range tasks {
		status := "Not Done"
		if task.Done {
			status = "Done"
		}
		fmt.Println("ID: %d | Title: %s | Description: %s | Status: %s\n", task.ID, task.Title, task.Description, status)
	}
}

func markTaskDone(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			break
		}
	}
}

func mainMenu() {
	fmt.Println("To-Do List CLI App")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Mark Tasks as Done")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Please enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var title, description string
		fmt.Print("Enter task title: ")
		fmt.Scan(&title)
		fmt.Print("Enter the task description")
		fmt.Print(&description)
		addTask(title, description)
	case 2:
		listTasks()
	case 3:
		fmt.Println("Please enter the task number to mark it as done")
		var id int
		fmt.Scan(&id)
		markTaskDone(id)
	case 4:
		fmt.Println("Bye")
		return
	default:
		fmt.Println("invalid choice")
	}
	mainMenu()
}

func main() {
	mainMenu()
}
