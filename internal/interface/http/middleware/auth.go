package middleware

import (
	"dast-api/internal/domain/model"
	"dast-api/internal/interface/http/presenter"
	"dast-api/internal/usecase"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	hierarchyCRUD usecase.HierarchyCRUD
	userUseCase   usecase.UserUseCase
}

func NewAuthHandler(useCase usecase.HierarchyCRUD, userUseCase usecase.UserUseCase) *AuthHandler {
	return &AuthHandler{hierarchyCRUD: useCase, userUseCase: userUseCase}
}

func (a AuthHandler) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		h, err := a.hierarchyCRUD.GetHierarchy(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_getting_hierarchy", err))
			return
		}

		if h == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, presenter.NewNotFound(fmt.Sprintf("hierarchy:%s not found", c.Param(":id"))))
			return
		}

		if h.Owner != model.GuestUsername {
			token := c.GetHeader("X-Auth-Token")
			if token == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.NewUnauthorized(h.Owner))
				return
			}

			err := a.userUseCase.AlreadyLogIn(h.Owner, token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.NewUnauthorized(h.Owner))
				return
			}
		}

		c.Set(h.ID, h)
	}
}

func (a AuthHandler) ValidateUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		owner := c.Param("username")
		if owner == "guest" {
			c.AbortWithStatusJSON(http.StatusBadGateway, presenter.NewBadRequest("guest_hierarchies", errors.New("you cannot obtain guest hierarchies")))
			return
		}

		//token := c.GetHeader("X-Auth-Token")
		//if token == "" {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.NewUnauthorized(owner))
		//	return
		//}
		//
		//err := a.userUseCase.AlreadyLogIn(c.Param("username"), token)
		//if err != nil {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.NewUnauthorized(owner))
		//	return
		//}
	}
}

func (a AuthHandler) IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
