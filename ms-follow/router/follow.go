package router

import (
	"foodSocialContact/ms-follow/api"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	group := engine.Group("/follow")
	{
		group.GET("", api.Follow)

	}

}
