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
		c.Redirect(http.StatusFound, "/lights")
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func authRequired(c *gin.Context) {
	session := sessions.Default(c)
	loggedIn := session.Get("loggedIn")
	if loggedIn == nil {
		// Redirect to login page if not logged in
		c.Redirect(http.StatusFound, "/login")
		c.Abort() // Prevents further handlers from executing
		return
	}
	c.Next()
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}
	c.Redirect(http.StatusFound, "/login")
}
