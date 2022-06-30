package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/k20human/roby2000/pkg/robot"
	"github.com/samber/lo"
	"net/http"
	"strings"
)

var lightActions = []string{
	robot.On,
	robot.Off,
}

var lightTypes = []string{
	"front",
	"back",
	"blinking-left",
	"blinking-right",
}

func (h *handler) Light(c *gin.Context) {
	action := c.Param("action")
	light := c.Param("type")
	color := c.Param("color")

	if !lo.Contains(lightActions, action) {
		c.JSON(http.StatusBadRequest, &errorResponse{Message: "Action not authorized"})
	}

	for _, v := range lightTypes {
		if v == light {
			switch light {
			case "front":
				if err := h.driver.LightsFront(action, color); err != nil {
					h.logAndServeError(c, err)
					return
				}
			case "back":
				if err := h.driver.LightsBack(action, color); err != nil {
					h.logAndServeError(c, err)
					return
				}
			case "blinking-left", "blinking-right":
				direction := strings.Split(light, "-")[1]

				if err := h.driver.LightsBlinking(direction); err != nil {
					h.logAndServeError(c, err)
					return
				}
			}

			c.JSON(http.StatusOK, &successResponse{Message: "OK"})
			return
		}
	}

	c.JSON(http.StatusBadRequest, &errorResponse{Message: "Type not authorized"})
}
