package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/lalizita/streaming-key-server-manager/internal/model"
)

var (
	QueryErr = errors.New("Error query key")
)

type KeysRepository interface {
	FindStreamKey(url, key string) (*model.Keys, error)
}

type keysRepository struct {
	*sql.DB
}

func NewKeysRepository(db *sql.DB) KeysRepository {
	return &keysRepository{
		db,
	}
}

func (r *keysRepository) FindStreamKey(url, key string) (*model.Keys, error) {
	fmt.Printf("============ %s : %s", url, key)
	keys := &model.Keys{}
	row := r.QueryRow(`SELECT * FROM "Keys" WHERE "Url"=$1 AND "Key"=$2`, url, key)

	err := row.Scan(&keys.Url, &keys.Key)
	if err != nil {
		log.Error(err.Error())
		return &model.Keys{}, QueryErr
	}

	log.Info("Here is the keys", keys)
	return keys, nil
}
