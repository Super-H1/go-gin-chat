package models

import (
	"gorm.io/gorm"
	"time"
)

type Psychologists struct {
	gorm.Model
	ID        uint
	Name      string    `json:"name"`
	RoomId    string    `json:"room_id"`
	AvatarId  string    `json:"avatar_id"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

func FindPsychologistsByRoomId(RoomId string) Psychologists {
	var p Psychologists
	ChatDB.Where("room_id = ?", RoomId).First(&p)

	return p
}
