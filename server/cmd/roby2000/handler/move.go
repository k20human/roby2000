package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var moveActions = [4]string{
	"forward",
	"backward",
	"left",
	"right",
}

func (h *handler) Move(c *gin.Context) {
	action := c.Param("action")

	for _, v := range moveActions {
		if v == action {
			switch action {
			case "forward":
				if err := h.driver.MoveForward(); err != nil {
					h.logger.Error(err.Error())
					return
				}
			case "backward":
				h.driver.MoveBackward()
			case "left":
				h.driver.MoveLeft()
			case "right":
				h.driver.MoveRight()
			}

			c.JSON(http.StatusOK, &successResponse{Message: "OK"})
			return
		}
	}

	c.JSON(http.StatusBadRequest, &errorResponse{Message: "Action not authorized"})
}
