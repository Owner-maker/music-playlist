package schema

import (
	"encoding/json"
	"github.com/google/uuid"
	"music-playlist/internal/domain"
	"time"
)

type DownloadedPlaylist struct {
	Data DownloadedSongList `json:"data"`
}

type DownloadedSong struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Duration Duration  `json:"duration"`
}

type Duration struct {
	Duration time.Duration
}

func (md *Duration) UnmarshalJSON(b []byte) error {
	var durationString string
	if err := json.Unmarshal(b, &durationString); err != nil {
		return err
	}

	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return err
	}

	md.Duration = duration

	return nil
}

type DownloadedSongList []DownloadedSong

func (l DownloadedSongList) Convert() []domain.Song {
	res := make([]domain.Song, len(l))
	for i, song := range l {
		res[i] = song.convert()
	}

	return res
}

func (s DownloadedSong) convert() domain.Song {
	return domain.Song{
		ID:       s.ID,
		Name:     s.Name,
		Duration: s.Duration.Duration,
	}
}

type UploadPlaylist struct {
	Data []UploadSong `json:"data"`
}

type UploadSong struct {
	ID       uuid.UUID
	Name     string
	Duration string
}

func ConvToUploadList(s []domain.Song) (UploadPlaylist, error) {
	data := make([]UploadSong, len(s))
	for i, song := range s {
		conv, err := convToUpload(song)
		if err != nil {
			return UploadPlaylist{}, err
		}

		data[i] = conv
	}

	return UploadPlaylist{Data: data}, nil
}

func convToUpload(s domain.Song) (UploadSong, error) {
	conv, err := time.ParseDuration(s.Duration.String())
	if err != nil {
		return UploadSong{}, err
	}

	res := UploadSong{
		ID:       s.ID,
		Name:     s.Name,
		Duration: conv.String(),
	}

	return res, nil
}
