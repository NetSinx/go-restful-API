package user

import (
	"github.com/NetSinx/go-restful-API/middleware"
	"github.com/NetSinx/go-restful-API/model/domain"
	service "github.com/NetSinx/go-restful-API/service/user"
	"github.com/NetSinx/go-restful-API/utils"
	"github.com/gin-gonic/gin"
)

type usercontroller struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *usercontroller {
	return &usercontroller{
		UserService: userService,
	}
}

func (controller *usercontroller) Register(ctx *gin.Context) {
	var user domain.Register

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		utils.BadRequest(ctx)

		return
	}

	controller.UserService.Register(user)

	utils.ResponseCreateUser(ctx)
}

func (controller *usercontroller) Delete(ctx *gin.Context) {
	nameUser := ctx.Param("nameUser")

	err := controller.UserService.Delete(nameUser)
	if err != nil {
		utils.NotFound(ctx)

		return
	}

	utils.ResponseDeleteUser(ctx)
}

func (controller *usercontroller) Login(ctx *gin.Context) {
	var users domain.Login

	err := ctx.ShouldBindJSON(&users)
	if err != nil {
		utils.BadRequest(ctx)
		
		return
	}
	
	err = controller.UserService.Login(users)
	if err != nil {
		utils.ResponseLoginFailed(ctx)

		return
	}

	token := middleware.GenerateToken()

	utils.ResponseLoginSuccess(ctx, token)
}