package db

import (
	"context"
	"github.com/eugene-waaagh/restaurantweb/util"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/stretchr/testify/require"
)

// Helper function to create a random order
func createRandomOrder(t *testing.T) Order {
	arg := CreateOrderParams{
		OrderDate:  "2024-12-27", // Use dynamic date if needed
		TotalPrice: util.RandomInt(1000, 5000),
		Status:     "pending",
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.NotZero(t, order.ID)
	require.Equal(t, arg.OrderDate, order.OrderDate)
	require.Equal(t, arg.TotalPrice, order.TotalPrice)
	require.Equal(t, arg.Status, order.Status)

	return order
}

// Other shared helper functions
func createRandomOrderItem(t *testing.T, order Order, food Food) Orderitem {
	arg := CreateOrderItemParams{
		OrderID:   order.ID,
		FoodID:    food.ID,
		Quantity:  util.RandomInt(1, 5),
		ItemPrice: util.RandomInt(100, 500),
	}

	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderItem)

	require.Equal(t, arg.OrderID, orderItem.OrderID)
	require.Equal(t, arg.FoodID, orderItem.FoodID)
	require.Equal(t, arg.Quantity, orderItem.Quantity)
	require.Equal(t, arg.ItemPrice, orderItem.ItemPrice)
	require.NotZero(t, orderItem.ID)

	return orderItem
}
