package usecase

import (
	"rest_api/model"
	"rest_api/repository"
)

type ItemUsecase struct {
	repository repository.ItemRepository
}

func NewItemUsecase(repo repository.ItemRepository) ItemUsecase {
	return ItemUsecase{
		repository: repo,
	}
}

func (iu *ItemUsecase) GetItems() ([]model.Item, error) {
	return iu.repository.GetItems()
}
