package main

import (
	"fmt"

	"example.com/m/v2/account"
	"example.com/m/v2/files"
)

func main() {
	files.ReadFile()
	files.WriteFile("ХЭЛОУ ФАЙЛ СОЗДАН", "file.txt")
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN")
		return
	}
	myAccount.OutputPassword()
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
