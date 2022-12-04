package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.LoadHTMLGlob("web/*.html")
	r.GET("/", handlerIndex)

	r.GET("/hangman", handlerHangman)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func handlerIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func handlerHangman(c *gin.Context) {
	c.HTML(http.StatusOK, "hangman.html", nil)
}
