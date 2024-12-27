package main

import (
	"fmt"

	"example.com/m/v2/account"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("__МАНАГЕР_ПАРОЛЕЙ__")
	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var varik int
	fmt.Println(`Выберите действие:
	1. Создать аккаунт
	2. Найти аккаунт
	3. Удалить аккаунт
	4. Выход`)
	fmt.Scanln(&varik)
	return varik
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите url для поиска")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}

}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите url для удаления")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
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
