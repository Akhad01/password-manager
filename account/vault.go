package account

import (
	"demo/password/encrypter"
	"demo/password/output"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db Db
	enc encrypter.Encrypter
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
			enc: enc,
		}
	}

	data := enc.Decrypt(file)

	var vault Vault
	err = json.Unmarshal(data, &vault)
	color.Cyan("Найдено %d аккаунтов", len(vault.Accounts))
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.vault")
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
			enc: enc,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db: db,
		enc: enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccounts(url string, checker func(Account, string) bool) []Account {
	var results []Account

	for _, acc := range vault.Accounts {
		if checker(acc, url) {
			results = append(results, acc)
		}
	}

	return results
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
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

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	encData := vault.enc.Encrypt(data)

	if err != nil {
		output.PrintError("Не удалось преобразовать")
	}

	vault.db.Write(encData)
} 