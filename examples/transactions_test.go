package main

import (
	"context"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/stretchr/testify/assert"
)

func TestTransactions(t *testing.T) {
	var (
		ctx  = context.TODO()
		repo = reltest.New()
	)

	/// [transactions]
	repo.ExpectTransaction(func(repo *reltest.Repository) {
		repo.ExpectUpdate(rel.Dec("stock")).ForType("main.Book")

		// mock process

		repo.ExpectTransaction(func(r *reltest.Repository) {
			repo.ExpectUpdateAny(rel.From("authors").Where(where.Eq("id", 0)), rel.Inc("popularity"))
			repo.ExpectUpdateAny(rel.From("publishers").Where(where.Eq("name", "")), rel.Inc("popularity"))
		})

		repo.ExpectUpdate(rel.Set("status", "paid")).ForType("main.Transaction")
	})
	/// [transactions]

	assert.Nil(t, Transactions(ctx, repo))
	repo.AssertExpectations(t)
}
