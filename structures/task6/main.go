package main

import (
	"fmt"
)

func main() {
	tm := TaskManager{}
	tm.AddTask("steps")
	tm.AddTask("practice go")
	err := tm.CompleteTask(2)
	if err == nil {
		fmt.Println("completed task 2")
	}
	listTasks := tm.ListTasks()
	for _, task := range listTasks {
		if task.Done {
			fmt.Printf("%d: %s [done]\n", task.ID, task.Title)
		} else {
			fmt.Printf("%d: %s [not done]\n", task.ID, task.Title)
		}
	}
	rem := tm.RemoveTask(1) //КОД ДНЯ 1154
	if rem != nil {
		fmt.Printf("dont remove task")
	} else {
		fmt.Printf("removed task\n")
	}
	listTasks = tm.ListTasks()
	for _, task := range listTasks {
		fmt.Printf("%d: %s\n", task.ID, task.Title)
	}

}

type Task struct {
	ID    int
	Title string
	Done  bool
}
type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) AddTask(title string) {
	id := len(tm.tasks) + 1
	tm.tasks = append(tm.tasks, Task{ID: id, Title: title, Done: false})
}
func (tm *TaskManager) CompleteTask(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			if tm.tasks[i].Done {
				return fmt.Errorf("task %d is already done", id)
			}
			tm.tasks[i].Done = true
			return nil

		}
	}
	return fmt.Errorf("task %d not found", id)
}
func (tm TaskManager) ListTasks() []Task {
	return tm.tasks
}

func (tm *TaskManager) RemoveTask(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}
