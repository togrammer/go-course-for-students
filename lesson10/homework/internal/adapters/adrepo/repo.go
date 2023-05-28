package adrepo

import (
	"context"
	"homework10/internal/ads"
	"homework10/internal/app"
	"sync"
	"time"
)

type RepoInit struct {
	mtx    *sync.RWMutex
	lastId int64
	adList []ads.Ad
}

func (a *RepoInit) AddAd(ctx context.Context, title string, text string, authorId int64) ads.Ad {
	a.mtx.Lock()
	defer func() {
		a.lastId++
		a.mtx.Unlock()
	}()
	a.adList = append(a.adList, ads.Ad{ID: a.lastId, Title: title, Text: text, AuthorID: authorId, Created: time.Now().UTC(), Updated: time.Now().UTC()})
	return a.adList[a.lastId]
}

func (a *RepoInit) ChangeStatus(ctx context.Context, adId int64, userId int64, published bool) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	a.adList[adId].Published = published
}

func (a *RepoInit) UpdateAd(ctx context.Context, adId int64, userId int64, title string, text string) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	a.adList[adId].Title = title
	a.adList[adId].Text = text
}

func (a *RepoInit) FindAd(ctx context.Context, adId int64) (ads.Ad, error) {
	a.mtx.RLock()
	defer a.mtx.RUnlock()
	if adId >= a.lastId {
		return ads.Ad{}, app.ErrWrongAdId
	}
	return a.adList[adId], nil
}

func (a *RepoInit) ListAds(ctx context.Context) []ads.Ad {
	a.mtx.RLock()
	defer a.mtx.RUnlock()
	return a.adList
}

func New() app.Repository {
	return &RepoInit{mtx: &sync.RWMutex{}}
}
