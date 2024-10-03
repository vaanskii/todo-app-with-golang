package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Description string
	Completed   bool
}

var todoList []Todo

func addTodo(description string, completed bool) {
	todo := Todo{Description: description, Completed: completed}
	todoList = append(todoList, todo)
	fmt.Println("\nTodo added successfully!")
}

func viewTodos() {
	if len(todoList) == 0 {
		fmt.Println("No todos available.")
		return
	}
	fmt.Println("\nTodo List:")
	for i, todo := range todoList {
		fmt.Printf("%d. Description: %s, Completed: %t\n", i+1, todo.Description, todo.Completed)
	}
}

func updateTodo(todo *Todo, newDescription string, completed bool) {
	todo.Description = newDescription
	todo.Completed = completed
	fmt.Printf("\nUpdated todo: %s (Completed: %t)\n", todo.Description, completed)
}

func deleteTodo(index int) {
	if index < 0 || index >= len(todoList) {
		fmt.Println("Invalid index. Cannot delete todo.")
		return
	} else {
		todoList = append(todoList[:index], todoList[index+1:]...)
		fmt.Println("Todo deleted successfully!")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		var choice int
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Todo")
		fmt.Println("2. View Todos")
		fmt.Println("3. Update Todo")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			var completed bool
			fmt.Print("Completed: (true/false): ")
			fmt.Scanln(&completed)

			addTodo(description, completed)

		case 2:
			viewTodos()

		case 3:
			if len(todoList) > 0 {
				viewTodos()
				var index int
				fmt.Print("\nEnter the number of the todo to update: ")
				fmt.Scanln(&index)

				// Clear the newline character left by fmt.Scanln
				fmt.Print("Enter new description: ")
				newDescription, _ := reader.ReadString('\n')
				newDescription = strings.TrimSpace(newDescription)

				var completed bool
				fmt.Print("Is it completed (true/false)? ")
				fmt.Scanln(&completed)

				if index > 0 && index <= len(todoList) {
					updateTodo(&todoList[index-1], newDescription, completed)
				} else {
					fmt.Println("Invalid todo number.")
				}
			} else {
				fmt.Println("\nNo todos available.")
			}
		case 4:
			if len(todoList) > 0 {
				viewTodos()
				var index int
				fmt.Print("\nEnter the number of the todo to delete: ")
				fmt.Scanln(&index)
				index--

				if index >= 0 && index < len(todoList) {
					deleteTodo(index)
				} else {
					fmt.Println("Invalid todo number.")
				}
			} else {
				fmt.Println("\nNo todos available.")
			}
		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
