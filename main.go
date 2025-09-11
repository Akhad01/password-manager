package main

import (
	"demo/password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault()	
Menu:	
	for {
		variant := getMenu()

		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccound(vault)
		case 3:
			deleteAccount(vault)
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

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите URL для удаления")
	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
}

func findAccound(vault *account.Vault) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrl(url) 

	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}

	for _, acc := range accounts {
		acc.Output()
	}
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	
	if err != nil {
		fmt.Println("Неверный формат URL или Логин") 
		return
	}

	vault.AddAccount(*myAccount) 
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}


