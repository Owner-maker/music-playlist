package schema

import (
	"encoding/json"
	"github.com/google/uuid"
	"music-playlist/internal/domain"
	"time"
)

type PlaylistDomain struct {
	Data []SongParseDomain `json:"data"`
}

type SongParseDomain struct {
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

func ConvertDomainPlaylist(s []SongParseDomain) []domain.Song {
	res := make([]domain.Song, len(s))
	for i, song := range s {
		res[i] = song.convertDomain()
	}

	return res
}

func (s SongParseDomain) convertDomain() domain.Song {
	return domain.Song{
		ID:       s.ID,
		Name:     s.Name,
		Duration: s.Duration.Duration,
	}
}

type PlaylistSchema struct {
	Data []SongParseSchema `json:"data"`
}

type SongParseSchema struct {
	ID       uuid.UUID
	Name     string
	Duration time.Duration
}

func ConvertSchemaList(s []domain.Song) (PlaylistSchema, error) {
	data := make([]SongParseSchema, len(s))
	for i, song := range s {
		conv, err := convertSchema(song)
		if err != nil {
			return PlaylistSchema{}, err
		}

		data[i] = conv
	}

	return PlaylistSchema{Data: data}, nil
}

func convertSchema(s domain.Song) (SongParseSchema, error) {
	conv, err := time.ParseDuration(s.Duration.String())
	if err != nil {
		return SongParseSchema{}, err
	}

	res := SongParseSchema{
		ID:       s.ID,
		Name:     s.Name,
		Duration: conv,
	}

	return res, nil
}
