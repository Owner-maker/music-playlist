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
	conv, err := schema.ConvToUploadList(s)
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
		return []domain.Song{}, err
	}
	defer file.Close()

	var playlist schema.DownloadedPlaylist
	err = json.NewDecoder(file).Decode(&playlist)
	if err != nil {
		return []domain.Song{}, err
	}

	return playlist.Data.Convert(), nil
}
