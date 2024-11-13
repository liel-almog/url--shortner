package repositories

import (
	"database/sql"
	"sync"

	"github.com/liel-almog/url-shortener/database"
	"github.com/liel-almog/url-shortener/errors/apperrors"
)

type UrlRepository interface {
	FindOriginalUrl(shortURL string) (string, error)
}

type urlRepositoryImpl struct {
	sqlite *database.Sqlite
}

var (
	initUrlRepositoryOnce sync.Once
	urlRepository         *urlRepositoryImpl
)

func newUrlRepository() *urlRepositoryImpl {
	return &urlRepositoryImpl{
		sqlite: database.GetDB(),
	}
}

func GetUrlRepository() UrlRepository {
	initUrlRepositoryOnce.Do(func() {
		urlRepository = newUrlRepository()
	})

	return urlRepository
}

func (repo *urlRepositoryImpl) FindOriginalUrl(shortURL string) (string, error) {
	var originalURL string

	row := repo.sqlite.Db.QueryRow("SELECT original_url FROM urls WHERE short_url = ?", shortURL)
	if err := row.Scan(&originalURL); err != nil {
		if err == sql.ErrNoRows {
			return "", apperrors.ErrUrlNotFound
		}

		return "", err
	}

	return originalURL, nil
}
