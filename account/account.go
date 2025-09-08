package account

import (
	"encoding/json"
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()")

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login)
	color.Blue(acc.Password)
	color.Yellow(acc.Url)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (acc *Account) generatePassword(n int) {
	passwordRune := make([]rune, n)
	for i := range passwordRune {
		passwordRune[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(passwordRune)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:    login,
		Password: password,
		Url:      urlString,
	}

	if password == "" {
		newAcc.generatePassword(32)
	}

	return newAcc, nil
}
