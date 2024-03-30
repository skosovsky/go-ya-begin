package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectSales(t *testing.T) {
	client := 208
	sales, err := selectSales(client)

	require.NoError(t, err)
	require.NotEmpty(t, sales)

	for _, sale := range sales {
		assert.NotEmpty(t, sale.Product)
		assert.NotEmpty(t, sale.Volume)
		assert.NotEmpty(t, sale.Date)
	}
}
