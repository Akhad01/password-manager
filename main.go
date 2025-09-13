package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJsonDb("data.json"))	

Menu:	
	for {
		variant := promptData([]string{
		"1. Создать аккаунт", 
		"2. Найти аккаунт", 
		"3. Удалить аккаунт", 
		"4. Выход", 
		"Выберите вариант",
	})

		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccound(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для удаления"})
	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func findAccound(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	accounts := vault.FindAccountsByUrl(url) 

	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}

	for _, acc := range accounts {
		acc.Output()
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})
	myAccount, err := account.NewAccount(login, password, url)
	
	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}

	vault.AddAccount(*myAccount) 
}

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt) - 1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	} 
	
	var res string
	fmt.Scanln(&res)
	return res
}
