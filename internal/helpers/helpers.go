package helpers

import "ecommerce/internal/models"

func GetProductIdsFromStruct(productOrders []*models.ProductOrder) []int {
	var result []int

	for i := range productOrders {
		result = append(result, (*productOrders[i]).ProductID)
	}
	return result
}

func GetShelveIdsFromStruct(productOrders []*models.ProductShelve) []int {
	result := make([]int, 0)

	for i := range productOrders {
		result = append(result, productOrders[i].ShelveID)
	}
	return result
}
