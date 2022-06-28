package main

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	username = "user"
	ucolor   = "color"
)

func loadLogin(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"login.html",
		gin.H{},
	)
}

func auth(ctx *gin.Context) {
	user := ctx.PostForm("user")
	color := ctx.PostForm("color")

	if strings.Contains(user, "<") || strings.Contains(user, ">") || user == "" {
		ctx.HTML(
			http.StatusOK,
			"login.html",
			gin.H{
				"error": "Use a valid username !!",
			},
		)
		return
	}

	session := sessions.Default(ctx)
	session.Set(username, user)
	session.Set(ucolor, color)

	if err := session.Save(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	} else {
		ctx.Redirect(http.StatusFound, "/moe/chat")
	}
}
