package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const secretPassword = "morron"

func login(c *gin.Context) {
	password := c.PostForm("password")
	if password == secretPassword {
		session := sessions.Default(c)
		session.Set("loggedIn", true)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func authRequired(c *gin.Context) {
	session := sessions.Default(c)
	loggedIn := session.Get("loggedIn")
	if loggedIn == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.Next()
}
