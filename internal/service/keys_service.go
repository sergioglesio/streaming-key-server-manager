package service

import (
	"github.com/sergioglesio/streaming-key-server-manager/internal/model"
	"github.com/sergioglesio/streaming-key-server-manager/internal/repository"
)

type IKeyService interface {
	AuthStreamingKey(name, key string) (*model.Keys, error)
}

type keysService struct {
	keysRepository repository.IKeysRepository
}

func NewKeysService(repo repository.IKeysRepository) IKeyService {
	return &keysService{
		keysRepository: repo,
	}
}

func (s *keysService) AuthStreamingKey(name, key string) (*model.Keys, error) {
	return s.keysRepository.FindStreamKey(name, key)
}
