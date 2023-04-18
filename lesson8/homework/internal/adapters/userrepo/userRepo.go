package userrepo

import (
	"context"
	"errors"
	"homework8/internal/app"
	"homework8/internal/users"
)

type Users struct {
	lastId   int64
	userList []users.User
}

func (u Users) FindUser(ctx context.Context, userID int64) (users.User, error) {
	if userID >= u.lastId {
		return users.User{}, errors.New("wrong userId")
	}
	return u.userList[userID], nil
}

func (u Users) AddUser(ctx context.Context, nickname, email string) users.User {
	defer func() {
		u.lastId++
	}()
	u.userList = append(u.userList, users.User{ID: u.lastId, Nickname: nickname, Email: email})
	return u.userList[u.lastId]
}

func (u Users) UpdateUser(ctx context.Context, userID int64, nickname, email string) {
	u.userList[userID].Nickname = nickname
	u.userList[userID].Email = email
}

func New() app.UsersRepository {
	return &Users{}
}
