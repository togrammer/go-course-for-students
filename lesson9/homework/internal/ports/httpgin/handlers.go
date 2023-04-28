package httpgin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"homework9/internal/app"
)

// Метод для создания объявления (ad)
func createAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody createAdRequest
		err := c.ShouldBindJSON(&reqBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		ad, er := a.CreateAd(c, reqBody.Title, reqBody.Text, reqBody.UserID)
		if er != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(er))
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, AdErrorResponse(err))
		}
		c.JSON(http.StatusOK, AdSuccessResponse(&ad))
	}
}

// Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
func changeAdStatus(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody changeAdStatusRequest
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		adIDstring := c.Param("ad_id")
		adID, err := strconv.Atoi(adIDstring)
		if err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		ad, er := a.ChangeAdStatus(c, int64(adID), reqBody.UserID, reqBody.Published)
		if er != nil {
			if errors.Is(er, app.ErrWrongUser) {
				c.JSON(http.StatusForbidden, AdErrorResponse(er))
				return
			} else {
				c.JSON(http.StatusBadRequest, AdErrorResponse(er))
				return
			}
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, AdErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, AdSuccessResponse(&ad))
	}
}

// Метод для обновления текста(Text) или заголовка(Title) объявления
func updateAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody updateAdRequest
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		adIDstring := c.Param("ad_id")
		adID, err := strconv.Atoi(adIDstring)
		if err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		ad, er := a.UpdateAd(c, int64(adID), reqBody.UserID, reqBody.Title, reqBody.Text)
		if er != nil {
			if errors.Is(er, app.ErrWrongUser) {
				c.JSON(http.StatusForbidden, AdErrorResponse(er))
				return
			} else {
				c.JSON(http.StatusBadRequest, AdErrorResponse(er))
				return
			}
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, AdErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, AdSuccessResponse(&ad))
	}
}

func listAds(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {

		authorIdstring := c.Query("author_id")
		authorId := int64(-1)
		if authorIdstring != "" {
			val, err := strconv.Atoi(authorIdstring)
			if err != nil {
				c.JSON(http.StatusBadRequest, AdErrorResponse(err))
				return
			}
			authorId = int64(val)
		}
		publishedOnlystr := c.Query("published_only")
		publishedOnly := true
		if publishedOnlystr != "" {
			val, err := strconv.ParseBool(publishedOnlystr)
			if err != nil {
				c.JSON(http.StatusBadRequest, AdErrorResponse(err))
				return
			}
			publishedOnly = val
		}

		createdTimestr := c.Query("created_time")
		createdTime := int64(-1)
		if publishedOnlystr != "" {
			val, err := strconv.ParseInt(createdTimestr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, AdErrorResponse(err))
				return
			}
			createdTime = val
		}

		ads := a.ListAds(c, authorId, publishedOnly, createdTime)

		c.JSON(http.StatusOK, AdSuccessResponseList(&ads))
	}
}

func getAdById(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		adIdstr := c.Param("ad_id")
		adId, err := strconv.Atoi(adIdstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}
		ad, e := a.FindAd(c, int64(adId))
		if e != nil {
			if errors.Is(e, app.ErrWrongAdId) {
				c.JSON(http.StatusBadRequest, AdErrorResponse(e))
			}
			c.JSON(http.StatusInternalServerError, AdErrorResponse(e))
			return
		}
		c.JSON(http.StatusOK, AdSuccessResponse(&ad))
	}
}

func createUser(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody User
		err := c.ShouldBindJSON(&reqBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, UserErrorResponse(err))
			return
		}

		user := a.CreateUser(c, reqBody.Nickname, reqBody.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, UserErrorResponse(err))
		}
		c.JSON(http.StatusOK, UserSuccessResponse(&user))
	}
}

func changeUser(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody changeUserRequest
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}
		UserIdstr := c.Param("user_id")
		userId, err := strconv.Atoi(UserIdstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, UserErrorResponse(err))
			return
		}
		user, e := a.UpdateUser(c, int64(userId), reqBody.Nickname, reqBody.Email)
		if e != nil {
			if errors.Is(e, app.ErrWrongUser) {
				c.JSON(http.StatusBadRequest, UserErrorResponse(e))
			}
			c.JSON(http.StatusInternalServerError, UserErrorResponse(e))
			return
		}
		c.JSON(http.StatusOK, UserSuccessResponse(&user))
	}
}

func getAdsByTitle(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Query("title")
		ads := a.FindAdsByTitle(c, title)
		c.JSON(http.StatusOK, AdSuccessResponseList(&ads))
	}
}
