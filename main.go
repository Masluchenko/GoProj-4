package main

import (
	"fmt"

	"example.com/m/v2/account"
)

func main() {
	fmt.Println("__МАНАГЕР_ПАРОЛЕЙ__")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
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

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func findAccount() {

}

func deleteAccount() {

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
