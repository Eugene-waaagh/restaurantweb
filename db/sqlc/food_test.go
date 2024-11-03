package db

import (
	"context"
	"github.com/eugene-waaagh/restaurantweb/util"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Foodcatalogue {
	name := util.RandomCategory()

	category, err := testQueries.CreateCategory(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, name, category.Name)
	require.NotZero(t, category.ID)

	return category
}

func TestQueries_CreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
}
