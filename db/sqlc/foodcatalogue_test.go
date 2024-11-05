package db

import (
	"context"
	"database/sql"
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

func TestQueries_GetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
}

func TestQueries_UpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:   category1.ID,
		Name: util.RandomCategory(),
	}

	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Name, category2.Name)
}

func TestQueries_DeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestQueries_ListCategory(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCategory(t)
	}

}
