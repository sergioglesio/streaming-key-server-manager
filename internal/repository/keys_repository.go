package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/lalizita/streaming-key-server-manager/internal/model"
)

var (
	ErrQuery = errors.New("error find string")
)

type IKeysRepository interface {
	FindStreamKey(name, key string) (*model.Keys, error)
}

type KeysRepository struct {
	*sql.DB
}

func NewKeyRepository(db *sql.DB) IKeysRepository {
	return &KeysRepository{
		db,
	}
}

func (kr *KeysRepository) FindStreamKey(name, key string) (*model.Keys, error) {
	fmt.Println("============== looking for:", name, key)
	keys := &model.Keys{}
	row := kr.QueryRow(`SELECT * FROM "Lives" WHERE "name"=$1 AND "stream-key"=$2`, name, key)

	err := row.Scan(&keys.Name, &keys.Key)
	if err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &model.Keys{}, nil
		}

		return &model.Keys{}, ErrQuery
	}
	return keys, nil
}
