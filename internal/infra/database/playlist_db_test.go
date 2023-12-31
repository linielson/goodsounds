package database

import (
	"fmt"
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

func TestFinalAllPlaylists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Playlist{})
	if err != nil {
		t.Error(err)
	}
	for i := 1; i < 24; i++ {
		dataPlaylist := dto.Playlist{Title: fmt.Sprintf("Playlist %d", i), Location: "Italy", Latitude: "45.468016", Longitude: "9.186114"}
		playlist, err := entity.NewPlaylist(dataPlaylist)
		assert.NoError(t, err)
		db.Create(playlist)
	}
	playlistDB := NewPlaylist(db)
	playlists, err := playlistDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, playlists, 10)
	assert.Equal(t, "Playlist 1", playlists[0].Title)
	assert.Equal(t, "Playlist 10", playlists[9].Title)

	playlists, err = playlistDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, playlists, 10)
	assert.Equal(t, "Playlist 11", playlists[0].Title)
	assert.Equal(t, "Playlist 20", playlists[9].Title)

	playlists, err = playlistDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, playlists, 3)
	assert.Equal(t, "Playlist 21", playlists[0].Title)
	assert.Equal(t, "Playlist 23", playlists[2].Title)
}

func TestFindPlaylistByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Playlist{})
	if err != nil {
		t.Error(err)
	}
	dataPlaylist := dto.Playlist{Title: "Playlist 1", Location: "Italy", Latitude: "45.468016", Longitude: "9.186114"}
	playlist, err := entity.NewPlaylist(dataPlaylist)
	assert.NoError(t, err)
	db.Create(playlist)
	playlistDB := NewPlaylist(db)
	playlist, err = playlistDB.FindByID(playlist.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Playlist 1", playlist.Title)
}

func TestUpdatePlaylist(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Playlist{})
	if err != nil {
		t.Error(err)
	}
	dataPlaylist := dto.Playlist{Title: "Playlist 1", Location: "Italy", Latitude: "45.468016", Longitude: "9.186114"}
	playlist, err := entity.NewPlaylist(dataPlaylist)
	assert.NoError(t, err)
	db.Create(playlist)
	playlistDB := NewPlaylist(db)
	playlist.Title = "Playlist 2"
	err = playlistDB.Update(playlist)
	assert.NoError(t, err)
	playlist, err = playlistDB.FindByID(playlist.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Playlist 2", playlist.Title)
}

func TestDeletePlaylist(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Playlist{})
	if err != nil {
		t.Error(err)
	}
	dataPlaylist := dto.Playlist{Title: "Playlist 1", Location: "Italy", Latitude: "45.468016", Longitude: "9.186114"}
	playlist, err := entity.NewPlaylist(dataPlaylist)
	assert.NoError(t, err)
	db.Create(playlist)
	playlistDB := NewPlaylist(db)

	err = playlistDB.Delete(playlist.ID.String())
	assert.NoError(t, err)
	_, err = playlistDB.FindByID(playlist.ID.String())
	assert.Error(t, err)
}

//fix coverage
