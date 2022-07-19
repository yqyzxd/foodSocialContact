package router

import (
	"foodSocialContact/ms-diner/api"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	group := engine.Group("/account")
	{
		group.GET("login", api.Login)
		group.POST("logout", api.Logout)
		group.POST("sign", api.Sign)
		group.POST("sign_month", api.GetSignInfoOfMonth)

	}

}
