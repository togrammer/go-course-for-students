package userrepo

import (
	"context"
	"homework10/internal/app"
	"homework10/internal/users"
	"sync"
)

type Users struct {
	lastId   int64
	userList []users.User
	mtx      *sync.RWMutex
}

func (u *Users) FindUser(ctx context.Context, userID int64) (users.User, error) {
	u.mtx.RLock()
	defer u.mtx.RUnlock()
	if userID >= int64(len(u.userList)) {
		return users.User{}, app.ErrWrongUserId
	}
	return u.userList[userID], nil
}

func (u *Users) AddUser(ctx context.Context, nickname, email string) users.User {
	u.mtx.Lock()
	defer func() {
		u.lastId++
		u.mtx.Unlock()
	}()
	user := users.User{ID: u.lastId, Nickname: nickname, Email: email}
	u.userList = append(u.userList, user)
	return u.userList[len(u.userList)-1]
}

func (u *Users) UpdateUser(ctx context.Context, userID int64, nickname, email string) {
	u.mtx.Lock()
	defer u.mtx.Unlock()
	u.userList[userID].Nickname = nickname
	u.userList[userID].Email = email
}

func New() app.UsersRepository {
	return &Users{mtx: &sync.RWMutex{}}
}
