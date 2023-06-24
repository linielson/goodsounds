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

func (p *Playlist) FindAll(page, limit int, sort string) ([]entity.Playlist, error) {
	var playlists []entity.Playlist
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&playlists).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&playlists).Error
	}
	return playlists, err
}

func (p *Playlist) FindByID(id string) (*entity.Playlist, error) {
	var playlist entity.Playlist
	err := p.DB.First(&playlist, "id = ?", id).Error
	return &playlist, err
}

func (p *Playlist) Update(playlist *entity.Playlist) error {
	_, err := p.FindByID(playlist.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(playlist).Error
}

