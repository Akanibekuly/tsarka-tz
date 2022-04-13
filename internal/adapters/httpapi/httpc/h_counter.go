package httpc

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Akanibekuly/tsarka-tz/internal/domain/errs"
	"github.com/gin-gonic/gin"
)

func (a *St) hAdd(c *gin.Context) {
	val, err := strconv.Atoi(c.Param("val"))
	if err != nil {
		a.lg.Errorw("[COUNTER] add:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  errs.BadRequest,
		})
		return
	}

	fmt.Println(val)
}

func (a *St) hSub(c *gin.Context) {
	val, err := strconv.Atoi(c.Param("val"))
	if err != nil {
		a.lg.Errorw("[COUNTER] add:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  errs.BadRequest,
		})
		return
	}

	fmt.Println(val)
}

func (a *St) hVal(c *gin.Context) {

}
