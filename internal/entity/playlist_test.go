package entity

import (
	"github.com/linielson/goodsounds/internal/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlaylist(t *testing.T) {
	dataPlaylist := dto.Playlist{Title: "Italy morning", Location: "Italy", Latitude: "45.468016", Longitude: "9.186114"}
	playlist, err := NewPlaylist(dataPlaylist)
	assert.Nil(t, err)
	assert.NotNil(t, playlist)
	assert.NotEmpty(t, playlist.ID)
	assert.Equal(t, "Italy morning", playlist.Title)
	assert.Equal(t, "Italy", playlist.Location)
	assert.Equal(t, "45.468016", playlist.Latitude)
	assert.Equal(t, "9.186114", playlist.Longitude)
}

func TestPlaylistWhenTitleIsRequired(t *testing.T) {
	dataPlaylist := dto.Playlist{Title: "", Location: "Italy", Latitude: "45.468016", Longitude: "9.186114"}
	playlist, err := NewPlaylist(dataPlaylist)
	assert.Nil(t, playlist)
	assert.Equal(t, ErrTitleIsRequired, err)
}

func TestPlaylistWhenLocationIsRequired(t *testing.T) {
	dataPlaylist := dto.Playlist{Title: "Italy morning", Location: "", Latitude: "45.468016", Longitude: "9.186114"}
	playlist, err := NewPlaylist(dataPlaylist)
	assert.Nil(t, playlist)
	assert.Equal(t, ErrLocationIsRequired, err)
}

func TestPlaylistWhenLatitudeIsRequired(t *testing.T) {
	dataPlaylist := dto.Playlist{Title: "Italy morning", Location: "Italy", Latitude: "", Longitude: "9.186114"}
	playlist, err := NewPlaylist(dataPlaylist)
	assert.Nil(t, playlist)
	assert.Equal(t, ErrLatitudeIsRequired, err)
}

func TestPlaylistWhenLongitudeIsRequired(t *testing.T) {
	dataPlaylist := dto.Playlist{Title: "Italy morning", Location: "Italy", Latitude: "45.468016", Longitude: ""}
	playlist, err := NewPlaylist(dataPlaylist)
	assert.Nil(t, playlist)
	assert.Equal(t, ErrLongitudeIsRequired, err)
}

func TestPlaylistValidate(t *testing.T) {
	dataPlaylist := dto.Playlist{Title: "Italy morning", Location: "Italy", Latitude: "45.468016", Longitude: "9.186114"}
	playlist, err := NewPlaylist(dataPlaylist)
	assert.Nil(t, err)
	assert.NotNil(t, playlist)
	assert.Nil(t, playlist.Validate())
}
