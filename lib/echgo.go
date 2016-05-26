package echgo

import (
    "github.com/gin-gonic/gin"
)

func Echgo(c *gin.Context) {
    c.JSON(200, c.Request.URL.Query())
    return
}
