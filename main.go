package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	setupRoutes(r)
	r.Run("0.0.0.0:8080")
}

func setupRoutes(r *gin.Engine) {
	lightController := NewLightController()

	r.GET("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		loggedIn := session.Get("loggedIn")
		if loggedIn != nil {
			c.Redirect(http.StatusFound, "/light")
			return
		}
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", login)
	r.GET("/logout", logout)

	protected := r.Group("/")
	protected.Use(authRequired)
	{
		protected.GET("/light", func(c *gin.Context) {
			c.HTML(http.StatusOK, "light.html", nil)
		})
		protected.GET("/light/:state", HandleLightState(lightController))
	}
}
