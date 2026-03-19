package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todo-go-cli/todo"
)

func main() {
	list := todo.NewTodoList()
	fmt.Println(`
Начало работы

	================================================
	add <задача> – добавить задачу в список
	list – показать все задачи в списке
	done <номер задачи> – отметить задачу выполненой
	delete <номер задачи> – удалить задачу
	================================================
	exit – выйти
	`)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]
		switch cmd {
		case "add":
			list.Add(parts[1], false)
		case "done":
			id, _ := strconv.Atoi(parts[1])
			list.Complete(id)
		case "list":
			list.Print()
		case "exit":
			return
		case "delete":
			id, _ := strconv.Atoi(parts[1])
			list.DeleteTask(id)
		}
	}
}
