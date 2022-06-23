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
				h.driver.Forward()
			case "backward":
				h.driver.Backward()
			case "left":
				h.driver.Left()
			case "right":
				h.driver.Right()
			}

			c.JSON(http.StatusOK, &successResponse{Message: "OK"})
			return
		}
	}

	c.JSON(http.StatusBadRequest, &errorResponse{Message: "Action not authorized"})
}
