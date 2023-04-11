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
	Add(ctx context.Context, title string, text string, userID int64) (ads.Ad, error)
	ChangeStatus(ctx context.Context, adId int64, userId int64, published bool) error
	UpdateAd(ctx context.Context, adId int64, userId int64, title string, text string) error
}

type RepoApp struct {
	repo Repository
}

func (a RepoApp) CreateAd(ctx context.Context, title string, text string, userID int64) (ads.Ad, error) {
	ad, err := a.repo.Add(ctx, title, text, userID)
	if err != nil {
		return ads.Ad{}, err
	}
	e := a.Validate(ctx, ad.ID)
	return ad, e
}

func (a RepoApp) ChangeAdStatus(ctx context.Context, adID int64, UserID int64, published bool) (ads.Ad, error) {
	err := a.repo.ChangeStatus(ctx, adID, UserID, published)
	if err != nil {
		return ads.Ad{}, err
	}
	ad, _ := a.repo.Find(ctx, adID)
	e := a.Validate(ctx, adID)
	return ad, e
}

func (a RepoApp) UpdateAd(ctx context.Context, adID int64, UserID int64, title string, text string) (ads.Ad, error) {
	err := a.repo.UpdateAd(ctx, adID, UserID, title, text)
	if err != nil {
		return ads.Ad{}, err
	}
	ad, _ := a.repo.Find(ctx, adID)
	e := a.Validate(ctx, adID)
	return ad, e
}

func (a RepoApp) Validate(ctx context.Context, adId int64) error {
	ad, err := a.repo.Find(ctx, adId)
	if err != nil {
		return err
	}
	e := validation.Validate(ad)
	if e != nil {
		return errors.New("validation error")
	}
	return nil
}

func NewApp(repo Repository) App {
	return RepoApp{repo: repo}
}
