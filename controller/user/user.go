package user

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Login(ctx *gin.Context)
}