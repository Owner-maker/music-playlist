package domain

import (
	"github.com/google/uuid"
	"time"
)

type Playlist struct {
	Songs DoublyLinkedList
}

type DoublyLinkedList struct {
	Start *Node
	End   *Node
}

type Node struct {
	next *Node
	prev *Node
	data Song
}

type Song struct {
	ID       uuid.UUID
	Name     string
	Duration time.Duration
}

type MusicService interface {
	Play()                   // Play начинает воспроизведение текущей песни
	Pause()                  // Pause приостанавливает воспроизведение
	AddSong(s Song)          // AddSong добавляет в конец плейлиста песню 	*
	RemoveSong(id uuid.UUID) // RemoveSong удаляют необходимую песню
	Next()                   // Next воспроизводит следующую песню 		    *
	Prev()                   // Prev воспроизводит предыдущую песню 		*
}
