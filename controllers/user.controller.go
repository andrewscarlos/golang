package controllers

import (
	"github.com/andrewscarlos/golang/models"
	"github.com/andrewscarlos/golang/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func New(userservice services.UserServiceInterface) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err := u.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "success"})
}

func (u *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := u.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": user})
}

func (u *UserController) GetAll(ctx *gin.Context) {
	users, err := u.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err := u.UserService.UpdateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "success"})
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := u.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRouter := rg.Group("/user")
	userRouter.POST("/create", uc.CreateUser)
	userRouter.GET("/get/:name", uc.GetUser)
	userRouter.GET("/getall", uc.GetAll)
	userRouter.PATCH("/update", uc.UpdateUser)
	userRouter.DELETE("/delete/:name", uc.DeleteUser)
}
