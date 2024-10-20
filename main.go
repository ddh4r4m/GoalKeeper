package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var todos []Todo

func main() {
	printMenu()
	for {
		choice := getUserInput("Enter your choice: ")

		switch choice {
		case "1":
			addTodo()
		case "2":
			listTodos()
		case "3":
			markAsCompleted()
		case "4":
			deleteTodo()
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("\n--- Todo App Menu ---")
	fmt.Println("1. Add a new todo")
	fmt.Println("2. List all todos")
	fmt.Println("3. Mark a todo as completed")
	fmt.Println("4. Delete a todo")
	fmt.Println("5. Exit")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func addTodo() {
	task := getUserInput("Enter the todo task: ")
	todos = append(todos, Todo{task: task, isCompleted: false})
	fmt.Println("Todo added successfully")
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No Todos found!")
		return
	}

	for i, todo := range todos {
		status := " "
		if todo.isCompleted {
			status = "X"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, todo.task)
	}
}

func markAsCompleted() {
	listTodos()
	if len(todos) == 0 {
		fmt.Println("All Todos are completed!")
		return
	}

	choice := getUserInput("Enter the number of TODO to mark it as completed: ")
	index := getIndex(choice)
	if index != -1 {
		todos[index].isCompleted = true
		fmt.Println("TODO marked as completed ")
	}
}

func deleteTodo() {

}

func getIndex(choice string) int {
	index, err := strconv.Atoi(choice)
	print(index, " ", err)
	if err != nil || index < 1 || index > len(todos) {
		fmt.Println("Invalid choice please try agian")
		return -1
	}

	return index - 1
}
