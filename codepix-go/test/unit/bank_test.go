package unit

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelmgr12/codepix-go/internal/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModelNewBank(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"

	bank, err := model.NewBank(code, name)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.NewString(), bank.ID)
	require.Equal(t, bank.Code, code)
	require.Equal(t, bank.Name, name)

	_, err = model.NewBank("", "")
	require.NotNil(t, err)
}
