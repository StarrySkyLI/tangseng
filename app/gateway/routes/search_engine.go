package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/tangseng/app/gateway/http"
)

func SearchRegisterHandlers(rg *gin.RouterGroup) {
	favoriteGroup := rg.Group("/search_engine")
	{
		favoriteGroup.GET("/search", http.SearchEngineSearch)
		favoriteGroup.GET("/query", http.WordAssociation)
	}
}
