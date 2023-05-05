package app

import (
	"context"
	"errors"
	"homework6/internal/ads"
	"homework6/validation"
)

type App interface {
	CreateAd(ctx context.Context, title string, text string, userID int64) (ads.Ad, error)
	ChangeAdStatus(ctx context.Context, adID int64, UserID int64, published bool) (ads.Ad, error)
	UpdateAd(ctx context.Context, adID int64, UserID int64, title string, text string) (ads.Ad, error)
}

type Repository interface {
	Find(ctx context.Context, adID int64) (ads.Ad, error)
	Add(ctx context.Context, title string, text string, userID int64) ads.Ad
	ChangeStatus(ctx context.Context, adId int64, userId int64, published bool)
	UpdateAd(ctx context.Context, adId int64, userId int64, title string, text string)
}

type RepoApp struct {
	repo Repository
}

var ErrWrongAdId = errors.New("wrong adId")
var ErrWrongUser = errors.New("invalid user")

func (a RepoApp) CreateAd(ctx context.Context, title string, text string, userID int64) (ads.Ad, error) {
	e := validation.Validate(ads.Ad{Text: text, Title: title})
	if e != nil {
		return ads.Ad{}, e
	}
	ad := a.repo.Add(ctx, title, text, userID)
	return ad, nil
}

func (a RepoApp) ChangeAdStatus(ctx context.Context, adID int64, UserID int64, published bool) (ads.Ad, error) {
	ad, e := a.repo.Find(ctx, adID)
	if UserID != ad.AuthorID {
		return ads.Ad{}, ErrWrongUser
	}
	if e != nil {
		return ads.Ad{}, ErrWrongAdId
	}
	a.repo.ChangeStatus(ctx, adID, UserID, published)
	ad, _ = a.repo.Find(ctx, adID)
	return ad, nil
}

func (a RepoApp) UpdateAd(ctx context.Context, adID int64, UserID int64, title string, text string) (ads.Ad, error) {
	e := validation.Validate(ads.Ad{Text: text, Title: title})
	if e != nil {
		return ads.Ad{}, e
	}
	ad, err := a.repo.Find(ctx, adID)
	if err != nil {
		return ads.Ad{}, ErrWrongAdId
	}
	if ad.AuthorID != UserID {
		return ads.Ad{}, ErrWrongUser
	}
	a.repo.UpdateAd(ctx, adID, UserID, title, text)
	ad, _ = a.repo.Find(ctx, adID)
	return ad, nil
}

func NewApp(repo Repository) App {
	return RepoApp{repo: repo}
}
