package db

import (
	"context"
	"github.com/eugene-waaagh/restaurantweb/util"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/stretchr/testify/require"
)

// Helper function to create a random food customization
func createRandomFoodCustomization(t *testing.T, food Food) Foodcustomization {
	arg := CreateFoodCustomizationParams{
		FoodID:            food.ID,
		CustomizationType: util.RandomString(8),
		Value:             util.RandomString(10),
	}

	customization, err := testQueries.CreateFoodCustomization(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customization)

	require.Equal(t, arg.FoodID, customization.FoodID)
	require.Equal(t, arg.CustomizationType, customization.CustomizationType)
	require.Equal(t, arg.Value, customization.Value)
	require.NotZero(t, customization.ID)

	return customization
}

func TestQueries_CreateFoodCustomization(t *testing.T) {
	category := createRandomFoodCategory(t)
	food := createRandomFood(t, category)
	createRandomFoodCustomization(t, food)
}

func TestQueries_ListCustomizationsByFood(t *testing.T) {
	category := createRandomFoodCategory(t)
	food := createRandomFood(t, category)

	for i := 0; i < 5; i++ {
		createRandomFoodCustomization(t, food)
	}

	customizations, err := testQueries.ListCustomizationsByFood(context.Background(), food.ID)
	require.NoError(t, err)
	require.NotEmpty(t, customizations)

	for _, customization := range customizations {
		require.NotEmpty(t, customization)
		require.Equal(t, food.ID, customization.FoodID)
	}
}

func TestQueries_UpdateFoodCustomization(t *testing.T) {
	category := createRandomFoodCategory(t)
	food := createRandomFood(t, category)
	customization1 := createRandomFoodCustomization(t, food)

	arg := UpdateFoodCustomizationParams{
		ID:                customization1.ID,
		CustomizationType: util.RandomString(8),
		Value:             util.RandomString(10),
	}

	customization2, err := testQueries.UpdateFoodCustomization(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customization2)

	require.Equal(t, customization1.ID, customization2.ID)
	require.Equal(t, food.ID, customization2.FoodID)
	require.Equal(t, arg.CustomizationType, customization2.CustomizationType)
	require.Equal(t, arg.Value, customization2.Value)
}

func TestQueries_DeleteFoodCustomization(t *testing.T) {
	category := createRandomFoodCategory(t)
	food := createRandomFood(t, category)
	customization := createRandomFoodCustomization(t, food)

	err := testQueries.DeleteFoodCustomization(context.Background(), customization.ID)
	require.NoError(t, err)

	// Ensure the customization no longer exists
	customizations, err := testQueries.ListCustomizationsByFood(context.Background(), food.ID)
	require.NoError(t, err)

	for _, c := range customizations {
		require.NotEqual(t, customization.ID, c.ID)
	}
}
