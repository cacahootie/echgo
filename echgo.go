package echgo

import (
    "github.com/gin-gonic/gin"
)

func Echgo(c *gin.Context) {
    c.JSON(200, c.Request.URL.Query())
    // querystring, _ := json.Marshal()
    // s := string(querystring[:])
    // fmt.Fprintf(w, s)
    return
}
