package app

import (
	"ecommerce/internal/helpers"
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
	"fmt"
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
	fmt.Println(*orders[0], *orders[1])

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

func (a *App) GetData(orderIDs []int) (*models.OrderPageDTO, error) {
	var data *models.OrderPageDTO
	orders, err := a.Repository.ProductOrder.GetProductOrder(orderIDs[0])
	if err != nil {
		return nil, err
	}
	productIDs := helpers.GetProductIdsFromStruct(orders)
	products, err := a.Repository.Product.GetProductsByIDs(productIDs)
	if err != nil {
		return nil, err
	}

	productShelve, err := a.Repository.ProductShelve.GetProductShelves(productIDs)
	if err != nil {
		return nil, err
	}

	shelveIDs := helpers.GetShelveIdsFromStruct(productShelve)

	shelves, err := a.Repository.Shelve.GetShelvesByIDs(shelveIDs)

	if err != nil {
		return nil, err
	}

	var ordersDTO models.OrderInfoDTO

}
