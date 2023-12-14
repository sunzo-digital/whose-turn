package main

import (
	"fmt"
)

var _ Storage = &InMemoryStorage{}

type InMemoryStorage struct {
	first *DishWasher
	head  *DishWasher
}

func (s *InMemoryStorage) Current() *DishWasher {
	return s.head
}

func (s *InMemoryStorage) SetCurrent(washer *DishWasher) {
	if washer == nil {
		panic("trying to set nil washer")
	}

	s.head = washer
}

func (s *InMemoryStorage) Add(ID int, name string) *DishWasher {
	washer := &DishWasher{
		ID:   ID,
		Name: name,
	}

	if s.first == nil {
		// cписок пуст, добавляем первый элемент
		s.first = washer
		s.head = washer
		return washer
	}

	if s.first.Next == nil {
		// добавляем второй элемент, создадим связь с первым
		washer.Next = s.head
		s.head.Next = washer

		return washer
	}

	// добавляем нового slave в конец списка
	latestDishWasher := s.head

	for latestDishWasher.Next != s.first {
		latestDishWasher = latestDishWasher.Next
	}

	washer.Next = s.first
	latestDishWasher.Next = washer

	return washer
}

func (s *InMemoryStorage) GetByID(ID int) (*DishWasher, error) {
	washer := s.first

	for washer != nil {
		if washer.ID == ID {
			return washer, nil
		}

		washer = washer.Next
	}

	return nil, fmt.Errorf("посудомойщик с ID %d не найден", ID)
}
