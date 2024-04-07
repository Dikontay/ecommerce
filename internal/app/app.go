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

	var data = make(map[string][]models.OrderInfoDTO)

	productOrders, err := a.Repository.ProductOrder.GetProductOrdersByOrderIDs(orderIDs)

	if err != nil {
		fmt.Println("GET PRODUCTS BY ID ERR")
		return nil, err
	}
	productIDs := helpers.GetProductIdsFromStruct(productOrders)

	if err != nil {

		return nil, err
	}

	productShelves, err := a.Repository.ProductShelve.GetProductShelvesByProductIDs(productIDs)

	if err != nil {

		return nil, err
	}

	for i := range productShelves {
		shelve, err := a.Repository.Shelve.GetShelfByID(productShelves[i].ShelveID)

		if err != nil {

			return nil, nil
		}

		product, err := a.Repository.Product.GetProductByID(productShelves[i].ProductID)
		if err != nil {
			fmt.Println("GET Product ERR")
			return nil, err
		}
		for j := range productOrders {
			productOrder, err := a.Repository.ProductOrder.GetProductOrderByOrderID(productOrders[j].OrderID)
			if err != nil {
				return nil, err
			}
			data[shelve.Name] = append(data[shelve.Name], models.OrderInfoDTO{Product: product, ProductOrder: productOrder})
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
