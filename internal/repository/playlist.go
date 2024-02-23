package repository

import (
	"encoding/json"
	"github.com/google/uuid"
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
	s = []domain.Song{
		{
			ID:       uuid.UUID{23},
			Name:     "asdasdasd",
			Duration: 1123123,
		},
	}

	conv, err := schema.ConvertSchemaList(s)
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(conv, " ", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(p.filename, bytes, 0666); err != nil {
		return err
	}

	return nil
}

func (p Playlist) Download() ([]domain.Song, error) {
	file, err := os.Open(p.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var playlist schema.PlaylistDomain
	err = json.NewDecoder(file).Decode(&playlist)
	if err != nil {
		return nil, err
	}

	return schema.ConvertDomainPlaylist(playlist.Data), nil
}
