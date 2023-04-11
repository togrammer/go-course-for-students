package adrepo

import (
	"context"
	"errors"
	"homework6/internal/ads"
	"homework6/internal/app"
)

type RepoInit struct {
	lastId int64
	adList []ads.Ad
}

var ErrWrongAdId = errors.New("wrong adId")
var ErrWrongUser = errors.New("invalid user")

func (a *RepoInit) Add(ctx context.Context, title string, text string, authorId int64) (ads.Ad, error) {
	defer func() {
		a.lastId++
	}()
	a.adList = append(a.adList, ads.Ad{ID: a.lastId, Title: title, Text: text, AuthorID: authorId})
	return a.adList[a.lastId], nil
}

func (a *RepoInit) ChangeStatus(ctx context.Context, adId int64, userId int64, published bool) error {
	if adId >= a.lastId {
		return ErrWrongAdId
	}
	if userId != a.adList[adId].AuthorID {
		return ErrWrongUser
	}
	a.adList[adId].Published = published
	return nil
}

func (a *RepoInit) UpdateAd(ctx context.Context, adId int64, userId int64, title string, text string) error {
	if adId >= a.lastId {
		return ErrWrongAdId
	}
	if userId != a.adList[adId].AuthorID {
		return ErrWrongUser
	}
	a.adList[adId].Title = title
	a.adList[adId].Text = text
	return nil
}

func (a *RepoInit) Find(ctx context.Context, adId int64) (ads.Ad, error) {
	if adId >= a.lastId {
		return ads.Ad{}, ErrWrongAdId
	}
	return a.adList[adId], nil
}

func New() app.Repository {
	return &RepoInit{}
}
