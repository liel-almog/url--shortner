package repositories

import (
	"sync"

	"github.com/liel-almog/url-shortener/database"
)

type UrlRepository interface {
}

type urlRepositoryImpl struct {
	db *database.Sqlite
}

var (
	initUrlRepositoryOnce sync.Once
	urlRepository         *urlRepositoryImpl
)

func newUrlRepository() *urlRepositoryImpl {
	return &urlRepositoryImpl{
		db: database.GetDB(),
	}
}

func GetUrlRepository() UrlRepository {
	initUrlRepositoryOnce.Do(func() {
		urlRepository = newUrlRepository()
	})

	return urlRepository
}
