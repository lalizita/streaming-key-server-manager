package service

import (
	"github.com/lalizita/streaming-key-server-manager/internal/model"
	"github.com/lalizita/streaming-key-server-manager/internal/repository"
)

type KeysService interface {
	GetStreamingKey(url, key string) (*model.Keys, error)
}

type keysService struct {
	keysRepository repository.KeysRepository
}

func NewKeysService(repo repository.KeysRepository) KeysService {
	return &keysService{
		keysRepository: repo,
	}
}

func (s *keysService) GetStreamingKey(url, key string) (*model.Keys, error) {
	return s.keysRepository.FindStreamKey(url, key)
}
