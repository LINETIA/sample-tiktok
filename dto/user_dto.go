package dto

import "Gin/model"

type UserDto struct {
	Nickname string `json: "nickname"`
	ID       uint   `json: "ID`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Nickname: user.Nickname,
		ID:       user.ID,
	}
}