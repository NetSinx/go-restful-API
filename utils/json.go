package utils

import (
	"net/http"

	"github.com/NetSinx/go-restful-API/model/domain"
	"github.com/NetSinx/go-restful-API/model/web"
	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, ErrorBadRequest{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: "Request not valid!",
	})
}

func Unauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, ErrorUnauthorized{
		Code: http.StatusUnauthorized,
		Status: "UNAUTHORIZED!",
	})
}

func NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, ErrorNotFound{
		Code: http.StatusNotFound,
		Status: "Not Found",
		Message: "Data not found!",
	})
}

func ResponseFindAll(ctx *gin.Context, products []domain.Product) {
	ctx.JSON(http.StatusOK, web.ProductResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: products,
	})
}

func ResponseDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, web.DeleteProductResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: "data deleted successfully!",
	})
}

func WriteResponse(ctx *gin.Context, product domain.Product) {
	ctx.JSON(http.StatusOK, web.ProductResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: product,
	})
}

func ResponseCreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, web.UserResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Registration success!",
	})
}

func ResponseDeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, web.UserResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Delete user success!",
	})
}

func ResponseLoginSuccess(ctx *gin.Context, token string) {
	ctx.JSON(http.StatusOK, web.Login{
		Token: token,
	})
}

func ResponseLoginFailed(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, web.UserResponse{
		Code: http.StatusInternalServerError,
		Status: "Internal Server Error",
		Message: "Invalid credentials!",
	})
}