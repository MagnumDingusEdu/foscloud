package db

import (
	"context"
	"database/sql"
	"fmt"
	"foscloud/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomizedNode(t *testing.T) Node {
	return createSpecificNode(t, 1)
}

func createSpecificNode(t *testing.T, parentId int64) Node {
	args := CreateNodeParams{
		ParentID: sql.NullInt64{Int64: parentId, Valid: true},
		Name:     utils.RandomName(),
		IsDir:    false,
		Filesize: sql.NullInt64{Int64: utils.RandomFilesize(), Valid: true},
		Owner:    sql.NullInt64{},
	}

	node, err := testQueries.CreateNode(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, node)

	require.NotEmpty(t, node.ID)
	require.Equal(t, args.Name, node.Name)
	require.Equal(t, args.Filesize, node.Filesize)
	require.Equal(t, args.ParentID, node.ParentID)
	require.Equal(t, args.IsDir, node.IsDir)
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
	require.NotEmpty(t, node2.Lineage) // Lineage is generated later
	require.NotEmpty(t, node2.Depth)   // Depth is generated later
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

func TestQueries_ListChildNodes(t *testing.T) {
	var err error
	parentNode := createSpecificNode(t, 1)
	parentNode, err = testQueries.GetNode(context.Background(), parentNode.ID)

	require.NoError(t, err)
	require.NotEmpty(t, parentNode)

	for i := 0; i < 10; i++ {
		createSpecificNode(t, parentNode.ID)
	}

	children, err := testQueries.ListChildNodes(context.Background(), sql.NullInt64{
		Int64: parentNode.ID,
		Valid: true,
	})

	require.NoError(t, err)
	require.Len(t, children, 10)

	for _, child := range children {
		require.NotEmpty(t, child)

		require.Equal(t, parentNode.Depth.Int32+1, child.Depth.Int32)
		require.Equal(t, fmt.Sprintf("%v%v/", parentNode.Lineage.String, child.ID), child.Lineage.String)
	}
}
