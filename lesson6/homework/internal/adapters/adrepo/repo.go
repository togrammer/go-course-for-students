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

func (a *RepoInit) Add(ctx context.Context, title string, text string, authorId int64) ads.Ad {
	defer func() {
		a.lastId++
	}()
	a.adList = append(a.adList, ads.Ad{ID: a.lastId, Title: title, Text: text, AuthorID: authorId})
	return a.adList[a.lastId]
}

func (a *RepoInit) ChangeStatus(ctx context.Context, adId int64, userId int64, published bool) {
	a.adList[adId].Published = published
}

func (a *RepoInit) UpdateAd(ctx context.Context, adId int64, userId int64, title string, text string) {
	a.adList[adId].Title = title
	a.adList[adId].Text = text
}

func (a *RepoInit) Find(ctx context.Context, adId int64) (ads.Ad, error) {
	if adId >= a.lastId {
		return ads.Ad{}, errors.New("wrong adId")
	}
	return a.adList[adId], nil
}

func New() app.Repository {
	return &RepoInit{}
}
