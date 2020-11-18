package db

import (
	"context"
	"database/sql"
	"foscloud/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomizedAuthToken(t *testing.T) Authtoken {
	account := createRandomizedAccount(t)
	args := CreateAuthTokenParams{
		Token:   utils.GenerateAuthToken(),
		Account: account.ID,
	}
	authtoken, err := testQueries.CreateAuthToken(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, authtoken)
	require.NotEmpty(t, authtoken.ID)

	require.Equal(t, args.Token, authtoken.Token)
	require.Equal(t, account.ID, authtoken.Account)
	require.NotEmpty(t, authtoken.CreatedAt)
	require.NotEmpty(t, authtoken.LastUsed)

	return authtoken
}

func TestQueries_CreateAuthToken(t *testing.T) {
	createRandomizedAuthToken(t)

}

func TestQueries_GetAuthTokenByID(t *testing.T) {
	authtoken1 := createRandomizedAuthToken(t)
	authtoken2, err := testQueries.GetAuthTokenByID(context.Background(), authtoken1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, authtoken2)

	require.Equal(t, authtoken1.ID, authtoken2.ID)
	require.Equal(t, authtoken1.Account, authtoken2.Account)
	require.Equal(t, authtoken1.Token, authtoken2.Token)

	require.WithinDuration(t, authtoken1.CreatedAt, authtoken2.CreatedAt, time.Second)
	require.WithinDuration(t, authtoken1.LastUsed, authtoken2.LastUsed, time.Second)
}

func TestQueries_GetAuthTokenByAccount(t *testing.T) {
	authtoken1 := createRandomizedAuthToken(t)
	authtoken2, err := testQueries.GetAuthTokenByAccount(context.Background(), authtoken1.Account)
	require.NoError(t, err)
	require.NotEmpty(t, authtoken2)

	require.Equal(t, authtoken1.ID, authtoken2.ID)
	require.Equal(t, authtoken1.Account, authtoken2.Account)
	require.Equal(t, authtoken1.Token, authtoken2.Token)

	require.WithinDuration(t, authtoken1.CreatedAt, authtoken2.CreatedAt, time.Second)
	require.WithinDuration(t, authtoken1.LastUsed, authtoken2.LastUsed, time.Second)


}

func TestQueries_UpdateAuthTokenDate(t *testing.T) {
	authtoken := createRandomizedAuthToken(t)
	args := UpdateAuthTokenDateParams{
		Account:  authtoken.Account,
		LastUsed: time.Now(),
	}
	authtoken2, err := testQueries.UpdateAuthTokenDate(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, authtoken2)

	require.Equal(t, authtoken.ID, authtoken2.ID)
	require.Equal(t, authtoken.Account, authtoken2.Account)
	require.Equal(t, authtoken.Token, authtoken2.Token)

	require.WithinDuration(t, authtoken.CreatedAt, authtoken2.CreatedAt, time.Second)
	require.WithinDuration(t, args.LastUsed, authtoken2.LastUsed, time.Second)

}

func TestQueries_UpdateAuthTokenValue(t *testing.T) {
	authtoken := createRandomizedAuthToken(t)
	args := UpdateAuthTokenValueParams{
		Account: authtoken.Account,
		Token:   utils.GenerateAuthToken(),
	}
	authtoken2, err := testQueries.UpdateAuthTokenValue(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, authtoken2)

	require.Equal(t, authtoken.ID, authtoken2.ID)
	require.Equal(t, authtoken.Account, authtoken2.Account)
	require.Equal(t, args.Token, authtoken2.Token)

	require.WithinDuration(t, authtoken.CreatedAt, authtoken2.CreatedAt, time.Second)
	require.WithinDuration(t, authtoken.LastUsed, authtoken2.LastUsed, time.Second)

}

func TestQueries_DeleteAuthToken(t *testing.T) {
	authtoken := createRandomizedAuthToken(t)

	err := testQueries.DeleteAuthToken(context.Background(), authtoken.Account)
	require.NoError(t, err)

	authtoken2, err := testQueries.GetAuthTokenByAccount(context.Background(), authtoken.Account)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, authtoken2)
}
