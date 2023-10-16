package unit

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelmgr12/codepix-go/internal/domain/model"
	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)
	require.Nil(t, err)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.NewString(), account.ID)
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.BankID, bank.ID)
	require.Equal(t, account.OwnerName, ownerName)

	_, err = model.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}
