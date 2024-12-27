package db

import (
	"context"
	"github.com/eugene-waaagh/restaurantweb/util"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/stretchr/testify/require"
)

// Helper function to create a random food category
func createRandomFoodCategory(t *testing.T) Foodcatalogue {
	name := util.RandomCategory()
	category, err := testQueries.CreateCategory(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	return category
}

// Helper function to create random food
func createRandomFood(t *testing.T, category Foodcatalogue) Food {
	arg := CreateFoodParams{
		Name:        util.RandomString(10),
		Description: util.RandomString(20),
		Price:       util.RandomInt(1, 1000),
		CategoryID:  category.ID,
	}

	food, err := testQueries.CreateFood(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, food)

	require.Equal(t, arg.Name, food.Name)
	require.Equal(t, arg.Description, food.Description)
	require.Equal(t, arg.Price, food.Price)
	require.Equal(t, arg.CategoryID, food.CategoryID)
	require.NotZero(t, food.ID)

	return food
}

func TestQueries_CreateFood(t *testing.T) {
	category := createRandomFoodCategory(t)
	createRandomFood(t, category)
}

func TestQueries_GetFood(t *testing.T) {
	category := createRandomFoodCategory(t)
	food1 := createRandomFood(t, category)

	food2, err := testQueries.GetFood(context.Background(), food1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, food2)

	require.Equal(t, food1.ID, food2.ID)
	require.Equal(t, food1.Name, food2.Name)
	require.Equal(t, food1.Description, food2.Description)
	require.Equal(t, food1.Price, food2.Price)
	require.Equal(t, food1.CategoryID, food2.CategoryID)
}

func TestQueries_ListFood(t *testing.T) {
	category := createRandomFoodCategory(t)
	for i := 0; i < 10; i++ {
		createRandomFood(t, category)
	}

	foods, err := testQueries.ListFood(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, foods)

	for _, food := range foods {
		require.NotEmpty(t, food)
		require.Equal(t, category.ID, food.CategoryID)
	}
}

func TestQueries_UpdateFood(t *testing.T) {
	category := createRandomFoodCategory(t)
	food1 := createRandomFood(t, category)

	arg := UpdateFoodParams{
		ID:          food1.ID,
		Name:        util.RandomString(10),
		Description: util.RandomString(20),
		Price:       util.RandomInt(1, 1000),
		CategoryID:  category.ID,
	}

	food2, err := testQueries.UpdateFood(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, food2)

	require.Equal(t, food1.ID, food2.ID)
	require.Equal(t, arg.Name, food2.Name)
	require.Equal(t, arg.Description, food2.Description)
	require.Equal(t, arg.Price, food2.Price)
	require.Equal(t, arg.CategoryID, food2.CategoryID)
}
