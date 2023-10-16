package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `json:"bank_id" valid:"uuid"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (account *Account) isValid() error {
	if account.Bank != nil {
		account.BankID = account.Bank.ID
	}
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {

	account := Account{
		OwnerName: ownerName,
		Bank:      bank,
		Number:    number,
	}
	account.ID = uuid.NewString()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
