package db

import (
	"context"
	"foscloud/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_RegisterAccount(t *testing.T) {
	store := NewStore(testDb)

	args := RegisterTxParams{
		Name:     utils.RandomName(),
		Username: utils.RandomUserName(),
		Email:    utils.RandomEmail(),
		Password: utils.RandomString(10),
	}

	account, err := store.RegisterAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Name, account.Name)
	require.Equal(t, args.Username, account.Username)
	require.Equal(t, args.Email, account.Email)
	result, err := utils.CheckPassword(account.Password, []byte(args.Password))
	require.NoError(t, err)
	require.True(t, result)
}
