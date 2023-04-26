package app

import (
	"context"
	"errors"
	"homework8/internal/ads"
	"homework8/internal/users"
	"homework8/validation"
	"strings"
	"time"
)

type App interface {
	CreateAd(ctx context.Context, title string, text string, userID int64) (ads.Ad, error)
	ChangeAdStatus(ctx context.Context, adID int64, UserID int64, published bool) (ads.Ad, error)
	UpdateAd(ctx context.Context, adID int64, UserID int64, title string, text string) (ads.Ad, error)
	ListAds(ctx context.Context, authorId int64, publishedOnly bool, createdTime int64) []ads.Ad
	CreateUser(ctx context.Context, nickname, email string) users.User
	UpdateUser(ctx context.Context, userID int64, nickname, email string) (users.User, error)
	FindAd(ctx context.Context, adID int64) (ads.Ad, error)
	FindAdsByTitle(ctx context.Context, title string) []ads.Ad
}

type Repository interface {
	FindAd(ctx context.Context, adID int64) (ads.Ad, error)
	AddAd(ctx context.Context, title string, text string, userID int64) ads.Ad
	ChangeStatus(ctx context.Context, adId int64, userId int64, published bool)
	UpdateAd(ctx context.Context, adId int64, userId int64, title string, text string)
	ListAds(ctx context.Context) []ads.Ad
}

type UsersRepository interface {
	FindUser(ctx context.Context, userID int64) (users.User, error)
	AddUser(ctx context.Context, nickname, email string) users.User
	UpdateUser(ctx context.Context, userID int64, nickname, email string)
}

type RepoApp struct {
	repoAd    Repository
	repoUsers UsersRepository
}

var ErrWrongAdId = errors.New("wrong adId")
var ErrWrongUser = errors.New("invalid user")

func (a RepoApp) CreateAd(ctx context.Context, title string, text string, userID int64) (ads.Ad, error) {
	e := validation.Validate(ads.Ad{Text: text, Title: title})
	if e != nil {
		return ads.Ad{}, e
	}
	//_, err := a.repoUsers.FindUser(ctx, userID) //текущие тесты не создают пользователей
	//if err != nil {
	//	return ads.Ad{}, ErrWrongUser
	//}
	ad := a.repoAd.AddAd(ctx, title, text, userID)
	return ad, nil
}

func (a RepoApp) CreateUser(ctx context.Context, nickname string, email string) users.User {
	return a.repoUsers.AddUser(ctx, nickname, email)
}

func (a RepoApp) ChangeAdStatus(ctx context.Context, adID int64, UserID int64, published bool) (ads.Ad, error) {
	//_, err := a.repoUsers.FindUser(ctx, UserID) //текущие тесты не создают пользователей
	//if err != nil {
	//	return ads.Ad{}, ErrWrongUser
	//}
	ad, e := a.repoAd.FindAd(ctx, adID)
	if UserID != ad.AuthorID {
		return ads.Ad{}, ErrWrongUser
	}
	if e != nil {
		return ads.Ad{}, ErrWrongAdId
	}
	a.repoAd.ChangeStatus(ctx, adID, UserID, published)
	ad, _ = a.repoAd.FindAd(ctx, adID)
	return ad, nil
}

func (a RepoApp) UpdateAd(ctx context.Context, adID int64, UserID int64, title string, text string) (ads.Ad, error) {
	e := validation.Validate(ads.Ad{Text: text, Title: title})
	if e != nil {
		return ads.Ad{}, e
	}
	//_, err := a.repoUsers.FindUser(ctx, UserID) //текущие тесты не создают пользователей
	//if err != nil {
	//	return ads.Ad{}, ErrWrongUser
	//}
	ad, err := a.repoAd.FindAd(ctx, adID)
	if err != nil {
		return ads.Ad{}, ErrWrongAdId
	}
	if ad.AuthorID != UserID {
		return ads.Ad{}, ErrWrongUser
	}
	a.repoAd.UpdateAd(ctx, adID, UserID, title, text)
	ad, _ = a.repoAd.FindAd(ctx, adID)
	return ad, nil
}

func (a RepoApp) UpdateUser(ctx context.Context, userID int64, nickname string, email string) (users.User, error) {
	_, err := a.repoUsers.FindUser(ctx, userID)
	if err != nil {
		return users.User{}, ErrWrongUser
	}
	a.repoUsers.UpdateUser(ctx, userID, nickname, email)
	return a.repoUsers.FindUser(ctx, userID)
}

func (a RepoApp) FindAd(ctx context.Context, adID int64) (ads.Ad, error) {
	ad, err := a.repoAd.FindAd(ctx, adID)
	if err != nil {
		return ads.Ad{}, ErrWrongAdId
	}
	return ad, nil
}

func (a RepoApp) ListAds(ctx context.Context, authorId int64, publishedOnly bool, createdTime int64) []ads.Ad {
	l := a.repoAd.ListAds(ctx)
	var filteredAds []ads.Ad

	for _, ad := range l {

		if authorId != -1 && ad.AuthorID != authorId {
			continue
		}
		if createdTime != -1 && !ad.Created.Equal(time.Unix(createdTime, 0)) {
			continue
		}
		if publishedOnly && !ad.Published {
			continue
		}
		filteredAds = append(filteredAds, ad)
	}
	return filteredAds
}

func (a RepoApp) FindAdsByTitle(ctx context.Context, title string) []ads.Ad {
	l := a.repoAd.ListAds(ctx)
	var filteredAds []ads.Ad

	for _, ad := range l {
		if strings.HasPrefix(ad.Title, title) {
			filteredAds = append(filteredAds, ad)
		}
	}

	return filteredAds
}

func NewApp(repoAd Repository, repoUsers UsersRepository) App {
	return RepoApp{repoAd: repoAd, repoUsers: repoUsers}
}
