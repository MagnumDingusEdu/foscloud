package db

import (
	"context"
	"database/sql"
	"foscloud/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomizedAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Name:     utils.RandomName(),
		Username: utils.RandomUserName(),
		Email:    utils.RandomEmail(),
		Password: utils.RandomPassword(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.Password, account.Password)
	require.NotZero(t, account.LastLogin)
	require.NotZero(t, account.ID)
	return account
}

func TestQueries_CreateAccount(t *testing.T) {
	createRandomizedAccount(t)
}

func TestQueries_GetAccount(t *testing.T) {
	account1 := createRandomizedAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.Password, account2.Password)
	require.Equal(t, account1.Email, account2.Email)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestQueries_UpdateAccount(t *testing.T) {
	account1 := createRandomizedAccount(t)

	args := UpdateAccountParams{
		ID:       account1.ID,
		Name:     utils.RandomName(),
		Username: utils.RandomUserName(),
		Email:    utils.RandomEmail(),
		Password: utils.RandomPassword(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, args.Name, account2.Name)
	require.Equal(t, args.Username, account2.Username)
	require.Equal(t, args.Email, account2.Email)
	require.Equal(t, args.Password, account2.Password)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestQueries_DeleteAccount(t *testing.T) {
	account1 := createRandomizedAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestQueries_UpdateLastLogin(t *testing.T) {
	account1 := createRandomizedAccount(t)

	args := UpdateLastLoginParams{
		ID:        account1.ID,
		LastLogin: time.Now(),
	}

	account2, err := testQueries.UpdateLastLogin(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Password, account2.Password)

	require.WithinDuration(t, args.LastLogin, account2.LastLogin, time.Second)

}

func TestQueries_GetAccountForUpdate(t *testing.T) {
	account1 := createRandomizedAccount(t)

	account2, err := testQueries.GetAccountForUpdate(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.Password, account2.Password)
	require.Equal(t, account1.Email, account2.Email)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestQueries_ListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomizedAccount(t)
	}
	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
