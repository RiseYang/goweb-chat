package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	AvatarId  string    `json:"avatar_id"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

func AddUser(value interface{}) User {
	var user User
	user.Username = value.(map[string]interface{})["username"].(string)
	user.Password = value.(map[string]interface{})["password"].(string)
	user.AvatarId = value.(map[string]interface{})["avatar_id"].(string)
	ChatDB.Create(&user)
	return user
}

func SaveAvatarId(AvatarId string, user User) User {
	user.AvatarId = AvatarId
	ChatDB.Save(&user)
	return user
}

func FindUserByField(field, value string) User {
	var user User
	if field == "id" || field == "username" {
		ChatDB.Where(field+"= ?", value).First(&user)
	}
	return user
}

func GetOnlineUserList(uids []float64) []map[string]interface{} {
	var results []map[string]interface{}
	ChatDB.Where("id IN ?", uids).Find(&results)

	return results
}
