package repositories

import (
	"database/sql"
	"sync"

	"github.com/liel-almog/url-shortener/database"
	"github.com/liel-almog/url-shortener/errors/apperrors"
	"github.com/liel-almog/url-shortener/models"
)

type UrlRepository interface {
	FindOriginalUrl(shortURL string) (string, error)
	InsertShortenUrl(url models.Url) error
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
	var originalUrl string

	row := repo.sqlite.Db.QueryRow("SELECT original_url FROM urls WHERE short_url = ?", shortURL)
	if err := row.Scan(&originalUrl); err != nil {
		if err == sql.ErrNoRows {
			return "", apperrors.ErrUrlNotFound
		}

		return "", err
	}

	return originalUrl, nil
}

func (repo *urlRepositoryImpl) InsertShortenUrl(url models.Url) error {
	_, err := repo.sqlite.Db.Exec("INSERT INTO urls (short_url, original_url) VALUES (?, ?)", url.ShortUrl, url.OriginalUrl)
	if err != nil {
		return err
	}

	return nil
}
