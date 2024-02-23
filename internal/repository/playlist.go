package repository

import (
	"encoding/json"
	"music-playlist/internal/domain"
	"music-playlist/internal/repository/schema"
	"os"
)

type Playlist struct {
	filename string
}

func NewPlaylist(filename string) Playlist {
	return Playlist{filename: filename}
}

func (p Playlist) Upload(s []domain.Song) error {
	return nil
}

func (p Playlist) Download() ([]domain.Song, error) {
	file, err := os.Open(p.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var playlist schema.Playlist
	err = json.NewDecoder(file).Decode(&playlist)
	if err != nil {
		return nil, err
	}

	return schema.ConvertSongs(playlist.Data), nil
}
