package main

type App struct {
	storage Storage
}

type Storage interface {
	Current() *DishWasher
	SetCurrent(washer *DishWasher)
	Add(ID int, name string) *DishWasher
	GetByID(ID int) (*DishWasher, error)
}

type DishWasher struct {
	ID     int
	Name   string
	Points int
	Next   *DishWasher
}

func (app App) CurrentDishWasher() *DishWasher {
	current := app.storage.Current()
	return current
}

func (app App) IWashedDishes(ID int) error {
	washer, err := app.storage.GetByID(ID)
	if err != nil {
		return err
	}

	if app.storage.Current() != washer {
		// посудомой - ровный фраер и помыл посуду вне очереди
		// текущий посудомой не меняется
		washer.Points++

		return nil
	}

	app.storage.SetCurrent(washer.Next)

	return nil
}
