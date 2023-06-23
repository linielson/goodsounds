package schemas

import (
	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	Title     string
	Location  string
	Latitude  string
	Longitude string
}
