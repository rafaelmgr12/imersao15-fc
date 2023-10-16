package unit

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelmgr12/codepix-go/internal/domain/model"

	"github.com/stretchr/testify/require"
)

func TestModelNewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)
	require.Nil(t, err)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, err := model.NewAccount(bank, accountNumber, ownerName)
	require.Nil(t, err)

	kind := "email"
	key := "j@j.com"
	pixKey, err := model.NewPixKey(kind, account, key)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.NewString(), pixKey.ID)
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, err = model.NewPixKey(kind, account, key)
	require.Nil(t, err)

	_, err = model.NewPixKey("nome", account, key)
	require.NotNil(t, err)
}
