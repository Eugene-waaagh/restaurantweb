package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/stretchr/testify/require"
)

// Helper function to create a random OrderCustomization
func addRandomCustomizationToOrderItem(t *testing.T, orderItemID int32, customizationValueID int32) Ordercustomization {
	arg := AddCustomizationToOrderItemParams{
		OrderItemID:         orderItemID,
		FoodCustomizationID: customizationValueID,
	}

	customization, err := testQueries.AddCustomizationToOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customization)

	require.Equal(t, arg.OrderItemID, customization.OrderItemID)
	require.Equal(t, arg.FoodCustomizationID, customization.FoodCustomizationID)

	return customization
}

func TestQueries_AddCustomizationToOrderItem(t *testing.T) {
	// Create dependencies
	category := createRandomFoodCategory(t)
	food := createRandomFood(t, category)
	customization := createRandomFoodCustomization(t, food)
	order := createRandomOrder(t)
	orderItem := createRandomOrderItem(t, order, food)

	// Add customization to order item
	addRandomCustomizationToOrderItem(t, orderItem.ID, customization.ID)
}

func TestQueries_GetOrderItemCustomizations(t *testing.T) {
	// Create dependencies
	category := createRandomFoodCategory(t)
	food := createRandomFood(t, category)
	customization := createRandomFoodCustomization(t, food)
	order := createRandomOrder(t)
	orderItem := createRandomOrderItem(t, order, food)

	// Add customization to order item
	addRandomCustomizationToOrderItem(t, orderItem.ID, customization.ID)

	// Retrieve customizations for the order item
	customizations, err := testQueries.GetOrderItemCustomizations(context.Background(), orderItem.ID)
	require.NoError(t, err)
	require.NotEmpty(t, customizations)

	for _, c := range customizations {
		require.NotEmpty(t, c.CustomizationType)
		require.NotEmpty(t, c.Value)
	}
}

func TestQueries_DeleteOrderItemCustomization(t *testing.T) {
	// Create dependencies
	category := createRandomFoodCategory(t)
	food := createRandomFood(t, category)
	customization := createRandomFoodCustomization(t, food)
	order := createRandomOrder(t)
	orderItem := createRandomOrderItem(t, order, food)

	// Add customization to order item
	orderCustomization := addRandomCustomizationToOrderItem(t, orderItem.ID, customization.ID)

	// Delete customization
	arg := DeleteOrderItemCustomizationParams{
		OrderItemID:         orderCustomization.OrderItemID,
		FoodCustomizationID: orderCustomization.FoodCustomizationID,
	}
	err := testQueries.DeleteOrderItemCustomization(context.Background(), arg)
	require.NoError(t, err)

	// Ensure customization is deleted
	customizations, err := testQueries.GetOrderItemCustomizations(context.Background(), orderItem.ID)
	require.NoError(t, err)
	require.Empty(t, customizations)
}
