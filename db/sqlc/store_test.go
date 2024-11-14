package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTrx(t *testing.T) {
	store := NewStore(testDB)

	account := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5

	errs := make(chan error)
	results := make(chan TransferTxResult)

	amount := int64(10)
	for i := 0; i < n; i++ {
		go func() {

			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		results := <-results
		require.NotEmpty(t, results)

		// check transfer

		transfer := results.Transfer
		require.Equal(t, account.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
	}
}
