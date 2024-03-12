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

func NewPlaylist(repo domain.MusicRepository, cache *domain.DoublyLinkedList) *Playlist {
	return &Playlist{repo: repo, cache: cache}
}

func InitCache(repo domain.MusicRepository) *domain.DoublyLinkedList {
	data, err := repo.Download()
	if err != nil {
		slog.Warn("failed to download songs from file", sl.Err(err))
	}

	cache := domain.NewDoublyLinkedList()
	cache.AppendMany(data...)
	return cache
}

// Save TODO проверить
func (s Playlist) Save() error {
	data := s.cache.GetAll()
	if err := s.repo.Upload(data); err != nil {
		return err
	}

	return nil
}

func (s Playlist) Add(data *domain.Song) {
	s.cache.Append(data)
}

func (s Playlist) AddMany(data []*domain.Song) {
	for _, v := range data {
		s.Add(v)
	}
}

func (s Playlist) Get() domain.Info {
	return s.cache.Get()
}

func (s Playlist) GetAll() []domain.Info {
	//TODO implement me
	panic("implement me")
}

func (s Playlist) Update(data *domain.Info) error {
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
