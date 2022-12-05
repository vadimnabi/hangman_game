package api

import (
	hangman "hangman_game/internal"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.LoadHTMLGlob("web/*.html")
	r.GET("/", handlerIndex)

	r.GET("/hangman", handlerHangman)

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hangman/start", handlerV1HangmanStart)
		v1.GET("/hangman/guess", handlerV1HangmanGuess)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func handlerIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func handlerHangman(c *gin.Context) {
	c.HTML(http.StatusOK, "hangman.html", nil)
}

func handlerV1HangmanStart(c *gin.Context) {
	targetWord, guessedLetters := hangman.Start()

	c.JSON(200, gin.H{
		"guessedLetters": guessedLetters,
		"word":           targetWord,
		"life":           0,
		"textImg":        readFile(getRoot() + "\\assets\\states\\hangman9"),
		"useChars":       nil,
	})
}

func handlerV1HangmanGuess(c *gin.Context) {
	c.JSON(200, gin.H{
		"word":     "_ _ _ _ _",
		"life":     9,
		"textImg":  readFile(getRoot() + "\\assets\\states\\hangman9"),
		"useChars": nil,
	})
}

func readFile(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes[:])
}

func getRoot() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}
