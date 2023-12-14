package main

import "testing"

func TestApp(t *testing.T) {
	storage := &InMemoryStorage{}

	first := storage.Add(1, "Илья")
	second := storage.Add(2, "Даник")
	third := storage.Add(3, "Артем")

	app := App{storage: storage}

	if app.CurrentDishWasher() != first {
		t.Fatal("текущий посудомой определен неправильно")
	}

	_ = app.IWashedDishes(first.ID)

	if app.CurrentDishWasher() != second {
		t.Fatal("текущий посудомой определен неправильно")
	}

	_ = app.IWashedDishes(second.ID)

	if app.CurrentDishWasher() != third {
		t.Fatal("текущий посудомой определен неправильно")
	}

	_ = app.IWashedDishes(third.ID)

	if app.CurrentDishWasher() != first {
		t.Fatal("текущий посудомой определен неправильно")
	}

	// TODO кейсы если человек помыл посуду вне очереди
}
