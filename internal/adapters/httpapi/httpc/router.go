package httpc

import "github.com/gin-gonic/gin"

func (a *St) router() *gin.Engine {
	router := gin.Default()

	rest := router.Group("/rest")
	{
		rest.POST("/substr/find", a.hSubstrFind)
		rest.POST("/email/check", a.hFindEmails)

		counter := rest.Group("/counter")
		{
			counter.POST("/add/:val", a.hAdd)
			counter.POST("/sub/:val", a.hSub)
			counter.GET("/val", a.hVal)
			counter.DELETE("/val", a.hDel)
		}

		user := rest.Group("/user")
		{
			user.GET("/:id", a.hUserGet)
			user.POST("/", a.hUserCreate)
			user.POST("/:id", a.hUserUpdate)
			user.DELETE("/:id", a.hUserDelete)
		}
	}

	return router
}
