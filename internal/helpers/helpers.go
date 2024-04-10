package helpers

import "ecommerce/internal/models"

func GetProductIdsFromStruct(productOrders []models.ProductOrder) []int {
	result := make([]int, 0, len(productOrders))

	for i := range productOrders {
		result = append(result, (productOrders[i]).ProductID)
	}
	return result
}
