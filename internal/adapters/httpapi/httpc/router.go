package httpc

import "github.com/gin-gonic/gin"

func (a *St) router() *gin.Engine {
	router := gin.Default()

	rest := router.Group("/rest")
	{
		rest.POST("/substr/find", a.hSubstrFind)
		rest.POST("/email/check", a.hFindEmails)
	}

	return router
}
