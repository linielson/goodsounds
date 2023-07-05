package entity

import (
	"errors"
	"github.com/linielson/goodsounds/internal/dto"
	"github.com/linielson/goodsounds/pkg/entity"
	"time"
)

var (
	ErrIDIsRequired        = errors.New("id is required")
	ErrInvalidID           = errors.New("invalid id")
	ErrTitleIsRequired     = errors.New("title is required")
	ErrLocationIsRequired  = errors.New("location is required")
	ErrLatitudeIsRequired  = errors.New("latitude is required")
	ErrLongitudeIsRequired = errors.New("longitude is required")
)

// schema
type Playlist struct {
	ID        entity.ID `json:"id"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	//Musics    []Music
}

func NewPlaylist(playlist dto.Playlist) (*Playlist, error) {
	newPlaylist := &Playlist{
		ID:        entity.NewID(),
		Title:     playlist.Title,
		Location:  playlist.Location,
		Latitude:  playlist.Latitude,
		Longitude: playlist.Longitude,
		CreatedAt: time.Now(),
	}

	if err := newPlaylist.Validate(); err != nil {
		return nil, err
	}

	return newPlaylist, nil
}

func (p *Playlist) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Title == "" {
		return ErrTitleIsRequired
	}
	if p.Location == "" {
		return ErrLocationIsRequired
	}
	if p.Latitude == "" {
		return ErrLatitudeIsRequired
	}
	if p.Longitude == "" {
		return ErrLongitudeIsRequired
	}
	return nil
}
