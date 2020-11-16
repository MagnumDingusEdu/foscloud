package db

import (
	"context"
	"database/sql"
	"foscloud/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomizedLink(t *testing.T, node *Node) Link {
	args := CreateLinkParams{
		Node:     node.ID,
		Link:     utils.RandomLink(),
		Password: utils.RandomString(8),
	}

	link, err := testQueries.CreateLink(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, link)

	require.NotEmpty(t, link.ID)
	require.NotEmpty(t, link.CreatedAt)
	require.Equal(t, int32(0), link.Clicks)
	require.Equal(t, args.Node, link.Node)
	require.Equal(t, args.Link, link.Link)
	require.Equal(t, args.Password, link.Password)

	return link
}

func TestQueries_CreateLink(t *testing.T) {
	node := createRandomizedNode(t)
	createRandomizedLink(t, &node)

}

func TestQueries_GetLink(t *testing.T) {
	node := createRandomizedNode(t)
	link1 := createRandomizedLink(t, &node)

	link2, err := testQueries.GetLink(context.Background(), link1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, link2)

	require.Equal(t, link1.ID, link2.ID)
	require.Equal(t, link1.Link, link2.Link)
	require.Equal(t, link1.Clicks, link2.Clicks)
	require.Equal(t, link1.Password, link2.Password)
	require.Equal(t, link1.Node, link2.Node)

	require.WithinDuration(t, link1.CreatedAt, link2.CreatedAt, time.Second)

}

func TestQueries_UpdateLink(t *testing.T) {
	node := createRandomizedNode(t)
	link1 := createRandomizedLink(t, &node)

	args := UpdateLinkParams{
		ID:       link1.ID,
		Link:     utils.RandomLink(),
		Password: utils.RandomString(8),
	}
	link2, err := testQueries.UpdateLink(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, link2)

	require.Equal(t, link1.ID, link2.ID)
	require.Equal(t, args.Link, link2.Link)
	require.Equal(t, link1.Clicks, link2.Clicks)
	require.Equal(t, args.Password, link2.Password)
	require.Equal(t, link1.Node, link2.Node)

	require.WithinDuration(t, link1.CreatedAt, link2.CreatedAt, time.Second)

}

func TestQueries_DeleteLink(t *testing.T) {
	node := createRandomizedNode(t)
	link1 := createRandomizedLink(t, &node)

	err := testQueries.DeleteLink(context.Background(), link1.ID)
	require.NoError(t, err)

	link2, err := testQueries.GetLink(context.Background(), link1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, link2)
}

func TestQueries_ListLinks(t *testing.T) {
	node := createRandomizedNode(t)
	for i:=0; i<10; i++ {
		createRandomizedLink(t, &node)
	}

	args := ListLinksParams{
		Limit:  5,
		Offset: 5,
	}

	links, err := testQueries.ListLinks(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, links)

	require.Len(t, links, 5)

	for _, link := range links {
		require.NotEmpty(t, link)
	}
	count, err := testQueries.CountLinks(context.Background())
	require.NoError(t, err)

	args = ListLinksParams{
		Limit:  5,
		Offset: int32(count),
	}

	links, err = testQueries.ListLinks(context.Background(), args)
	require.NoError(t, err)
	require.Empty(t, links)

	// invalid query
	args = ListLinksParams{
		Limit:  -1,
		Offset: -1,
	}

	links, err = testQueries.ListLinks(context.Background(), args)
	require.Error(t, err)
	require.Empty(t, links)
}

func TestQueries_CountLinks(t *testing.T) {
	node := createRandomizedNode(t)
	for i:=0; i<10; i++ {
		createRandomizedLink(t, &node)
	}

	count, err := testQueries.CountLinks(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, int64(10))
}