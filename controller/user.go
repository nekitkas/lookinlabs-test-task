package controller

import (
	"lookinlabs-test/model"
	"lookinlabs-test/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type userController struct {
	connection repository.Connection
}

func NewUserController(connection repository.Connection) UserController {
	return &userController{
		connection: connection,
	}
}

func (u *userController) CreateUser(ctx *gin.Context) {
	user := model.NewUser()
	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if user.Name == "" || user.Email == "" {
		ctx.JSON(400, gin.H{"error": "name and email is required"})
		return
	}

	u.connection.DB().Create(user)
	ctx.JSON(http.StatusCreated, user)
}

func (u *userController) GetUsers(ctx *gin.Context) {
	var users []model.User
	u.connection.DB().Find(&users)
	ctx.JSON(http.StatusOK, users)
}

func (u *userController) GetUserByID(ctx *gin.Context) {
	var user model.User
	u.connection.DB().First(&user, ctx.Param("id"))
	if user.ID == 0 {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *userController) UpdateUser(ctx *gin.Context) {
	var user model.User
	u.connection.DB().First(&user, ctx.Param("id"))
	if user.ID == 0 {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u.connection.DB().Save(&user)
	ctx.JSON(http.StatusOK, user)
}
