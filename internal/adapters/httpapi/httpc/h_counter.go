package httpc

import (
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

	val, err = a.core.Add(val)
	if err != nil {
		if err == errs.ObjectNotFound {
			a.lg.Warnw("[COUNTER] add: object not found")
			c.JSON(http.StatusNotFound, gin.H{
				"status": "error",
				"error":  errs.ObjectNotFound,
			})
		} else {
			a.lg.Errorw("[COUNTER] add: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  errs.InternalServerError,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"value":  val,
	})
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
	val, err = a.core.Sub(val)
	if err != nil {
		if err == errs.ObjectNotFound {
			a.lg.Warnw("[COUNTER] add: object not found")
			c.JSON(http.StatusNotFound, gin.H{
				"status": "error",
				"error":  errs.ObjectNotFound,
			})
		} else {
			a.lg.Errorw("[COUNTER] add: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  errs.InternalServerError,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"value":  val,
	})

}

func (a *St) hVal(c *gin.Context) {
	val, err := a.core.Val()
	if err == errs.ObjectNotFound {
		a.lg.Warnw("[COUNTER] add: object not found")
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  errs.ObjectNotFound,
		})
		return
	}
	if err != nil {
		a.lg.Errorw("[COUNTER] add: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  errs.InternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"value":  val,
	})
}

func (a *St) hDel(c *gin.Context) {
	if err := a.core.Del(); err != nil {
		a.lg.Errorw("[COUNTER] delete", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  errs.InternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   "successfully deleted",
	})
}
