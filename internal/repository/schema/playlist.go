package schema

import (
	"encoding/json"
	"github.com/google/uuid"
	"music-playlist/internal/domain"
	"time"
)

type Playlist struct {
	Data []Song `json:"data"`
}

type Song struct {
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

func (s Song) Convert() domain.Song {
	return domain.Song{
		ID:       s.ID,
		Name:     s.Name,
		Duration: s.Duration.Duration,
	}
}

func ConvertSongs(songs []Song) []domain.Song {
	domainSongs := make([]domain.Song, len(songs))
	for i, song := range songs {
		domainSongs[i] = song.Convert()
	}
	return domainSongs
}
