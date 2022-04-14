package httpc

import (
	"fmt"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/errs"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (a *St) hCalc(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		a.lg.Errorw("[HANDLER] read from body", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  errs.InternalServerError,
		})
		return
	}

	fmt.Println(string(data))

	id, err := a.core.Hash.Calc(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  errs.InternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"id":     id,
	})
}

func (a *St) hResult(c *gin.Context) {
	id := c.Param("id")

	result, err := a.core.Hash.Result(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  errs.InternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
