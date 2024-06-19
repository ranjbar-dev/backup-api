package apicontroller

import (
	"github.com/gin-gonic/gin"
)

func (controller *Controller) Test(c *gin.Context) {

	controller.ok(c, nil)
}
