package httpc

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/Akanibekuly/tsarka-tz/internal/domain/errs"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/utils"
	"github.com/gin-gonic/gin"
)

var re = regexp.MustCompile(`^[\w]+$`)

func (a *St) hSubstrFind(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  errs.InternalServerError,
		})
		return
	}

	if !re.Match(data) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  errs.BadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"substr": utils.FindSubstr(string(data)),
	})
}
