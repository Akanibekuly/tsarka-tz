package httpc

import (
	"net/http"
	"strconv"

	"github.com/Akanibekuly/tsarka-tz/internal/domain/entities"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/errs"
	"github.com/gin-gonic/gin"
)

func (a *St) hUserGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		a.lg.Errorw("[HANDLER] user get", err, "id", id)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  errs.BadRequest,
		})
		return
	}

	user, err := a.core.User.Get(id)
	if err != nil {
		if err == errs.ObjectNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "error",
				"error":  errs.ObjectNotFound,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  errs.InternalServerError,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"user":   user,
	})
}

func (a *St) hUserCreate(c *gin.Context) {
	var user entities.UserSt
	if err := c.ShouldBindJSON(&user); err != nil {
		a.lg.Errorw("[HANDLER] user create", err, "user", user)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  errs.BadRequest,
		})
		return
	}

	id, err := a.core.User.Create(&user)
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

func (a *St) hUserUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		a.lg.Errorw("[HANDLER] user update", err, "id", id)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  errs.BadRequest,
		})
		return
	}

	var user entities.UserUpdateSt
	if err := c.ShouldBindJSON(&user); err != nil {
		a.lg.Errorw("[HANDLER] user update", err, "user", user)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  errs.BadRequest,
		})
		return
	}

	if err := a.core.User.Update(id, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  errs.InternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"info":   "successfully updated",
	})
}

func (a *St) hUserDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		a.lg.Errorw("[HANDLER] user update", err, "id", id)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  errs.BadRequest,
		})
		return
	}

	if err := a.core.User.Delete(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "err",
			"error":  errs.InternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"info":   "successfully deleted",
	})
}
