package app

import (
	"ecommerce/internal/helpers"
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
	"fmt"
	"log"
	"os"
	"strconv"
)

type App struct {
	Repository *repository.Repository
}

func NewApp(repository *repository.Repository) *App {
	return &App{
		Repository: repository,
	}
}

func (app *App) Start() error {
	args, err := validate(os.Args[1:])

	if err != nil {
		return err
	}
	orders, err := app.GetData(args)
	if err != nil {
		log.Fatal(err)
	}

	app.outputData(orders)

	return nil
}

func validate(args []string) ([]int, error) {
	var argsInt []int

	for _, a := range args {
		number, err := strconv.Atoi(a)
		if err != nil {
			return nil, err
		}
		argsInt = append(argsInt, number)
	}
	return argsInt, nil
}

func (a *App) GetData(orderIDs []int) (map[string][]models.OrderInfoDTO, error) {
	// Using a map to prevent repetitive database queries for the same product
	productCache := make(map[int]models.Product)
	shelfCache := make(map[int]*models.Shelf)

	var data = make(map[string][]models.OrderInfoDTO)

	// Fetch product orders once
	productOrders, err := a.Repository.ProductOrder.GetProductOrdersByOrderIDs(orderIDs)
	if err != nil {
		return nil, fmt.Errorf("GET PRODUCTS BY ID ERR: %w", err)
	}

	// Map product IDs to product orders
	productIDToOrders := make(map[int][]models.ProductOrder)
	for _, po := range productOrders {
		productIDToOrders[po.ProductID] = append(productIDToOrders[po.ProductID], po)
	}

	// Fetch product shelves once
	productShelves, err := a.Repository.ProductShelve.GetProductShelvesByProductIDs(helpers.GetProductIdsFromStruct(productOrders))
	if err != nil {
		return nil, fmt.Errorf("GET PRODUCT SHELVES ERR: %w", err)
	}

	for _, ps := range productShelves {
		shelve, ok := shelfCache[ps.ShelveID]
		if !ok {
			shelve, err = a.Repository.Shelve.GetShelfByID(ps.ShelveID)
			if err != nil {
				return nil, fmt.Errorf("GET SHELF ERR: %w", err)
			}
			shelfCache[ps.ShelveID] = shelve
		}

		orders := productIDToOrders[ps.ProductID]
		for _, po := range orders {
			product, ok := productCache[po.ProductID]
			if !ok {
				product, err = a.Repository.Product.GetProductByID(po.ProductID)
				if err != nil {
					return nil, fmt.Errorf("GET Product ERR: %w", err)
				}
				productCache[po.ProductID] = product
			}

			data[shelve.Name] = append(data[shelve.Name], models.OrderInfoDTO{Product: product, ProductOrder: po})
		}
	}

	return data, nil
}

func (a *App) outputData(data map[string][]models.OrderInfoDTO) {

	for k, v := range data {
		fmt.Printf("Стеллаж : %s\n", k)
		for j := range v {
			fmt.Printf("Заказ : %d, %s, %d штук \n", v[j].ProductOrder.OrderID, v[j].Product.Name, v[j].ProductOrder.Quantity)
		}
	}
}
