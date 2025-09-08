package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	
Menu:	
	for {
		variant := getMenu()

		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccound()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scanln(&variant)

	return variant
}

func deleteAccount() {}

func findAccound() {
}

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	
	if err != nil {
		fmt.Println("Неверный формат URL или Логин") 
		return
	}

	file, err := myAccount.ToBytes()

	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return 
	}

	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}


