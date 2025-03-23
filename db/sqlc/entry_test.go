package db

import (
	"context"
	"testing"
	"time"

	"github.com/osisupermoses/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, accountID int64) Entry {
	account := createRandomAccount(t)
	require.NotEmpty(t, account)

	var accID int64
	if accountID != 0 {
		accID = accountID
	} else {
		accID = account.ID
	}
	entryArg := CreateEntryParams{
		AccountID: accID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testStore.CreateEntry(context.Background(), entryArg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	if accountID == 0 {
		require.Equal(t, account.ID, entry.AccountID)
	}

	require.Equal(t, entryArg.AccountID, entry.AccountID)
	require.Equal(t, entryArg.Amount, entry.Amount)

	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t, 0)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t, 0)
	entry2, err := testStore.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	// create new account
	acct := createRandomAccount(t)

	// create 5 new entries with the same accountID
	for range 5 {
		createRandomEntry(t, acct.ID)
	}

	arg := ListEntriesParams{
		AccountID: acct.ID,
		Limit:     3,
		Offset:    2,
	}
	entries, err := testStore.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 3)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, acct.ID)
	}
}
