package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

// }

// func (h *Handler) CreateUser(c *gin.Context) {
// 	var u CreateUserReq
// 	if err := c.ShouldBindJSON(&u); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	res, err := h.Service.CreateUser(c.Request.Context(), &u)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.SetCookie("jwt", res.accessToken, 3600, "/", "localhost", false, true)
// 	c.JSON(http.StatusOK, res)
// }

func (h *Handler) CreateRoom(c *gin.Context) {
	var r CreateRoomReq
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Service.CreateRoom(c.Request.Context(), &r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful!"})
}
