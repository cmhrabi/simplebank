package db

import "context"

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{arg.FromAccountID, arg.ToAccountID, arg.Amount})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreatEntry(ctx, CreatEntryParams{arg.FromAccountID, -arg.Amount})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreatEntry(ctx, CreatEntryParams{arg.ToAccountID, arg.Amount})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{-arg.Amount, arg.FromAccountID})
			if err != nil {
				return err
			}

			result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{arg.Amount, arg.ToAccountID})
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{arg.Amount, arg.ToAccountID})
			if err != nil {
				return err
			}

			result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{-arg.Amount, arg.FromAccountID})
			if err != nil {
				return err
			}
		}

		return err
	})

	return result, err
}
