package database

import "github.com/linielson/goodsounds/internal/entity"

type PlaylistInterface interface {
	Create(playlist *entity.Playlist) error
	FindAll(page, limit int, sort string) ([]entity.Playlist, error)
	FindByID(id string) (*entity.Playlist, error)
	Update(playlist *entity.Playlist) error
	Delete(id string) error
}
