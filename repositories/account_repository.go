package repositories

import (
	"errors"
	"go-music/models"
	"go-music/utils"
)

type IAccountRepository interface {
	GetAccountByUserName(username string) (models.Account, error)
}

type AccountRepository struct {

}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (repo *AccountRepository) GetAccountByUserName(username string) (models.Account, error) {
	password := "10101986"
	err := utils.HashPassword(&password)
	if err != nil {
		return models.Account{}, errors.New("An error occured while hash password")
	}

	return models.Account{ Username: "minhlam", Password: password }, nil
}