package database

import (
	"github.com/linielson/goodsounds/internal/entity"
	"gorm.io/gorm"
)

type Playlist struct {
	DB *gorm.DB
}

func NewPlaylist(db *gorm.DB) *Playlist {
	return &Playlist{DB: db}
}

func (p *Playlist) Create(playlist *entity.Playlist) error {
	return p.DB.Create(playlist).Error
}
