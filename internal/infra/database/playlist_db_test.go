package database

import (
	"github.com/linielson/goodsounds/internal/dto"
	"github.com/linielson/goodsounds/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateNewPlaylist(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Playlist{})
	if err != nil {
		t.Error(err)
	}
	dataPlaylist := dto.Playlist{Title: "Italy morning", Location: "Italy", Latitude: "45.468016", Longitude: "9.186114"}
	playlist, err := entity.NewPlaylist(dataPlaylist)
	assert.NoError(t, err)
	playlistDB := NewPlaylist(db)
	err = playlistDB.Create(playlist)
	assert.NoError(t, err)
	assert.NotEmpty(t, playlist.ID)
}
