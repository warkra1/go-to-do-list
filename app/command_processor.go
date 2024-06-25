package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"to-do-list/app/model"
	"to-do-list/app/service"
)

func readString(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

func readNumber(message string) int {
	for {
		numberStr := readString(message)
		number, err := strconv.Atoi(numberStr)
		if err == nil {
			return number
		}
		fmt.Println("Invalid number!")
	}
}

type Command int64

const (
	Login      Command = iota
	REGISTER   Command = iota
	GetList    Command = iota
	CreateItem Command = iota
	UpdateItem Command = iota
	DeleteItem Command = iota
	Exit       Command = iota
	Help       Command = iota
)

func (c Command) String() string {
	switch c {
	case Login:
		return "login"
	case REGISTER:
		return "register"
	case GetList:
		return "get_list"
	case CreateItem:
		return "create_item"
	case UpdateItem:
		return "update_item"
	case DeleteItem:
		return "delete_item"
	case Exit:
		return "exit"
	case Help:
		return "help"
	}

	return "undefined"
}

func printCommands() {
	fmt.Println("List of commands: ")
	fmt.Printf("\t%s - %s\n", Login.String(), "authorize to application")
	fmt.Printf("\t%s - %s\n", REGISTER.String(), "create new user")
	fmt.Printf("\t%s - %s\n", GetList.String(), "get current user's todo list")
	fmt.Printf("\t%s - %s\n", CreateItem.String(), "create a new item in todo list")
	fmt.Printf("\t%s - %s\n", UpdateItem.String(), "updates a specific item in list")
	fmt.Printf("\t%s - %s\n", DeleteItem.String(), "deletes a specific item in list")
	fmt.Printf("\t%s - %s\n", Help.String(), "show commands")
	fmt.Printf("\t%s - %s\n", Exit.String(), "exit from application")
}

type CommandProcessor struct {
	authService     *service.AuthService
	toDoListService *service.ToDoListService
}

func NewCommandProcessor(authService *service.AuthService, toDoListService *service.ToDoListService) *CommandProcessor {
	return &CommandProcessor{authService: authService, toDoListService: toDoListService}
}

func (cp *CommandProcessor) Run() {
	fmt.Println("Hello, this is a to do list CLI App")

	for {
		command := readString("Please, Enter a command (or type \"help\" tp see all commands): ")
		switch command {
		case Login.String():
			cp.login()
			break
		case REGISTER.String():
			cp.register()
			break
		case GetList.String():
			cp.getList()
			break
		case CreateItem.String():
			cp.createItem()
			break
		case UpdateItem.String():
			cp.updateItem()
			break
		case DeleteItem.String():
			cp.deleteItem()
			break
		case Help.String():
			printCommands()
			break
		case Exit.String():
			exit()
		default:
			fmt.Println("Invalid command!")
		}
	}
}

func (cp *CommandProcessor) checkAuth() bool {
	if cp.authService.GetCurrentUser().Login != "" {
		return true
	}
	fmt.Println("You need to login first!")
	return false
}

func (cp *CommandProcessor) login() {
	login := readString("Enter login: ")
	password := readString("Enter password: ")

	err := cp.authService.Login(login, password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Login success!")
}

func (cp *CommandProcessor) register() {
	login := readString("Enter login: ")
	password := readString("Enter password: ")

	err := cp.authService.Register(login, password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully registered!")
}

func (cp *CommandProcessor) getList() {
	if !cp.checkAuth() {
		return
	}

	list := cp.toDoListService.GetList(cp.authService.GetCurrentUser())
	fmt.Println("To Do List: ")
	if len(list.Items) == 0 {
		fmt.Println("\tYour list is empty")
	} else {
		for i, item := range list.Items {
			fmt.Printf("\t[%d] %s\n", i+1, item.Title)
		}
	}
}

func (cp *CommandProcessor) createItem() {
	if !cp.checkAuth() {
		return
	}

	title := readString("Enter title: ")
	number := readNumber("Enter number: ")

	cp.toDoListService.CreateItem(cp.authService.GetCurrentUser(), model.ToDoItem{Title: title}, number-1)
	fmt.Println("Item successfully created!")
}

func (cp *CommandProcessor) updateItem() {
	if !cp.checkAuth() {
		return
	}

	number := readNumber("Enter number of item: ")
	title := readString("Enter title: ")

	cp.toDoListService.UpdateItem(cp.authService.GetCurrentUser(), model.ToDoItem{Title: title}, number-1)
	fmt.Println("Item successfully updated!")
}

func (cp *CommandProcessor) deleteItem() {
	if !cp.checkAuth() {
		return
	}

	number := readNumber("Enter number of item: ")
	cp.toDoListService.DeleteItem(cp.authService.GetCurrentUser(), number-1)
	fmt.Println("Item successfully deleted!")
}

func exit() {
	fmt.Println("Bye!")
	os.Exit(0)
}
