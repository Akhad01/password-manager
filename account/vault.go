package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")

	if err != nil {
		return &Vault{
			Accounts: []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts: []Account{},
			UpdatedAt: time.Now(),
		}
	}

	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var results []Account

	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			results = append(results, acc)
		}
	}

	return results
}

func (vault *Vault) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false

	for _, acc := range vault.Accounts {
		if !strings.Contains(acc.Url, url) {
			accounts = append(accounts, acc)
			continue
		}
		isDeleted = true
	}

	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
} 