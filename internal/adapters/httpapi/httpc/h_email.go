package httpc

import (
	"io/ioutil"
	"net/http"

	"github.com/Akanibekuly/tsarka-tz/internal/domain/errs"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/utils"
	"github.com/gin-gonic/gin"
)

func (a *St) hFindEmails(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  errs.InternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"emails": utils.FindEmails(string(data)),
	})
}
