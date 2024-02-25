package domain

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type Song struct {
	ID       uuid.UUID
	Name     string
	Duration time.Duration
	Meta     Meta
}

type Meta struct {
	mu        sync.Mutex
	PlayedFor time.Duration
	IsPlaying bool
}

// Сохранение данных будет происходить при graceful shutdown

type MusicService interface {
	Add(data *Song) error       // Add добавляет в конец плейлиста песню 	 	TODO конкурентный доступ
	AddMany(data []*Song) error // AddMany добавляет все песни в конец плейлиста
	Get() (*Song, error)        // Get Получает текущую песню
	GetAll() ([]*Song, error)   // GetAll Получает все песни 					TODO конкурентный доступ
	Update(data *Song) error    // Update Обновляет данные песни
	Remove(id uuid.UUID) error  // Remove удаляют необходимую песню        		TODO можно удалить если не воспроизводится сейчас
	Play() error                // Play начинает воспроизведение текущей песни  	TODO 1) не должно блокировать остальное методы 2)  ки; 3) воспроизведение может быть остановлено извне
	Pause() error               // Pause приостанавливает воспроизведение		TODO дальнейшее воспроизведение с момента паузы, или с нуля
	Next() error                // Next воспроизводит следующую песню 		    TODO конкурентный доступ
	Prev() error                // Prev воспроизводит предыдущую песню 		 	TODO конкурентный доступ
}

type MusicRepository interface {
	Upload(s []*Song) error     // Upload выгружает данные плейлиста в файл, создает файл при необходимости
	Download() ([]*Song, error) // Download загружает данные плейлиста из файла
}
