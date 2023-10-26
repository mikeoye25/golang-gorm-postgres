package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imprinno-solutions/golang-gorm-postgres/controllers"
	"github.com/imprinno-solutions/golang-gorm-postgres/middlewares"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}
