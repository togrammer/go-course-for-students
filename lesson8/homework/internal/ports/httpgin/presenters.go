package httpgin

import (
	"github.com/gin-gonic/gin"
	"homework8/internal/ads"
	"homework8/internal/users"
	"time"
)

type createAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	ID       int64  `json:"user_id"`
}

type changeUserRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type adResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	AuthorID  int64     `json:"author_id"`
	Published bool      `json:"published"`
	Created   time.Time `json:"created_time"`
	Updated   time.Time `json:"updated_time"`
}

type changeAdStatusRequest struct {
	Published bool  `json:"published"`
	UserID    int64 `json:"user_id"`
}

type updateAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

func AdSuccessResponse(ad *ads.Ad) *gin.H {
	return &gin.H{
		"data": adResponse{
			ID:        ad.ID,
			Title:     ad.Title,
			Text:      ad.Text,
			AuthorID:  ad.AuthorID,
			Published: ad.Published,
		},
		"error": nil,
	}
}

func UserSuccessResponse(u *users.User) *gin.H {
	return &gin.H{
		"data": User{
			Nickname: u.Nickname,
			Email:    u.Email,
			ID:       u.ID,
		},
		"error": nil,
	}
}

func AdSuccessResponseList(ads *[]ads.Ad) *gin.H {
	var l []adResponse
	for _, ad := range *ads {
		l = append(l, adResponse{
			ID:        ad.ID,
			Title:     ad.Title,
			Text:      ad.Text,
			AuthorID:  ad.AuthorID,
			Published: ad.Published,
			Created:   ad.Created,
			Updated:   ad.Updated,
		})
	}
	return &gin.H{
		"data":  l,
		"error": nil,
	}
}

func AdErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

func UserErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}
