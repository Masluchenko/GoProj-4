package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWYZ123456789*!")

type Account struct {
	login    string `demo:"login" xml:"test"`
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *Account) OutputPassword() {
	color.Red(acc.login)
}

func (acco AccountWithTimeStamp) OutputPassword() {
	fmt.Println(acco.login, acco.password, acco.url, acco.createdAt, acco.updatedAt)
	color.Magenta(acco.password, "Raskraska ebychaya")
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		acc: Account{
			password: password,
			login:    login,
			url:      urlString,
		},
	}
	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	fmt.Println(string(field.Tag))
	if password == "" {
		newAcc.acc.generatePassword(12)
	}
	return newAcc, nil
}
