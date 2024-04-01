package app

import (
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
	fmt.Println(app.Repository.Product.GetProductByID(args[0]))

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
