package db

import (
	"context"
	"database/sql"
	"foscloud/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomizedNode(t *testing.T) Node {
	args := CreateNodeParams{
		ParentID: sql.NullInt64{},
		Name:     utils.RandomName(),
		Filesize: sql.NullInt64{},
		Depth:    sql.NullInt32{},
		Lineage:  sql.NullString{},
		Owner:    sql.NullInt64{},
	}

	node, err := testQueries.CreateNode(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, node)

	require.NotEmpty(t, node.ID)
	require.Equal(t, args.Name, node.Name)
	require.Equal(t, args.Filesize, node.Filesize)
	require.Equal(t, args.Depth, node.Depth)
	require.Equal(t, args.Lineage, node.Lineage)
	require.Equal(t, args.Owner, node.Owner)
	require.Equal(t, false, node.IsDir)
	return node
}

func TestQueries_CreateNode(t *testing.T) {
	createRandomizedNode(t)
}

func TestQueries_GetNode(t *testing.T) {
	node1 := createRandomizedNode(t)

	node2, err := testQueries.GetNode(context.Background(), node1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, node2)

	require.Equal(t, node1.ID, node2.ID)
	require.Equal(t, node1.IsDir, node2.IsDir)
	require.Equal(t, node1.Owner, node2.Owner)
	require.Equal(t, node1.Lineage, node2.Lineage)
	require.Equal(t, node1.Depth, node2.Depth)
	require.Equal(t, node1.Filesize, node2.Filesize)
	require.Equal(t, node1.Name, node2.Name)
	require.Equal(t, node1.ParentID, node2.ParentID)

	require.WithinDuration(t, node1.CreatedAt, node2.CreatedAt, time.Second)

}

func TestQueries_ListNodes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomizedNode(t)
	}
	args := ListNodesParams{
		Limit:  5,
		Offset: 5,
	}

	nodes, err := testQueries.ListNodes(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, nodes)

	require.Len(t, nodes, 5)

	for _, node := range nodes {
		require.NotEmpty(t, node)
	}

	count, err := testQueries.CountNodes(context.Background())
	require.NoError(t, err)

	args = ListNodesParams{
		Limit:  5,
		Offset: int32(count),
	}

	nodes, err = testQueries.ListNodes(context.Background(), args)
	require.NoError(t, err)
	require.Empty(t, nodes)

	// invalid query
	args = ListNodesParams{
		Limit:  -1,
		Offset: -1,
	}

	nodes, err = testQueries.ListNodes(context.Background(), args)
	require.Error(t, err)
	require.Empty(t, nodes)
}

func TestQueries_DeleteNode(t *testing.T) {
	node1 := createRandomizedNode(t)

	err := testQueries.DeleteNode(context.Background(), node1.ID)
	require.NoError(t, err)

	node2, err := testQueries.GetNode(context.Background(), node1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, node2)
}

func TestQueries_CountNodes(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomizedNode(t)
	}

	count, err := testQueries.CountNodes(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, int64(10))
}

