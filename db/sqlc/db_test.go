package db

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_WithTx(t *testing.T) {
	tx, err := testDb.Begin()
	require.NoError(t, err)
	require.NotEmpty(t, tx)

	tq := testQueries.WithTx(tx)

	require.NotEmpty(t, tq)

}
