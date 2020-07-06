package user

import (
	"dast-api/internal/interface/http/presenter"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

type logInRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type userRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type tokenValidationRequest struct {
	ID    string `json:"id" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type userController struct {
	useCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) *userController {
	return &userController{useCase: useCase}
}

func (hac userController) Create(c *gin.Context) {
	var input userRequest
	if err := c.BindJSON(&input); err != nil {
		logger.Error("error parsing user create request", err)
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_parsing_user_request", err))
		return
	}

	res, err := hac.useCase.RegisterUser(input.Name, input.Email, input.Password)
	if err != nil {
		logger.Error("error creating user", err)
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_saving_user", err))
		return
	}

	logger.Infof("user:%s created successfully", input.Name)
	c.JSON(http.StatusCreated, presenter.RenderUser(res))
}

func (hac userController) LogIn(c *gin.Context) {
	var input logInRequest
	if err := c.BindJSON(&input); err != nil {
		logger.Error("error trying to parse log in request", err)
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_parsing_log_in_request", err))
		return
	}

	token, err := hac.useCase.LogIn(input.Name, input.Password)
	if err != nil {
		logger.Error("error creating login", err)
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_log_in", err))
		return
	}

	logger.Infof("user:%s logged successfully", input.Name)
	c.JSON(http.StatusCreated, presenter.RenderSuccessLogIn(input.Name, token))
}

func (hac userController) Get(c *gin.Context) {
	res, err := hac.useCase.GetUser(c.Param("id"))
	if err != nil {
		logger.Error("error getting user", err)
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_getting_user", err))
		return
	}

	c.JSON(http.StatusOK, presenter.RenderUser(res))
}

func (hac userController) ValidateToken(c *gin.Context) {
	var input tokenValidationRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_parsing_token", err))
		return
	}

	err := hac.useCase.AlreadyLogIn(input.ID, input.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, presenter.NewUnauthorized(c.Param("id")))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
