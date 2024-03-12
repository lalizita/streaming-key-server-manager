package service

import (
	"github.com/lalizita/streaming-key-server-manager/internal/model"
	"github.com/lalizita/streaming-key-server-manager/internal/repository"
)

type KeysService interface {
	AuthStreamingKey(name, key string) (*model.Keys, error)
}

type keysService struct {
	keysRepository repository.KeysRepository
}

func NewKeysService(repo repository.KeysRepository) KeysService {
	return &keysService{
		keysRepository: repo,
	}
}

func (s *keysService) AuthStreamingKey(name, key string) (*model.Keys, error) {
	return s.keysRepository.FindStreamKey(name, key)
}
