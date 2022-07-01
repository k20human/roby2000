package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var audioActions = []string{
	"play",
	"pause",
	"stop",
}

func (h *handler) Audio(c *gin.Context) {
	action := c.Param("action")
	filename := c.Param("filename")

	for _, v := range audioActions {
		if v == action {
			switch action {
			case "play":
				if err := h.driver.PlaySound(filename); err != nil {
					h.logger.Error(err.Error())
					return
				}
			case "pause":

			case "stop":

			}

			c.JSON(http.StatusOK, &successResponse{Message: "OK"})
			return
		}
	}

	c.JSON(http.StatusBadRequest, &errorResponse{Message: "Action not authorized"})
}
