package db

import (
	"context"
	"testing"

	"github.com/osisupermoses/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, from, to int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from,
		ToAccountID:   to,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testStore.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, from)
	require.Equal(t, transfer.ToAccountID, to)
	require.Equal(t, transfer.Amount, arg.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	// create random from and to accounts
	fromAcct := createRandomAccount(t)
	toAcct := createRandomAccount(t)

	createRandomTransfer(t, fromAcct.ID, toAcct.ID)
}

func TestListTransfers(t *testing.T) {
	// create random from and to accounts
	fromAcct := createRandomAccount(t)
	toAcct := createRandomAccount(t)

	for range 5 {
		createRandomTransfer(t, fromAcct.ID, toAcct.ID)
	}
	arg := ListTransfersParams{
		FromAccountID: fromAcct.ID,
		ToAccountID:   toAcct.ID,
		Limit:         3,
		Offset:        2,
	}
	transfers, err := testStore.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 3)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
		require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	}
}
