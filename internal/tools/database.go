package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database collections
type Auth struct {
	UserId int 
	Username string 
	Name string 
	Password string
}

type LoginDetails struct {
	UserId int 
	Username string 
	Name string 
	Password string
	Status int 
	CreatedAt string
}
type CoinDetails struct {
	Coins   int64
	Username string
}
type DatabaseInterface interface {
	GetAuth(token string) *Auth
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &dB{}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}