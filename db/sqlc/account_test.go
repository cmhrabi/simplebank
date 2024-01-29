package db

import (
	"context"
	"testing"

	"github.com/cmhrabi/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func CreateNewAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateNewAccount(t)
}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateNewAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: int64(utils.RandomInt(int(account1.Balance)+1, 1500)),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, arg.Balance, account2.Balance)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateNewAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Balance, account2.Balance)
}

func TestListAccount(t *testing.T) {
	CreateNewAccount(t)
}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateNewAccount(t)
	err1 := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err1)

	account2, err2 := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err2)
	require.Empty(t, account2)
}
