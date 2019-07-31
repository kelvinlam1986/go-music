package repositories

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var ErrorInvalidPassword = errors.New("Invalid Password")

type MusicContext struct {
	*gorm.DB
}

func NewMusicContext(typedb, con string) (*MusicContext, error) {
	db, err := gorm.Open(typedb, con)
	return &MusicContext{
		DB: db,
	}, err
}
