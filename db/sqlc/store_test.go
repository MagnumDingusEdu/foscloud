package db

import (
	"context"
	"foscloud/utils"
	"github.com/lib/pq"
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

func TestStore_RegisterDuplicate(t *testing.T) {
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
	account2, err := store.RegisterAccount(context.Background(), args)
	require.Error(t, err)
	require.Equal(t, err.(*pq.Error).Code, pq.ErrorCode("23505")) // Duplicate object error

	require.Empty(t, account2)
}

func TestStore_CheckPassword(t *testing.T) {
	store := NewStore(testDb)
	account := createRandomizedAccount(t)
	password, err := store.CheckPassword(context.Background(), CheckPasswordTx{
		LoginID: account.Username,
	})
	require.NoError(t, err)
	require.NotEmpty(t, password)

	password, err = store.CheckPassword(context.Background(), CheckPasswordTx{})
	require.Error(t, err)
	require.Empty(t, password)
}
