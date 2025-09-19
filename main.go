package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccoundByUrl,
	"3": findAccoundByName,
	"4": deleteAccount,
} 

var menuVariants = []string{
	"1. Создать аккаунт", 
	"2. Найти аккаунт по URL", 
	"3. Найти аккаунт по логин", 
	"4. Удалить аккаунт", 
	"5. Выход", 
	"Выберите вариант",
}

func main() {
	err := godotenv.Load()

	if err != nil {
		output.PrintError("Неудалось найти env файл")
	}

	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())	

Menu:	
	for {
		variant := promptData(menuVariants...)

		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для удаления")
	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func findAccoundByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	}) 

	outputResult(&accounts)
}

func findAccoundByName(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	}) 

	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}

	for _, acc := range *accounts {
		acc.Output()
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	
	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}

	vault.AddAccount(*myAccount) 
}

func promptData(prompt ...string) string {
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
