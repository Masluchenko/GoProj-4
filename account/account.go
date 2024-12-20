package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWYZ123456789*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputPassword() {
	color.Magenta(acc.Password, "Raskraska ebychaya")
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func (acc *Account) generateUrl(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Url = "https://" + string(res) + ".com"

}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  password,
		Login:     login,
		Url:       urlString,
	}

	if urlString == "" {
		newAcc.generateUrl(8)
		fmt.Println(newAcc.Url)
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	// field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	// fmt.Println(string(field.Tag))
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
