package main

import (
	"fmt"
	"strings"

	"example.com/m/v2/account"
	"example.com/m/v2/files"
	"example.com/m/v2/output"

	"github.com/fatih/color"
)

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите действие",
}

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountUrl,
	"3": findAccountLogin,
	"4": deleteAccount,
}

func main() {
	fmt.Println("__МАНАГЕР_ПАРОЛЕЙ__")
	vault := account.NewVault(files.NewJsonDb("data.json"))
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

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintErorr("Неверный формат URL или LOGIN")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccountUrl(vault *account.VaultWithDb) {
	url := promptData("Введите url для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func findAccountLogin(vault *account.VaultWithDb) {
	login := promptData("Введите login для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url для удаления")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintErorr("Не найдено")
	}
}

func promptData(prompt ...string) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

// func generatePassword(n int) string {
// 	res := make([]rune, n)
// 	for i := range res {
// 		res[i] = letterRunes[rand.IntN(len(letterRunes))]
// 	}
// 	return string(res)
// }

// str := []rune("Привет!)")
// for _, ch := range str {
// 	fmt.Println(ch, string(ch))
// }

// a := [4]int{1, 2, 3, 4}
// 	reverse(&a)
// 	fmt.Print(a)
// func reverse(arr *[4]int) {
// 	for index, value := range *arr {
// 		(*arr)[len(arr)-1-index] = value
// 	}
// }
