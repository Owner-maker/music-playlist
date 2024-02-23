package service

import (
	"github.com/google/uuid"
	"log/slog"
	"music-playlist/internal/domain"
	"music-playlist/pkg/logger/sl"
)

type Playlist struct {
	cache *domain.DoublyLinkedList
	repo  domain.MusicRepository
}

func NewPlaylist(repo domain.MusicRepository) (*Playlist, error) {
	p := &Playlist{repo: repo}
	err := p.initCache()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s Playlist) initCache() error {
	_, err := s.repo.Download()
	if err != nil {
		slog.Error("failed to download songs from file", sl.Err(err))
		return err
	}

	cache := domain.NewDoublyLinkedList()
	//cache.AppendMany(songs...)
	s.cache = cache

	return nil
}

func (s Playlist) Add(data domain.Song) error {
	return nil
}

func (s Playlist) AddMany(data []domain.Song) error {
	// TODO transaction

	for _, v := range data {
		if err := s.Add(v); err != nil {
			return err
		}
	}

	return nil
}

func (s Playlist) Get(id uuid.UUID) (domain.Song, error) {
	//TODO implement me
	panic("implement me")
}

func (s Playlist) GetAll() ([]domain.Song, error) {
	//TODO implement me
	panic("implement me")
}

func (s Playlist) Update(data domain.Song) error {
	//TODO implement me
	panic("implement me")
}

func (s Playlist) Remove(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Playlist) Play() error {
	//TODO implement me
	panic("implement me")
}

func (s Playlist) Pause() error {
	//TODO implement me
	panic("implement me")
}

func (s Playlist) Next() error {
	//TODO implement me
	panic("implement me")
}

func (s Playlist) Prev() error {
	//TODO implement me
	panic("implement me")
}
