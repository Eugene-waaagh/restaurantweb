package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestQueries_GetFullOrderDetails(t *testing.T) {
	// Step 1: Create dependencies
	category := createRandomFoodCategory(t)
	food := createRandomFood(t, category)
	order := createRandomOrder(t)
	orderItem := createRandomOrderItem(t, order, food) // Required to link order details
	createRandomFoodCustomization(t, food)             // Optional for customization testing

	// Step 2: Fetch full order details
	orderDetails, err := testQueries.GetFullOrderDetails(context.Background(), order.ID)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetails)

	// Step 3: Validate details
	for _, detail := range orderDetails {
		require.Equal(t, order.ID, detail.OrderID)
		require.Equal(t, food.Name, detail.FoodName)
		require.Equal(t, orderItem.ID, detail.OrderItemID)
	}
}
