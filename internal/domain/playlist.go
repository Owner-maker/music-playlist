package domain

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type Song struct {
	Info Info
	Meta Meta
}

type Meta struct {
	Mu        sync.Mutex
	PlayedFor time.Duration
	IsPlaying bool
}

type Info struct {
	ID       uuid.UUID
	Name     string
	Duration time.Duration
}

// Сохранение данных будет происходить при graceful shutdown

type MusicService interface {
	Add(data *Song)            // Add добавляет в конец плейлиста песню 	 		TODO конкурентный доступ
	AddMany(data []*Song)      // AddMany добавляет песни в конец плейлиста
	Get() Info                 // Get Получает текущую песню
	GetAll() []Info            // GetAll Получает все песни 						TODO конкурентный доступ
	Update(data *Info) error   // Update Обновляет данные песни 					TODO конкурентный доступ
	Remove(id uuid.UUID) error // Remove Удаляет текущую песню        				TODO можно удалить если не воспроизводится сейчас
	Play() error               // Play начинает воспроизведение текущей песни  		TODO 1) не должно блокировать остальное методы 2) след песня - автоматически; 3) воспроизведение может быть остановлено извне
	Pause() error              // Pause приостанавливает воспроизведение			TODO дальнейшее воспроизведение с момента паузы, или с нуля
	Next() error               // Next воспроизводит следующую песню 		    	TODO конкурентный доступ
	Prev() error               // Prev воспроизводит предыдущую песню 		 		TODO конкурентный доступ
}

type MusicRepository interface {
	Upload(s []*Song) error     // Upload выгружает данные плейлиста в файл, создает файл при необходимости
	Download() ([]*Song, error) // Download загружает данные плейлиста из файла
}
